package test

import (
	"libraries/lib/system/database"

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
	db := database.Open()
	defer db.Close()

	mahasiswi := Mahasiswa{Id: "001", Name: "Marlina", Age: 31, Grade: 8}

	db.NewRecord(mahasiswi)
	db.Create(&mahasiswi)

}
