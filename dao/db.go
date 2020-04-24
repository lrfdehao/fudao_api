package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DbCli *gorm.DB
)

func InitDb() {
	var err error
	DbCli, err = gorm.Open("mysql", "root:root@/fudao?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DbCli.LogMode(true)
	DbCli.SingularTable(true)
}
