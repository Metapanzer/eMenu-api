package controllers

import (
	"eMenu-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InsertOrder godoc
// @Summary Insert new order.
// @Description Add new order, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order
// @Accept       json
// @Produce      json
// @Param        Body    body     models.OrderInput  true  "the body to add an order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order
// @Router /order [post]
func InsertOrder(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := db.Create(&newOrder)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Order added successfully", "data": newOrder})
}

// UpdateOrder godoc
// @Summary Update user order.
// @Description Update existing user order, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order
// @Accept       json
// @Produce      json
// @Param        Body    body     models.OrderUpdateInput  true  "the body to update order"
// @Param        id    path     int  true  "Order ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order
// @Router /order/{id} [patch]
func UpdateOrder(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var order models.Order
	if err := db.Where("id = ?", ctx.Param("id")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.OrderUpdateInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&order).Updates(&input)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "order updated successfully", "data": order})
}

// DeleteOrder godoc
// @Summary Delete user order.
// @Description Delete one user order, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Order ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order
// @Router /order/{id} [delete]
func DeleteOrder(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	user_idStr, _ := ctx.Get("user_id")
	user_id, err := uuid.Parse(user_idStr.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid user ID"})
		return
	}
	var order models.Order
	if err := db.Where(`id = ?`, ctx.Param("id")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Record not found!"})
		return
	}

	if order.UserID != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
		return
	}
	db.Delete(&order)
	ctx.JSON(http.StatusOK, gin.H{"status": " success", "message": "Order deleted successfully"})
}

// GetOrderByUser godoc
// @Summary Get all order on specific user.
// @Description Get all order by user ID, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "User ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order
// @Router /user/{id}/order [get]
func GetOrderByUser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var order []models.Order

	if err := db.Where(`user_id = ?`, ctx.Param("id")).Find(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": order})
}
