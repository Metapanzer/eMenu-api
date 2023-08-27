package controllers

import (
	"eMenu-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllCategory godoc
// @Summary Get all Item Category.
// @Description Get a list of Category.
// @Tags Category
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "category search by name"
// @Success 200 {object} []models.Category
// @Router /category [get]
func GetAllCategory(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var category []models.Category
	nameQuery := ctx.Query("name")
	if nameQuery != "" {
		db.Where("name LIKE ?", "%"+nameQuery+"%").Find(&category)
	} else {
		db.Find(&category)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// InsertCategory godoc
// @Summary Insert Item Category.
// @Description Add a new category, This endpoint requires an admin role. The role is checked based on the user's token.
// @Tags Category
// @Accept       json
// @Produce      json
// @Param        Body    body     models.CategoryInput  true  "the body to add a category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Category
// @Router /category [post]
func InsertCategory(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var newCategory models.Category
	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var isExist models.User
	if err := db.Where("name =?", newCategory.Name).First(&isExist).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Category already exist"})
		return
	}

	result := db.Create(&newCategory)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newCategory})
}

// UpdateCategory godoc
// @Summary Update Item Category.
// @Description Update existing category, This endpoint requires an admin role. The role is checked based on the user's token.
// @Tags Category
// @Accept       json
// @Produce      json
// @Param        Body    body     models.CategoryInput  true  "the body to update category"
// @Param        id    path     int  true  "Category ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Category
// @Router /category/{id} [patch]
func UpdateCategory(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", ctx.Param("id")).First(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.CategoryInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isExist models.Category
	if err := db.Where("name = ? AND id != ?", input.Name, ctx.Param("id")).First(&isExist).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Category with the same name already exists"})
		return
	}

	db.Model(&category).Updates(&input)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// DeleteCategory godoc
// @Summary Delete Item Category.
// @Description Delete one category, This endpoint requires an admin role. The role is checked based on the user's token.
// @Tags Category
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Category ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Category
// @Router /category/{id} [delete]
func DeleteCategory(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var catagory models.Category
	if err := db.Where(`id = ?`, ctx.Param("id")).First(&catagory).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Record not found!"})
		return
	}
	db.Delete(&catagory)
	ctx.JSON(http.StatusOK, gin.H{"status": " success", "message": "Category deleted successfully"})
}
