package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	setDBEnv()
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
		log.Println("Database connection error!")
		log.Fatalf("error message: %s", err)
	} else {
		log.Println("*******************************")
		log.Println("Database connected sucessfully")
		log.Println("*******************************")

		Migration(Db)
		DB = Db
	}

}

// setDBEnv func will set environent variables for local and production environmet
// a more secure approach to be implemented later
func setDBEnv() {
	// Initialize Database configuration names and set in Environment Variables
	const (
		host     = "dpg-cmc4fk6n7f5s7396iru0-a"
		port     = "5432"
		user     = "politicapp_db_user"
		password = "44yTnogDdtKCJhaQv1dLcAlZKUOf26bT"
		dbname   = "politicapp_db"
	)

	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port)
	os.Setenv("DB_USERNAME", user)
	os.Setenv("DB_PASSWORD", password)
	os.Setenv("DB_NAME", dbname)
}
