package handlers

import (
	"net/http"
	models "subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
	"subscription-billing-system/middleware"

	"github.com/gin-gonic/gin"
)

type SubscriberAPI struct {
	usecase models.SubscriberUseCase
}

func NewSubscriberAPI(usecase models.SubscriberUseCase) *SubscriberAPI {
	return &SubscriberAPI{usecase}
}

func (h *SubscriberAPI) GetSubscribersByOwner(c *gin.Context) {

	ownerID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, apiModels.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	subscribers, err := h.usecase.GetSubscribersByOwner(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to fetch subscribers"))
		return
	}

	// Return the response with the data in a generalized format
	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(subscribers, "Subscribers retrieved successfully"))

}
