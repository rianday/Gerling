package sys

import (
	"libraries/lib/system/database"
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" form:"username" json:"username"` // binding:"required"`
	Email    string `gorm:"not null;unique" form:"email" json:"user"`        // binding:"required"`
	Phone    string `gorm:"not null;unique" form:"phone" json:"phone"`       // binding:"required"`
	Password string `gorm:"not null" form:"password" json:"password"`        // binding:"required"`
	Token    string `json:"token"`
	//PinToken      string
	Salt          string `json:"-"`
	ResetPassword string `json:"-"`
}

func (User) TableName() string {
	return "sys_user_mst"
}

func (User) CreateTable() {
	db := database.Open()

	log.Println("Connection Established")

	//Drops table if already exists
	db.Debug().DropTableIfExists(&User{})

	//Auto create table based on Model
	db.Debug().AutoMigrate(&User{})

	defer db.Close()
}

//be to Add, Update, Active/NonActive, Find by Id

func (user *User) Add() (error, interface{}) {
	db := database.Open()
	db.LogMode(true)

	tx := db.Begin()
	defer tx.Close()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err, nil
	}

	return tx.Commit().Error, user

}

func (user *User) Get() (error, interface{}) {
	db := database.Open()
	//db.LogMode(true)

	defer db.Close()

	if err := db.Where(user).First(&user).Error; err != nil {
		return err, nil
	}

	return nil, user
}
