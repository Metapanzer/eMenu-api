package routes

import (
	"eMenu-api/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	// set db to gin context
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	route.POST("/register", controllers.CreateUser)

	// route.GET("/items", controllers.GetAllMovie)
	// route.POST("/items", controllers.CreateMovie)
	// route.GET("/items/:id", controllers.GetMovieById)
	// route.PATCH("/items/:id", controllers.UpdateMovie)
	// route.DELETE("items/:id", controllers.DeleteMovie)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}
