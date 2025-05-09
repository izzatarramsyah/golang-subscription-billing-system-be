package models

import (
	"github.com/google/uuid"
)

type Ebook struct {
    ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    ProductID uuid.UUID
    Title     string 
    FilePath  string  // Lokasi file di server
}

func (Ebook) TableName() string {
	return "subscription.ebooks"
}   

type EbookRepository interface {
	Save(ebook *Ebook) error
	Update(ebook *Ebook) error
	FindAll() (*[]Ebook, error)
	FindByID(id uuid.UUID) (*Ebook, error)
	FindByProductID(productId uuid.UUID) (*Ebook, error)
}

type EbookUseCase interface {
	UploadEbook(title, filePath string, productId uuid.UUID) error
	UpdateEbook(filePath string, productId uuid.UUID) error
	ListEbooks() (*[]Ebook, error)
	GetEbook(id uuid.UUID) (*Ebook, error)
}
