// handlers/user_handlers.go
package handlers

import (
	"net/http"
	"log"
	"strconv"
	"subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlanAPI struct {
	planUseCase models.PlanUseCase
}

func NewPlanAPI(planUseCase models.PlanUseCase) *PlanAPI {
	return &PlanAPI{planUseCase}
}

func (h *PlanAPI) Create(c *gin.Context) {
	var input models.Plan

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Optional: log ID untuk debug
	log.Println("Received ProductID:", input.ProductID)

	if err := h.planUseCase.CreatePlan(&input); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to create plan"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Plan created successfully"))
}

func (h *PlanAPI) Update(c *gin.Context) {
	var req struct {
		ID             string `json:"ID"`
		ProductID      string `json:"ProductID"`
		Name           string `json:"Name"`
		Price          string `json:"Price"`
		DurationMonths string `json:"DurationMonths"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Request"))
		return
	}

	// Parsing UUIDs
	planID, err := uuid.Parse(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Plan ID: "+err.Error()))
		return
	}

	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Product ID: "+err.Error()))
		return
	}

	// Parsing Price
	price, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Price format: "+err.Error()))
		return
	}

	// Parsing DurationMonths (with base 10)
	durationMonths, err := strconv.ParseInt(req.DurationMonths, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid DurationMonths format: "+err.Error()))
		return
	}

	// Prepare input model
	input := models.Plan{
		ID:             planID,
		ProductID:      productID,
		Name:           req.Name,
		Price:          price,
		DurationMonths: int(durationMonths), // Converting to int (if necessary)
	}

	// Call use case to update plan
	if err := h.planUseCase.UpdatePlan(&input); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to update plan"))
		return
	}

	// Success response
	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Plan updated successfully"))
}

func (h *PlanAPI) List(c *gin.Context) {

	plans, err := h.planUseCase.GetAllPlans()
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive Plans"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(plans, "Plans retrieved successfully"))
}

func (h *PlanAPI) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	planId, err := uuid.Parse(idParam)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Plan ID"))
		return
	}

	plan, err := h.planUseCase.GetPlanByID(planId)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive plan"))
		return
	}
	
	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(plan, "Plan retrive successfully"))
}

func (h *PlanAPI) Delete(c *gin.Context) {
	idParam := c.Param("id")
	log.Println("Received Plan ID:", idParam) // log ini bantu debug

	planId, err := uuid.Parse(idParam)

	if err != nil {
		log.Println("UUID Parse Error:", err) // tambahan log
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Plan ID"))
		return
	}

	if err := h.planUseCase.DeletePlan(planId); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to delete Plan"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Plan deleted successfully"))
}


func (h *PlanAPI) GetByProductID(c *gin.Context) {
	idParam := c.Param("id")
	productId, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Product ID"))
		return
	}

	plan, err := h.planUseCase.GetByProductID(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive plan"))
		return
	}
	
	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(plan, "Plan retrive successfully"))
}