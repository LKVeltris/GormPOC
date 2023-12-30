package database

import (
	"fmt"
	"myapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres dbname=gormdb password=mysecretpassword port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database!")
	}

	DB = db
	db.AutoMigrate(&models.User{}, &models.Organization{}, &models.Router{}, &models.Billing{})
	fmt.Println("connect to the database!")
}
