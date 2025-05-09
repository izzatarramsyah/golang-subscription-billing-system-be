package repository

import (
	"gorm.io/gorm"
	"subscription-billing-system/models"
	"github.com/google/uuid"
	"log"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) models.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(product *models.Product) error {
	err := r.db.Create(product).Error
	if err != nil {
		log.Printf("Failed to create product: %v\n", err)
	} else {
		log.Printf("Successfully created product with ID: %d\n", product.ID)
	}
	return err
}

func (r *productRepository) GetAll() (*[]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	if err != nil {
		log.Printf("Failed to get all products: %v\n", err)
	} else {
		log.Printf("Successfully retrieved %d products\n", len(products))
	}
	return &products, err
}

func (r *productRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		log.Printf("Failed to get product with ID %d: %v\n", id, err)
	} else {
		log.Printf("Successfully retrieved product with ID: %d\n", id)
	}
	return &product, err
}

func (r *productRepository) Update(product *models.Product) error {
	err := r.db.Save(product).Error
	if err != nil {
		log.Printf("Failed to update product with ID %d: %v\n", product.ID, err)
	} else {
		log.Printf("Successfully updated product with ID: %d\n", product.ID)
	}
	return err
}

func (r *productRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Product{}, id).Error
	if err != nil {
		log.Printf("Failed to delete product with ID %d: %v\n", id, err)
	} else {
		log.Printf("Successfully deleted product with ID: %d\n", id)
	}
	return err
}

func (r *productRepository) GetByOwnerID(ownerID uuid.UUID) (*[]models.Product, error){
	log.Println("Get products:")
	var products []models.Product
	if err := r.db.Where("owner_id = ?", ownerID).Find(&products).Error; err != nil {
		log.Printf("Error get products: %v\n", err)
		return nil, err
	}
	log.Println("get products successfully")
	return &products, nil
}

func (r *productRepository) UpdateStatus(id uuid.UUID, status string) error {
	log.Printf("STATUS: %s\n", status)
	log.Printf("id: %s\n", id)
	err := r.db.Model(&models.Product{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		log.Printf("Failed to update status product with ID %s: %v\n", id.String(), err)
	} else {
		log.Printf("Successfully updated status product with ID: %s\n", id.String())
	}

	return err
}