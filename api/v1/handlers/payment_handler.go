package handlers

import (
	"net/http"
	"subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
	"subscription-billing-system/middleware"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type PaymentAPI struct {
	paymentUseCase models.PaymentUseCase
}

func NewPaymentAPI(paymentUseCase models.PaymentUseCase) *PaymentAPI {
	return &PaymentAPI{paymentUseCase}
}

func (p *PaymentAPI) CreatePayment(c *gin.Context) {
	var req struct {
		SubscriptionID uuid.UUID `json:"SubscriptionID"`
		Amount         float64 `json:"Amount"`
		PaymentMethod  string `json:"PaymentMethod"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Unauthorized"))
		return
	}

	payment, err := p.paymentUseCase.ProcessPayment(userID, req.SubscriptionID, req.Amount, req.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to prosses payment"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(payment, "Payment proccessed successfully"))
}

func (p *PaymentAPI) GetPayments(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Unauthorized"))
		return
	}

	payments, err := p.paymentUseCase.GetPayments(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive payment"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(payments, "Payment retrive successfully"))
}

func (p *PaymentAPI) GetAllPaymentDetails(c *gin.Context) {

	payments, err := p.paymentUseCase.GetAllPaymentDetails()
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive payment"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(payments, "Payment retrive successfully"))
}

func (uc *PaymentAPI) UpdatePaymentStatus(c *gin.Context) {
	var req struct {
		Id string `json:"id"`
		Status string `json:"status"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	paymentID, err := uuid.Parse(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid user ID format"))
		return
	}

	isUpdated, err := uc.paymentUseCase.UpdatePaymentStatus(paymentID, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed update payment status"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(isUpdated, "Payment status updated successfully"))
}