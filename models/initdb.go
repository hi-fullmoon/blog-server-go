package models

import (
	"fmt"
	"log"
	"zhengbiwen/blog-server/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// init db
func InitDB() (*gorm.DB, error) {
	sec, err := utils.ConfFile.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section `database`': %v", err)
	}

	dbType := sec.Key("DB_TYPE").String()
	dbName := sec.Key("DB_NAME").String()
	username := sec.Key("USERNAME").String()
	password := sec.Key("PASSWORD").String()
	dbHost := sec.Key("DB_HOST").String()
	logMode, _ := sec.Key("LOG_MODE").Bool()

	db, err = gorm.Open(dbType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, dbHost, dbName))

	if err != nil {
		return nil, err
	}

	db.LogMode(logMode)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{}, &Category{}, &Tag{}, &Article{}, &Session{})

	return db, nil
}
