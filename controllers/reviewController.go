package controllers

import (
	"eMenu-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InsertReview godoc
// @Summary Insert Item Review.
// @Description Add new item review, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Review
// @Accept       json
// @Produce      json
// @Param        Body    body     models.ReviewInput  true  "the body to add a review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Review
// @Router /review [post]
func InsertReview(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var newReview models.Review
	if err := ctx.ShouldBindJSON(&newReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := db.Create(&newReview)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newReview})
}
