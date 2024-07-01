package initilizer

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

var DB *gorm.DB

func DBConnect() {
	// connect db with goorm
	var err error
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DB == nil {
		log.Fatal("DB is nil")
	}
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Error; err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Goorm MySQL")
}
