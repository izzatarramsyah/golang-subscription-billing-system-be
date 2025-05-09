package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func GetUserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		return uuid.Nil, false
	}

	// Pastikan userID disimpan sebagai string di context
	userIDStr, ok := userIDVal.(string)
	if !ok {
		return uuid.Nil, false
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, false
	}

	return userID, true
}
