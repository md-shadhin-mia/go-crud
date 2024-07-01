package models

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
