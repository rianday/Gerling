package database

import (
	"libraries/lib/system/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open() *gorm.DB {
	//fmt.Println(config.Cfg.Recaptcha)
	//os.Exit(1)

	// db, err := gorm.Open("mysql", config.Cfg.Database.MySQLTest.Username+":"+config.Cfg.Database.MySQLTest.Password+"@/"+config.Cfg.Database.MySQLTest.Name+"?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", config.Cfg.Database.MySQL.Username+":"+config.Cfg.Database.MySQL.Password+"@/"+config.Cfg.Database.MySQL.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("cannot connect to db", err.Error())
	}

	return db
}
