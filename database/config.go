package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize Database configuration names and set in Environment Variables
func init() {
	const (
		host     = "127.0.0.1"
		port     = "5432"
		user     = "postgres"
		password = "commandme007!@~"
		dbname   = "politicapp_db"
	)

	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port)
	os.Setenv("DB_USERNAME", user)
	os.Setenv("DB_PASSWORD", password)
	os.Setenv("DB_NAME", dbname)
}

var DB *gorm.DB

// ConnectDB func will be called in main to initialize Database
func ConnectDB() {
	// Get Environment variables to initialize Database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUserName, dbPassword, dbName)

	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error!")
		log.Fatalf("error message: %s", err)
	} else {
		fmt.Println("*******************************")
		fmt.Println("Database connected sucessfully")
		fmt.Println("*******************************")
	}

	Migration(Db)

	DB = Db
}
