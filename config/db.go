package config

import (
	"eMenu-api/models"
	"eMenu-api/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {

	username := utils.Getenv("MYSQLUSER", "root")
	password := utils.Getenv("MYSQLPASSWORD", "1234")
	host := utils.Getenv("MYSQLHOST", "127.0.0.1")
	port := utils.Getenv("MYSQLPORT", "3306")
	database := utils.Getenv("MYSQLDATABASE", "db_emenu")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Item{}, &models.Order{}, &models.Order_detail{}, &models.Review{})
	fmt.Println("Database migrated successfully")
	return db
}
