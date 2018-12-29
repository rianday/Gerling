package sys

import (
	"fmt"
	"libraries/lib/system/database"
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Token    string
	//PinToken      string
	Salt          string
	ResetPassword string
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

func (User) Add(params []interface{}) (error, interface{}) {
	db := database.Open()

	tx := db.Begin()
	defer tx.Close()

	users := User{}

	if 0 < len(params) {
		for _, param := range params {
			users = User{Username: "test", Email: "test@gmail.com", Password: "asda12368asd&&^%@&@("}
			fmt.Println(param)
		}
	}

	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		return err, nil
	}

	return tx.Commit().Error, nil

}
