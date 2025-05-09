package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
	"subscription-billing-system/utils"
)

type ReminderAPI struct {
	reminderUseCase models.ReminderUseCase
}

func NewReminderAPI(reminderUseCase models.ReminderUseCase) *ReminderAPI {
	return &ReminderAPI{reminderUseCase}
}

func (h *ReminderAPI) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Subscriber ID"))
		return
	}

	reminders, err := h.reminderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive Reminder"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(reminders, "Reminders retrieved successfully"))
}

func (h *ReminderAPI) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid Reminder ID"))
		return
	}

	if err := h.reminderUseCase.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to delete Reminder"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Reminders delete successfully"))
}

func (h *ReminderAPI) Update(c *gin.Context) {
	var reminder models.Reminder
	var req apiModels.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	if err := utils.MapToStruct(req.Data, &reminder); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request data: "+err.Error()))
		return
	}

	if err := h.reminderUseCase.Update(reminder); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to updated reminder"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Reminders update successfully"))
}

func (h *ReminderAPI) Create(c *gin.Context) {
	var createReminderRq struct {
		Type         string `json:"type"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ReminderDate string `json:"reminderDate"`
	}

	if err := c.ShouldBindJSON(&createReminderRq); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request: "+err.Error()))
		return
	}

	if err := h.reminderUseCase.Create(
		createReminderRq.Type,
		createReminderRq.Title,
		createReminderRq.Description,
		createReminderRq.ReminderDate,
	); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to create reminder"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(true, "Reminder created successfully"))
}


func (h *ReminderAPI) GetAll(c *gin.Context) {
	
	reminders, err := h.reminderUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to retrive reminders"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(reminders, "Reminders retrieved successfully"))
}