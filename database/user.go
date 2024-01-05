package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	PhoneNumber int    `gorm:"unique"`
}

func (u *User) AddUser() (*User, error) {
	var err error = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
