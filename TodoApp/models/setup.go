package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "log"
    "fmt"
    "github.com/joho/godotenv"

)

var DB *gorm.DB

func ConnectDatabase() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
        os.Getenv("DB_TIMEZONE"),
    )
    
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Post{})

    DB = database
}
