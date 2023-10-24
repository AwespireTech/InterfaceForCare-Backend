package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DATABASE_URL string
)

func init() {
	var databaseCred string
	if os.Getenv("DATABASE_USERNAME") != "" && os.Getenv("DATABASE_PASSWORD") != "" && os.Getenv("DATABASE_HOST") != "" {
		databaseCred = os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@"
	} else {
		databaseCred = ""
	}
	DATABASE_URL = "mongodb://" + databaseCred + os.Getenv("DATABASE_HOST")

}
func PrintConfig() {
	println("DATABASE_URL: " + DATABASE_URL)
}
