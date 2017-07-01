package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
// TODO DB設定を外部ファイルに
var db = NewDBConn()

func NewDBConn()(*gorm.DB){
	db, err := gorm.Open("mysql", "root@/dev_kienu?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	return db
}

func GetDBConn()(*gorm.DB){
	return db
}