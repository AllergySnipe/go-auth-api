package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AllergySnipe/go-auth-api/initializer"
	"github.com/AllergySnipe/go-auth-api/models"
	"github.com/AllergySnipe/go-auth-api/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RevokeToken(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var revokedToken models.RevokedToken
	revokedToken.Token = tokenString
	revokedToken.RevokedAt = time.Now()

	if err := initializer.DB.Create(&revokedToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not revoke token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token revoked successfully"})
}

func RefreshToken(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	newToken, err := util.GenerateToken(uint(claims["user_id"].(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not refresh token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", newToken, 3600*24*7, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
