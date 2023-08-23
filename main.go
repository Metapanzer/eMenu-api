package main

import (
	"eMenu-api/config"
	"eMenu-api/docs"
	"eMenu-api/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
func main() {
	//Programmatically set swagger info
	docs.SwaggerInfo.Title = "eMenu API"
	docs.SwaggerInfo.Description = "This is a sample of eMenu API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//Connect to database
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	route := routes.SetupRouter(db)
	route.Run(":8080")
}
