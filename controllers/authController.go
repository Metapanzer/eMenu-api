package controllers

import (
	"eMenu-api/models"
	"eMenu-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body RegisterInput true "the body to register a user"
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(ctx *gin.Context) {
	// Validate input
	var input models.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check if user with same email is exist
	db := ctx.MustGet("db").(*gorm.DB)
	var isExist models.User
	if err := db.Where("email =?", input.Email).First(&isExist).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Account already exist"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Create User
	newUser := models.User{Email: input.Email, Password: hashedPassword, Name: input.Name}
	result := db.Create(&newUser)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	message := "Account created successfully"
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}

// Login godoc
// @Summary User login
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body LoginInput true "the body to login a user"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(ctx *gin.Context) {
	// Validate input
	var input models.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db := ctx.MustGet("db").(*gorm.DB)
	// if err := db.Where("email = ?", input.Email).Take(&user).Error; err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "Account not found"})
	// 	return
	// }

	if result := db.First(&user, "email = ?", strings.ToLower(input.Email)); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Account not found"})
		return
	}

	if err := utils.VerifyPassword(user.Password, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid password"})
		return

	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Account logged in", "token": token})

}
