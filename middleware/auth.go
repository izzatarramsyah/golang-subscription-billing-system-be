package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")

        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }

        claims, err := ValidateJWT(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Set user ID & Role from token in context
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}


func RoleAuthorization(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			log.Println("[RoleAuthorization] No role found in context")
			c.JSON(http.StatusForbidden, gin.H{"error": "No role found"})
			c.Abort()
			return
		}

		userRole := role.(string)
		log.Printf("[RoleAuthorization] User role: %s | Allowed roles: %v\n", userRole, allowedRoles)

		for _, allowed := range allowedRoles {
			if userRole == allowed {
				log.Printf("[RoleAuthorization] Access granted for role: %s\n", userRole)
				c.Next()
				return
			}
		}

		log.Printf("[RoleAuthorization] Access denied for role: %s\n", userRole)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
