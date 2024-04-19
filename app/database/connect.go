package database

import (
	"fmt"
	"log"
	"strconv"

	"chat/app/config"
	// "chat/app/internal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() {
	var db_error error
	db_port := config.Config("DB_PORT")
	port, db_error := strconv.ParseUint(db_port, 10, 32)

	if db_error != nil {
		log.Println("Fail!")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	DB, db_error = gorm.Open(postgres.Open(dsn))

	if db_error != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// DB.Migrator().DropTable(&model.Note{})

	// DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}
