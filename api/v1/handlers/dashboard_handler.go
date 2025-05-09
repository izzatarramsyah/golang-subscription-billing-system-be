package handlers

import (
	"net/http"
	"subscription-billing-system/models"
    apiModels "subscription-billing-system/api/v1/models"

	"github.com/gin-gonic/gin"
)

type DashboardAPI struct {
	useCase models.DashboardUseCase
}

func NewDashboardAPI(useCase models.DashboardUseCase) *DashboardAPI {
	return &DashboardAPI{useCase}
}

func (uc *DashboardAPI) Dashboard(c *gin.Context) {
	data, err := uc.useCase.GetDashboardData()

	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive dashboard info"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(data, "Dashboard info retrive successfully"))
}
