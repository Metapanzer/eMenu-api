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

	route.GET("/", Redirect)

	v1 := route.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)

		v1.GET("/category", controllers.GetAllCategory)
		v1.GET("/category/:id/item", controllers.GetItemByCategory)

		v1.GET("/item", controllers.GetAllItem)
		v1.GET("/item/:id/review", controllers.GetReviewByItem)

		user := v1.Use(middleware.JwtAuth())
		user.PATCH("/account/password", controllers.ChangePassword)

		user.GET("/user/:id/order", controllers.GetOrderByUser)
		user.POST("/order", controllers.InsertOrder)
		user.PATCH("/order/:id", controllers.UpdateOrder)
		user.DELETE("/order/:id", controllers.DeleteOrder)

		user.GET("/order/:id/order-detail", controllers.GetOrderDetailByOrder)
		user.POST("/order-detail", controllers.InsertOrderDetail)
		user.PATCH("/order-detail/:id", controllers.UpdateOrderDetail)
		user.DELETE("/order-detail/:id", controllers.DeleteOrderDetail)

		user.POST("/review", controllers.InsertReview)
		user.PATCH("/review/:id", controllers.UpdateReview)
		user.DELETE("/review/:id", controllers.DeleteReview)

		admin := v1.Use(middleware.AdminOnly())
		admin.POST("/category", controllers.InsertCategory)
		admin.PATCH("/category/:id", controllers.UpdateCategory)
		admin.DELETE("/category/:id", controllers.DeleteCategory)
		admin.POST("/item", controllers.InsertItem)
		admin.PATCH("/item/:id", controllers.UpdateItem)
		admin.DELETE("/item/:id", controllers.DeleteItem)

	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}

func Redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/swagger/index.html")
}
