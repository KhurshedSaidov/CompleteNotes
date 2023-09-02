package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Notes    []Note `gorm:"foreignkey:UserId"`
}

type Note struct {
	gorm.Model
	Content string `json:"content"`
	UserId  uint   `json:"user_Id"`
}
