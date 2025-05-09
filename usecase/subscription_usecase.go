package usecase

import (
	"time"

	"github.com/google/uuid"

	"subscription-billing-system/models"
)

type subscriptionUseCase struct {
    subscriptionRepo models.SubscriptionRepository
    planRepo         models.PlanRepository
    productRepo      models.ProductRepository
    ebookRepo        models.EbookRepository
}

func NewSubscriptionUseCase(
    subscriptionRepo models.SubscriptionRepository,
    planRepo models.PlanRepository,
    productRepo models.ProductRepository,
    ebookRepo models.EbookRepository,
) models.SubscriptionUseCase {
    return &subscriptionUseCase{
        subscriptionRepo: subscriptionRepo,
        planRepo:         planRepo,
        productRepo:      productRepo,
        ebookRepo:        ebookRepo,
    }
}

func (u *subscriptionUseCase) Subscribe(userID uuid.UUID, planID uuid.UUID) (*uuid.UUID, error) {
	sub := &models.Subscription{
		UserID:    userID,
		PlanID:    planID,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 1, 0),
		Status:    "active",
	}
	
	if err := u.subscriptionRepo.Create(sub); err != nil {
		nilUUID := uuid.Nil
		return &nilUUID, err
	}

	return &sub.ID, nil
}


func (u *subscriptionUseCase) MySubscription(userID uuid.UUID) (*[]models.Subscription, error) {
	return u.subscriptionRepo.GetActiveByUser(userID)
}

func (u *subscriptionUseCase) Unsubscribe(subscriptionID uuid.UUID) error {
	return u.subscriptionRepo.UpdateStatus(subscriptionID, models.SubscriptionStatusCanceled)
}

func (u *subscriptionUseCase) IsSubscriptionActive(userID uuid.UUID) (bool, error) {
    // Ambil array langganan yang aktif untuk pengguna
    subscriptions, err := u.subscriptionRepo.GetActiveByUser(userID)
    if err != nil {
        return false, err
    }

    // Pastikan subscriptions bukan nil dan memiliki data
    if len(*subscriptions) == 0 {
        return false, nil // Tidak ada subscription aktif
    }

    // Periksa masing-masing langganan
    for _, subscription := range *subscriptions { // Menggunakan dereference di sini
        now := time.Now()

        // Cek apakah langganan masih aktif
        if subscription.Status == "active" && subscription.StartDate.Before(now) && subscription.EndDate.After(now) {
            // Ambil plan_id dari langganan untuk mencari produk
            plan, err := u.planRepo.GetByID(subscription.PlanID)
            if err != nil {
                return false, err
            }

            // Dari plan, ambil product_id dan temukan ebook terkait
            product, err := u.productRepo.GetByID(plan.ProductID)
            if err != nil {
                return false, err
            }

            // Ambil ebook berdasarkan product_id
            ebook, err := u.ebookRepo.FindByProductID(product.ID)
            if err != nil {
                return false, err
            }

            // Pastikan ebook ditemukan untuk produk
            if ebook != nil {
                return true, nil // Langganan aktif dan ada ebook untuk diakses
            }
        }
    }

    return false, nil // Tidak ada langganan aktif dengan ebook terkait
}

