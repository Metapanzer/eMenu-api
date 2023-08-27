package controllers

import (
	"eMenu-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllItem godoc
// @Summary Get all Item menu.
// @Description Get a list of item.
// @Tags item
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "item search by name"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "items per page"
// @Param        sortBy  query  string  false  "field to sort by"
// @Param        orderBy  query  string  false  "sort order: 'asc' or 'desc'"
// @Success 200 {object} []models.Item
// @Router /item [get]
func GetAllItem(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var items []models.Item

	nameQuery := ctx.Query("name")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	sortBy := ctx.Query("sortBy")
	orderBy := ctx.DefaultQuery("orderBy", "asc")

	query := db
	if nameQuery != "" {
		query = query.Where("name LIKE ?", "%"+nameQuery+"%")
	}

	if sortBy != "" {
		order := fmt.Sprintf("%s %s", sortBy, orderBy)
		query = query.Order(order)
	}

	var totalItems int64
	query.Model(&models.Item{}).Count(&totalItems)

	offset := (page - 1) * limit
	query.Limit(limit).Offset(offset).Find(&items)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": items})
}

// GetItemByCategory godoc
// @Summary Get all Item on specific Category.
// @Description Get all item by Category ID.
// @Tags Item
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Category ID"
// @Success 200 {object} []models.Item
// @Router /category/{id}/item [get]
func GetItemByCategory(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var item []models.Item

	if err := db.Where(`category_id = ?`, ctx.Param("id")).First(&item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": item})
}

// InsertItem godoc
// @Summary Insert Item Menu.
// @Description Add a new item menu, This endpoint requires an admin role. The role is checked based on the user's token.
// @Tags Item
// @Accept       json
// @Produce      json
// @Param        Body    body     models.ItemInput  true  "the body to add an item menu"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Item
// @Router /item [post]
func InsertItem(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var newItem models.Item
	if err := ctx.ShouldBindJSON(&newItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var isExist models.Item
	if err := db.Where("name =?", newItem.Name).First(&isExist).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Menu already exist"})
		return
	}

	result := db.Create(&newItem)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newItem})
}

// UpdateItem godoc
// @Summary Update Item Menu.
// @Description Update existing item, This endpoint requires an admin role. The role is checked based on the user's token.
// @Tags Item
// @Accept       json
// @Produce      json
// @Param        Body    body     models.ItemInput  true  "the body to update item"
// @Param        id    path     int  true  "Item ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Item
// @Router /item/{id} [patch]
func UpdateItem(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var Item models.Item
	if err := db.Where("id = ?", ctx.Param("id")).First(&Item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.ItemInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isExist models.Item
	if err := db.Where("name =?", input.Name).First(&isExist).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Item already exist"})
		return
	}

	db.Model(&Item).Updates(&input)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": Item})
}

// DeleteItem godoc
// @Summary Delete Item Menu.
// @Description Delete one item, This endpoint requires an admin role. The role is checked based on the user's token.
// @Tags Item
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Item ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Item
// @Router /item/{id} [delete]
func DeleteItem(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var item models.Item
	if err := db.Where(`id = ?`, ctx.Param("id")).First(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Record not found!"})
		return
	}
	db.Delete(&item)
	ctx.JSON(http.StatusOK, gin.H{"status": " success", "message": "Item deleted successfully"})
}
