package models

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
    ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Name        string
    Description string
	Status		string
    OwnerID     uuid.UUID `gorm:"type:uuid"`
    CreatedAt   time.Time
}

func (Product) TableName() string {
	return "subscription.products"
}

type ProductRepository interface {
	Create(Product *Product) error
	GetAll() (*[]Product, error)
	GetByID(id uuid.UUID) (*Product, error)
	GetByOwnerID(id uuid.UUID) (*[]Product, error)
	Update(Product *Product) error
	Delete(id uint) error
	UpdateStatus(id uuid.UUID, status string) error
}

type ProductUseCase interface {
	CreateProduct(name string, description string, ownerId string) (uuid.UUID, error)
	GetAllProducts() (*[]Product, error)
	GetProductByID(id uuid.UUID) (*Product, error)
	GetProductByOwnerID(id uuid.UUID) (*[]Product, error)
	UpdateProduct(Product *Product) error
	DeleteProduct(id uint) error
	UpdateStatusProduct(id uuid.UUID, status string) error
}