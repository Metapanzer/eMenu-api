package main

import (
	"eMenu-api/config"
	"eMenu-api/docs"
	"eMenu-api/routes"
	"eMenu-api/utils"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
func main() {
	// for load godotenv
	// for env
	utils.LoadEnv(".env.production")

	//Programmatically set swagger info
	docs.SwaggerInfo.Title = "eMenu API"
	docs.SwaggerInfo.Description = "This is a sample of eMenu API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("API_URL", "fp-sanbercode-go-48-krisna-production.up.railway.app")
	docs.SwaggerInfo.BasePath = "/api/v1"

	//Connect to database
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	route := routes.SetupRouter(db)
	route.Run()
}
