package handlers

import (
	// "log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	models "subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
	"subscription-billing-system/middleware"
	"subscription-billing-system/utils"
)

type ReportAPI struct {
	usecase models.ReportUseCase
}

func NewReportAPI(usecase models.ReportUseCase) *ReportAPI {
	return &ReportAPI{usecase}
}

func (h *ReportAPI) GetRevenueReport(c *gin.Context) {
	var req apiModels.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request format"))
		return
	}

	var revenueReportRequest struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
	}
	if err := utils.MapToStruct(req.Data, &revenueReportRequest); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request data: "+err.Error()))
		return
	}

	ownerID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, apiModels.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	report, err := h.usecase.GetRevenueReport(ownerID, revenueReportRequest.StartDate, revenueReportRequest.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(report, "Revenue report retrieved successfully"))
}
