package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
)

type LibraryItemAPI struct {
	usecase models.LibraryItemUseCase
}

func NewLibraryItemAPI(usecase models.LibraryItemUseCase) *LibraryItemAPI {
	return &LibraryItemAPI{usecase}
}

func (h *LibraryItemAPI) GetByUser(c *gin.Context) {
	
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Unauthorized"))
		return
	}

	userIDStr, ok := userIDRaw.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid User ID"))
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid User ID"))
		return
	}

	libraryItems, err := h.usecase.GetUserLibrary(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive library items"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(libraryItems, "Library item retrive successfully"))

}
