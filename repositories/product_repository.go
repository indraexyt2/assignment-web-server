package repositories

import (
	"context"
	"errors"
	"golang-web-server/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, req *models.Product) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *ProductRepository) UpdateInventory(ctx context.Context, productID int, req *models.Inventory) error {
	var inventory models.Inventory
	err := r.DB.WithContext(ctx).Where("product_id = ?", productID).First(&inventory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newInventory := &models.Inventory{
				ProductID: productID,
				Quantity:  req.Quantity,
				Location:  req.Location,
			}

			return r.DB.WithContext(ctx).Create(newInventory).Error
		}
	}

	inventory.Quantity += req.Quantity
	return r.DB.WithContext(ctx).Save(&inventory).Error
}

func (r *ProductRepository) GetInventoryByProductID(ctx context.Context, productID int) (*models.Inventory, error) {
	var inventory models.Inventory
	if err := r.DB.WithContext(ctx).Preload("Product").Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	var product models.Product
	if err := r.DB.WithContext(ctx).Preload("Inventory").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.WithContext(ctx).Preload("Inventory").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id int, req map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Updates(req).Error
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.Product{}).Error
}

func (r *ProductRepository) CreateNewOrder(ctx context.Context, req *models.Order) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range req.OrderItems {
			var inventory models.Inventory
			if err := tx.Where("product_id = ?", item.ProductID).First(&inventory).Error; err != nil {
				return err
			}

			if inventory.Quantity < item.Quantity {
				return errors.New("not enough inventory")
			}

			inventory.Quantity -= item.Quantity
			if err := tx.Save(&inventory).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(req).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *ProductRepository) GetOrder(ctx context.Context, orderID int) (models.Order, error) {
	var order models.Order
	if err := r.DB.WithContext(ctx).Preload("OrderItems").Preload("OrderItems.Product").Where("id = ?", orderID).First(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}
