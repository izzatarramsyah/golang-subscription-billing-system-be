package usecase

import (
	"github.com/google/uuid"
	"subscription-billing-system/models"
)

type ebookUseCase struct {
	repo models.EbookRepository
}

func NewEbookUseCase(repo models.EbookRepository) models.EbookUseCase {
	return &ebookUseCase{repo}
}

func (u *ebookUseCase) UploadEbook(title, filePath string, productId uuid.UUID) error{
    ebook := models.Ebook{
        ProductID: productId,
        Title:    title,
        FilePath: filePath,
    }
   return u.repo.Save(&ebook)
}

func (u *ebookUseCase) ListEbooks() (*[]models.Ebook, error) {
    return u.repo.FindAll()
}

func (u *ebookUseCase) GetEbook(id uuid.UUID) (*models.Ebook, error) {
    return u.repo.FindByProductID(id)
}

func (u *ebookUseCase) UpdateEbook(filePath string, productId uuid.UUID) error{
    ebook, err := u.repo.FindByProductID(productId)
    if err != nil {
        return err 
    }
    ebook.FilePath = filePath

    return u.repo.Update(ebook)
}