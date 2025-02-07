package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-web-server/models"
	"golang-web-server/repositories"
	"golang-web-server/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	UserRepo *repositories.UserRepository
}

func NewUserController(userRepo *repositories.UserRepository) *UserController {
	return &UserController{
		UserRepo: userRepo,
	}
}

func (uc *UserController) RegisterNewUser(c *gin.Context) {
	var (
		log = utils.Logger
		req = &models.User{}
	)

	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("Failed to bind JSON: ", err)
		utils.SendResponse(c, http.StatusBadRequest, "Invalid data", nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Validation failed: ", err)
		utils.SendResponse(c, http.StatusBadRequest, "Validation failed", nil)
		return
	}

	if err := req.HashPassword(); err != nil {
		log.Error("Failed to hash password: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	if err := uc.UserRepo.CreateUser(c.Request.Context(), req); err != nil {
		log.Error("Failed to create user: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user", nil)
		return
	}

	resp := req
	resp.Password = ""

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (uc *UserController) Login(c *gin.Context) {
	var (
		log = utils.Logger
		req = &models.User{}
	)

	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("Failed to bind JSON: ", err)
		utils.SendResponse(c, http.StatusBadRequest, "Invalid data", nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Validation failed: ", err)
		utils.SendResponse(c, http.StatusBadRequest, "Validation failed", nil)
		return
	}

	userData, err := uc.UserRepo.GetUserByUsername(c.Request.Context(), req.Username)
	if err != nil {
		log.Error("Failed to get user: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get user", nil)
		return
	}

	err = userData.CheckPassword(req.Password)
	if err != nil {
		log.Error("Failed to check password: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Invalid username or password", nil)
		return
	}

	resp := userData
	resp.Password = ""

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (uc *UserController) GetUser(c *gin.Context) {
	var (
		log = utils.Logger
	)

	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)

	resp, err := uc.UserRepo.GetUserByID(c.Request.Context(), userId)
	if err != nil {
		log.Error("Failed to get user: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get user", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}
