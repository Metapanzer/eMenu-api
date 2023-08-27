package controllers

import (
	"eMenu-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InsertOrderDetail godoc
// @Summary Insert new order detail.
// @Description Add new order detail, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order_detail
// @Accept       json
// @Produce      json
// @Param        Body    body     models.Order_detailInput  true  "the body to add an order detail"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order_detail
// @Router /order-detail [post]
func InsertOrderDetail(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var newOrderDetail models.Order_detail
	if err := ctx.ShouldBindJSON(&newOrderDetail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := db.Create(&newOrderDetail)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Order detail added successfully", "data": newOrderDetail})
}

// UpdateOrderDetail godoc
// @Summary Update user order detail.
// @Description Update existing order detail, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order_detail
// @Accept       json
// @Produce      json
// @Param        Body    body     models.Order_detailUpdateInput  true  "the body to update order detail"
// @Param        id    path     int  true  "Order_detail ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order_detail
// @Router /order-detail/{id} [patch]
func UpdateOrderDetail(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var orderDetail models.Order_detail
	if err := db.Where("id = ?", ctx.Param("id")).First(&orderDetail).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Order_detailUpdateInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&orderDetail).Updates(&input)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Order detail updated successfully", "data": orderDetail})
}

// DeleteOrderDetail godoc
// @Summary Delete user order detail.
// @Description Delete one order detail, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order_detail
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Order_detail ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order_detail
// @Router /order-detail/{id} [delete]
func DeleteOrderDetail(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var orderDetail models.Order_detail
	if err := db.Where(`id = ?`, ctx.Param("id")).First(&orderDetail).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Record not found!"})
		return
	}

	db.Delete(&orderDetail)
	ctx.JSON(http.StatusOK, gin.H{"status": " success", "message": "Order deleted successfully"})
}

// GetOrderDetailByOrder godoc
// @Summary Get all order detail on specific order.
// @Description Get all order detail by order ID, This endpoint requires user role. The role is checked based on the user's token.
// @Tags Order_detail
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Order ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order_detail
// @Router /order/{id}/order-detail [get]
func GetOrderDetailByOrder(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var orderDetail []models.Order_detail

	if err := db.Where(`order_id = ?`, ctx.Param("id")).Find(&orderDetail).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": orderDetail})
}
