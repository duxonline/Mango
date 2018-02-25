package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {

	dir, dirErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if dirErr != nil {
		log.Fatal(dirErr)
	}
	fmt.Println(dir)

	fileErr := godotenv.Load("/.env")
	if fileErr != nil {
		// log.Fatal("Error loading .env file")
		fmt.Println(fileErr)
	}
}
