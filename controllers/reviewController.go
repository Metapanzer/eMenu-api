package controllers

import (
	"eMenu-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Review added successfully", "data": newReview})
}

// UpdateReview godoc
// @Summary Update user review.
// @Description Update existing user review, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Review
// @Accept       json
// @Produce      json
// @Param        Body    body     models.ReviewUpdateInput  true  "the body to update review"
// @Param        id    path     int  true  "Review ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Review
// @Router /review/{id} [patch]
func UpdateReview(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", ctx.Param("id")).First(&review).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.ReviewUpdateInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&review).Updates(&input)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Review updated successfully", "data": review})
}

// DeleteReview godoc
// @Summary Delete Item review.
// @Description Delete one review, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Review
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Review ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Review
// @Router /review/{id} [delete]
func DeleteReview(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	user_idStr, _ := ctx.Get("user_id")
	user_id, err := uuid.Parse(user_idStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid user ID"})
		return
	}
	var review models.Review
	if err := db.Where(`id = ?`, ctx.Param("id")).First(&review).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Record not found!"})
		return
	}

	if review.UserID != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
		return
	}
	db.Delete(&review)
	ctx.JSON(http.StatusOK, gin.H{"status": " success", "message": "review deleted successfully"})
}

// GetReviewByItem godoc
// @Summary Get all review on specific item.
// @Description Get all review by item ID.
// @Tags Review
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Item ID"
// @Success 200 {object} []models.Review
// @Router /item/{id}/review [get]
func GetReviewByItem(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var review []models.Review

	if err := db.Where(`item_id = ?`, ctx.Param("id")).Find(&review).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": review})
}
