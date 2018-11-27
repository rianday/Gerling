package test

import (
	"fmt"
	"libraries/lib/system/database"
	"math/rand"

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

func (Mahasiswa) Create() (int64, error) { //(status bool, updeted int) {
	db := database.Open()

	tx := db.Begin()
	defer tx.Close()

	mahasiswi := Mahasiswa{Id: "001", Name: "Marlina", Age: 31, Grade: 8}

	if err := tx.Create(&mahasiswi).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return tx.Commit().RowsAffected, tx.Error
}

func (mahasiswi *Mahasiswa) Update() (int64, error) {
	db := database.Open()

	tx := db.Begin()
	defer tx.Close()

	tx.Where("id = ?", "001").First(&mahasiswi)
	// tx.Model(&mahasiswi).Where("Name = ?", "Marlena")

	//fmt.Println(mahasiswi)
	mahasiswi.Name = "Marlina"
	mahasiswi.Grade = rand.Intn(1000)

	if err := tx.Save(&mahasiswi).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	// return 0, nil

	return tx.RowsAffected, tx.Error
}

func (mahasiswi *Mahasiswa) Delete() (int64, error) {
	db := database.Open()

	tx := db.Begin()
	defer tx.Close()

	tx.Where("id = ?", "001").First(&mahasiswi)

	if err := tx.Delete(&mahasiswi).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return tx.RowsAffected, tx.Error
}

func (Mahasiswa) Read() []Mahasiswa {
	db := database.Open()
	defer db.Close()

	mahasiswi := []Mahasiswa{}

	// err := db.Where("name = ? AND age >= ?", "Marlina", "31").Find(&mahasiswi)
	err := db.Where("age > ?", "10").Find(&mahasiswi)
	if err.Error != nil {
		fmt.Println(err.Error)
		return nil
	}
	return mahasiswi
}
