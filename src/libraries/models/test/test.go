package test

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mahasiswa struct {
	//gorm.Model
	Id    string
	Name  string
	Age   int
	Grade int
}

func (Mahasiswa) TableName() string {
	return "tb_student"
}

func (Mahasiswa) Create() { //(status bool, updeted int) {
	db, err := gorm.Open("mysql", "developer:developer@/db_belajar_golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	mahasiswi := Mahasiswa{Id: "001", Name: "Marlina", Age: 31, Grade: 8}

	db.NewRecord(mahasiswi)
	db.Create(&mahasiswi)

}
