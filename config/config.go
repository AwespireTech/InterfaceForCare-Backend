package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DATABASE_URL  string
	DATABASE_NAME string
	IPFS_URL      string
	IPFS_USER     string
	IPFS_PASSWORD string
)

func init() {
	var databaseCred string
	if os.Getenv("DATABASE_USERNAME") != "" && os.Getenv("DATABASE_PASSWORD") != "" && os.Getenv("DATABASE_HOST") != "" {
		databaseCred = os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@"
	} else {
		databaseCred = ""
	}
	if os.Getenv("DATABASE_NAME") != "" {
		DATABASE_NAME = os.Getenv("DATABASE_NAME")
	} else {
		DATABASE_NAME = "RiverCare"
	}
	DATABASE_URL = "mongodb://" + databaseCred + os.Getenv("DATABASE_HOST")
	if os.Getenv("IPFS_URL") != "" {
		IPFS_URL = os.Getenv("IPFS_URL")
	} else {
		IPFS_URL = "https://ipfs.infura.io:5001/api/v0/add?recursive=false&cid-version=0&hash=sha2-256"
	}
	IPFS_USER = os.Getenv("IPFS_USER")
	IPFS_PASSWORD = os.Getenv("IPFS_PASSWORD")

}
func PrintConfig() {
	println("DATABASE_URL: " + DATABASE_URL)
}
