package main

import (
	"log"

	"github.com/md-shadhin-mia/go-crud/initilizer"
	"github.com/md-shadhin-mia/go-crud/models"
)

func init() {
	initilizer.LoadEnv()
	initilizer.DBConnect()
}

func main() {

	if initilizer.DB == nil {
		log.Fatal("DB is nil")
	}
	initilizer.DB.AutoMigrate(&models.User{})
	initilizer.DB.AutoMigrate(&models.Demo{})

}
