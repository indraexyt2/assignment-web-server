package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-web-server/models"
	"golang-web-server/repositories"
	"golang-web-server/utils"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductRepo *repositories.ProductRepository
}

func NewProductController(productRepo *repositories.ProductRepository) *ProductController {
	return &ProductController{
		ProductRepo: productRepo,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var (
		log = utils.Logger
		req = &models.Product{}
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

	if err := pc.ProductRepo.CreateProduct(c.Request.Context(), req); err != nil {
		log.Error("Failed to create product: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to create product", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", req)
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var (
		log = utils.Logger
	)

	resp, err := pc.ProductRepo.GetProducts(c.Request.Context())
	if err != nil {
		log.Error("Failed to get products: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get products", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	var (
		log = utils.Logger
	)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	resp, err := pc.ProductRepo.GetProductByID(c.Request.Context(), id)
	if err != nil {
		log.Error("Failed to get product: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get product", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var (
		log = utils.Logger
		req = &models.Product{}
	)

	productIdStr := c.Param("id")
	productIdInt, _ := strconv.Atoi(productIdStr)

	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("Failed to bind JSON: ", err)
		utils.SendResponse(c, http.StatusBadRequest, "Invalid data", nil)
		return
	}

	req.ID = productIdInt
	if err := req.Validate(); err != nil {
		log.Error("Validation failed: ", err)
		utils.SendResponse(c, http.StatusBadRequest, "Validation failed", nil)
		return
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Error("Failed to marshal JSON: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to marshal JSON", nil)
		return
	}

	var reqMap map[string]interface{}
	if err := json.Unmarshal(reqJson, &reqMap); err != nil {
		log.Error("Failed to unmarshal JSON: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to unmarshal JSON", nil)
		return
	}

	if err := pc.ProductRepo.UpdateProduct(c.Request.Context(), req.ID, reqMap); err != nil {
		log.Error("Failed to update product: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to update product", nil)
		return
	}

	resp, err := pc.ProductRepo.GetProductByID(c.Request.Context(), req.ID)
	if err != nil {
		log.Error("Failed to get product: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get product", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	var (
		log = utils.Logger
	)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := pc.ProductRepo.DeleteProduct(c.Request.Context(), id); err != nil {
		log.Error("Failed to delete product: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to delete product", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", nil)
}

func (pc *ProductController) GetInventoryByProductID(c *gin.Context) {
	var (
		log = utils.Logger
	)

	productIDStr := c.Param("id")
	productID, _ := strconv.Atoi(productIDStr)

	resp, err := pc.ProductRepo.GetInventoryByProductID(c.Request.Context(), productID)
	if err != nil {
		log.Error("Failed to get inventory: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get inventory", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (pc *ProductController) UpdateInventory(c *gin.Context) {
	var (
		log = utils.Logger
		req = &models.Inventory{}
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

	if err := pc.ProductRepo.UpdateInventory(c.Request.Context(), req.ProductID, req); err != nil {
		log.Error("Failed to update inventory: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to update inventory", nil)
		return
	}

	resp, err := pc.ProductRepo.GetInventoryByProductID(c.Request.Context(), req.ProductID)
	if err != nil {
		log.Error("Failed to get inventory: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get inventory", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}

func (pc *ProductController) CreateNewOrder(c *gin.Context) {
	var (
		log = utils.Logger
		req = &models.Order{}
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

	if err := pc.ProductRepo.CreateNewOrder(c.Request.Context(), req); err != nil {
		log.Error("Failed to create order: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to create order", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", nil)
}

func (pc *ProductController) GetOrder(c *gin.Context) {
	var (
		log = utils.Logger
	)

	orderIDStr := c.Param("id")
	orderID, _ := strconv.Atoi(orderIDStr)

	resp, err := pc.ProductRepo.GetOrder(c.Request.Context(), orderID)
	if err != nil {
		log.Error("Failed to get order: ", err)
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to get order", nil)
		return
	}

	utils.SendResponse(c, http.StatusOK, "success", resp)
}
