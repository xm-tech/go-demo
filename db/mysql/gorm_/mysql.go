package gorm_

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db_ *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Errorf("db conn fail")
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(15)
	db.SingularTable(true)
	db_ = db
}

func GetDb() *gorm.DB {
	return db_
}
