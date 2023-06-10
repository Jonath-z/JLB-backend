package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	DBHost := os.Getenv("POSTGRES_HOST")
	DBUserName := os.Getenv("POSTGRES_USER")
	DBUserPassword := os.Getenv("POSTGRES_PASSWORD")
	DBName := os.Getenv("POSTGRES_DB")
	DBPort := os.Getenv("POSTGRES_PORT")

	fmt.Printf(DBHost, DBUserName, DBUserPassword, DBName, DBPort)

	// var err error

	// dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DBHost, DBUserName, DBUserPassword, DBName, DBPort)
	// db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected Successfully to the Database")
	// DB = db
}
