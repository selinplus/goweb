package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/selinplus/goweb/pkg/setting"
	"log"
	"strings"
)

var db *gorm.DB

type Model struct {
}

func Setup() {
	var err error
	var tablePrefix = setting.DatabaseSetting.TablePrefix
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if !db.HasTable(strings.ToUpper(tablePrefix+"_auth")) {
		db.CreateTable(&Auth{})
	}
	if !db.HasTable(strings.ToUpper(tablePrefix+"_tag")) {
		db.CreateTable(&Tag{})
	}
	if !db.HasTable(strings.ToUpper(tablePrefix+"_article")) {
		db.CreateTable(&Article{})
	}
}

func CloseDB() {
	defer db.Close()
}