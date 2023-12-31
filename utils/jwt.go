package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var TOKEN_SECRET = Getenv("TOKEN_SECRET", "JWT-nuclear-code")

// Function to generate JWT Token
func GenerateToken(user_id interface{}, role string) (string, error) {
	token_expired_in, err := strconv.Atoi(Getenv("TOKEN_EXPIRED_IN", "60"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_expired_in)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(TOKEN_SECRET))

}

// Function to validate JWT Token and extract user_id
func ValidateToken(ctx *gin.Context) (jwt.MapClaims, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims, nil
}

// Function to separate Bearer and token string
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
