package setup

import (
	"angular-twitter/cmd/backend/models"
	"angular-twitter/common/config"
	"fmt"
	"github.com/jinzhu/gorm"

	// import mysql driver for qa and production
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConnect creates connect to database and return it.
func DBConnect() *gorm.DB {
	var db *gorm.DB

	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=5432 user=%s dbname=twitter password=%s sslmode=disable", config.DatabaseURL(), config.DatabaseUser(), config.DatabasePassword()))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Subscription{})

	db.Model(&models.Subscription{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	return db
}
