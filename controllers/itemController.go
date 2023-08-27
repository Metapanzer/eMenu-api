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
// @Tags Category
// @Accept       json
// @Produce      json
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
