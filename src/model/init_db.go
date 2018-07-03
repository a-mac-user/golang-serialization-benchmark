package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init_data() {
	var err error
	db, err = gorm.Open("mysql", "root:777777Mn!@tcp(192.168.0.107:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	if !db.HasTable(&Role{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(&Role{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&UserGroup{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(&UserGroup{}).Error; err != nil {
			panic(err)
		}
	}
}
