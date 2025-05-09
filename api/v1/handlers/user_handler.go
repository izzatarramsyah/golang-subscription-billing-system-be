// handlers/user_handlers.go
package handlers

import (
	"net/http"
	models "subscription-billing-system/models"
	apiModels "subscription-billing-system/api/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"subscription-billing-system/middleware"
)

type UserAPI struct {
	userUseCase models.UserUseCase
}

func NewUserAPI(userUseCase models.UserUseCase) *UserAPI {
	return &UserAPI{userUseCase}
}

func (uc *UserAPI) GetListUsers(c *gin.Context) {

	users, err := uc.userUseCase.GetListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve users"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(users, "Users retrieved successfully"))
}

func (uc *UserAPI) UpdateUser(c *gin.Context) {
	var user models.User

	// Bind langsung ke struct models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request: "+err.Error()))
		return
	}

	// Konversi ID ke UUID jika perlu
	userID, err := uuid.Parse(user.ID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}
	user.ID = userID

	updatedUser, err := uc.userUseCase.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed to update user"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(updatedUser, "User updated successfully"))
}

func (uc *UserAPI) UpdateUserStatus(c *gin.Context) {
	var req struct {
		Id string `json:"id"`
		Status string `json:"status"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	userID, err := uuid.Parse(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid user ID format"))
		return
	}

	isUpdated, err := uc.userUseCase.UpdateUserStatus(userID, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Failed update user status"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(isUpdated, "User status updated successfully"))
}

func (h *UserAPI) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	userId, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, apiModels.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	user, err := h.userUseCase.FindUserByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve user"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(user, "User retrieved successfully"))
}

func (h *UserAPI) GetUser(c *gin.Context) {
	userId, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, apiModels.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	user, err := h.userUseCase.FindUserByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve user"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(user, "User retrieved successfully"))
}

func (h *UserAPI) GetByRole(c *gin.Context) {
	role := c.Param("role")

	user, err := h.userUseCase.FindUserByRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiModels.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve user"))
		return
	}

	c.JSON(http.StatusOK, apiModels.NewSuccessResponse(user, "User retrieved successfully"))
}
