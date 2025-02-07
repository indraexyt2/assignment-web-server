package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price" gorm:"not null" validate:"required"`
	Category    string    `json:"category" gorm:"not null" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`

	Inventory *Inventory `json:"inventory,omitempty" gorm:"foreignKey:ProductID"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Validate() error {
	v := validator.New()
	return v.Struct(p)
}

type Inventory struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID int       `json:"product_id" gorm:"not null" validate:"required"`
	Quantity  int       `json:"quantity" gorm:"not null" validate:"required"`
	Location  string    `json:"location" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (i *Inventory) TableName() string {
	return "inventories"
}

func (i *Inventory) Validate() error {
	v := validator.New()
	return v.Struct(i)
}

type Order struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	OrderItems []OrderItems `json:"order_items,omitempty" gorm:"foreignKey:OrderID" validate:"required"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) Validate() error {
	v := validator.New()
	return v.Struct(o)
}

type OrderItems struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int       `json:"order_id" gorm:"not null"`
	ProductID int       `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
