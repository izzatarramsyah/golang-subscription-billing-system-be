package middleware

import (
    "github.com/golang-jwt/jwt/v5"
    "subscription-billing-system/models"
    "time"
    "os"
    "fmt"
    "errors"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    UserID uuid.UUID `json:"user_id"`
    jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token
func GenerateJWT(user models.User) (string, error) {
    claims := Claims{
        UserID: user.ID,
        RegisteredClaims: jwt.RegisteredClaims{
            Issuer:    "subscription-billing-system",
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // expires in 24 hours
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
