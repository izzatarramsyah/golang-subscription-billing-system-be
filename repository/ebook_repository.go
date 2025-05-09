package repository

import (
	"subscription-billing-system/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ebookRepository struct {
	db *gorm.DB
}

func NewEbookRepository(db *gorm.DB) models.EbookRepository {
	return &ebookRepository{db}
}

func (r *ebookRepository) Save(eb *models.Ebook) error {
	err := r.db.Create(eb).Error
	return err
}

func (r *ebookRepository) Update(eb *models.Ebook) error {
	err := r.db.Save(eb).Error
	return err
}

func (r *ebookRepository) FindAll() (*[]models.Ebook, error) {
	var ebooks []models.Ebook
	err := r.db.Find(&ebooks).Error
	return &ebooks, err
}

func (r *ebookRepository) FindByID(id uuid.UUID) (*models.Ebook, error) {
	var ebook models.Ebook
	if err := r.db.Where("id = ?", id).Find(&ebook).Error; err != nil {
		return nil, err
	}
	return &ebook, nil
}

func (r *ebookRepository) FindByProductID(productId uuid.UUID) (*models.Ebook, error) {
	var ebook models.Ebook
	if err := r.db.Where("product_id = ?", productId).Find(&ebook).Error; err != nil {
		return nil, err
	}
	return &ebook, nil
}