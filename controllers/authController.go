package controllers

import (
	"eMenu-api/models"
	"eMenu-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateUser godoc
// @Summary Create New User.
// @Description Creating a new User.
// @Tags User
// @Param Body body UserInput true "the body to create a new User"
// @Produce json
// @Success 200 {object} models.User
// @Router /register [post]
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Create User
	newUser := models.User{Email: input.Email, Password: hashedPassword, Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&newUser)

	message := "Account created successfully"
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}
