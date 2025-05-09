package handlers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"subscription-billing-system/middleware"
	models "subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
)

type SubscriptionAPI struct {
	subscriptionUseCase models.SubscriptionUseCase
}

func NewSubscriptionAPI(subscriptionUseCase models.SubscriptionUseCase) *SubscriptionAPI {
	return &SubscriptionAPI{subscriptionUseCase}
}

func (s *SubscriptionAPI) Subscribe(c *gin.Context) {
	var subscriptionRequest struct {
		PlanID string `json:"plan_id"`
	}

	if err := c.ShouldBindJSON(&subscriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	log.Printf("Received plan_id: %+v", subscriptionRequest.PlanID)

	planUUID, err := uuid.Parse(subscriptionRequest.PlanID)
	if err != nil {
		log.Printf("Failed to parse UUID: %v", err)
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Plan ID format"))
		return
	}
	
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, apiModels.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	subsID, err := s.subscriptionUseCase.Subscribe(userID, planUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to subscribe"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(subsID, "Subscribed successfully"))

}

func (s *SubscriptionAPI) MySubscription(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, apiModels.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	subscriptions, err := s.subscriptionUseCase.MySubscription(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve subscriptions"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(subscriptions, "Subscriptions retrieved successfully"))
}

func (s *SubscriptionAPI) Unsubscribe(c *gin.Context) {
	idParam := c.Param("id")
	subID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid subscription ID"))
		return
	}

	if err := s.subscriptionUseCase.Unsubscribe(subID); err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to unsubscribe"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Unsubscribed successfully"))

}


