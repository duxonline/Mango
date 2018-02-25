package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func OpenDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSSL)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database " + connectionString)
	}

	return db
}
