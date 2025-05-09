package usecase

import (
	"subscription-billing-system/models"
	"time"
	"github.com/google/uuid"
)

type productUseCase struct {
    productRepo  models.ProductRepository
}

func NewProductUseCase(productRepo models.ProductRepository) models.ProductUseCase {
    return &productUseCase{productRepo}
}

func (uc *productUseCase) CreateProduct(name string, description string, ownerId string) (uuid.UUID, error) {
	var product models.Product

	product.Name = name
	product.Description = description

	id, err := uuid.Parse(ownerId)
	if err != nil {
		return uuid.Nil, err
	}
	product.OwnerID = id

	product.CreatedAt = time.Now()

	if err := uc.productRepo.Create(&product); err != nil {
		return uuid.Nil, err
	}

	return product.ID, nil
}

func (uc *productUseCase) GetAllProducts() (*[]models.Product, error) {
	return uc.productRepo.GetAll()
}

func (uc *productUseCase) GetProductByID(id uuid.UUID) (*models.Product, error) {
	return uc.productRepo.GetByID(id)
}

func (uc *productUseCase) UpdateProduct(product *models.Product) error {
	return uc.productRepo.Update(product)
}

func (uc *productUseCase) DeleteProduct(id uint) error {
	return uc.productRepo.Delete(id)
}

func (uc *productUseCase) GetProductByOwnerID(id uuid.UUID) (*[]models.Product, error) {
	return uc.productRepo.GetByOwnerID(id)
}

func (uc *productUseCase) UpdateStatusProduct(id uuid.UUID, status string) error {
	return uc.productRepo.UpdateStatus(id, status)
}