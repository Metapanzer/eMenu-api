package routes

import (
	"eMenu-api/controllers"
	"eMenu-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	// set db to gin context
	route.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	route.GET("/", Welcome)

	v1 := route.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)

		v1.GET("/category", controllers.GetAllCategory)
		v1.GET("/category/:id/item", controllers.GetItemByCategory)

		v1.GET("/item", controllers.GetAllItem)

		admin := v1.Use(middleware.AdminOnly())
		admin.POST("/category", controllers.InsertCategory)
		admin.PATCH("/category/:id", controllers.UpdateCategory)
		admin.DELETE("/category/:id", controllers.DeleteCategory)

		// route.GET("/items", controllers.GetAllMovie)
		// route.POST("/items", controllers.CreateMovie)
		// route.GET("/items/:id", controllers.GetMovieById)
		// route.PATCH("/items/:id", controllers.UpdateMovie)
		// route.DELETE("items/:id", controllers.DeleteMovie)
	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}

func Welcome(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/swagger/index.html")
}
