package models

import (
	"Gin/learnGin/golangDemo/pkg/logging"
	"Gin/learnGin/golangDemo/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func Setup(){
	var err error
	db ,err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		logging.Fatalf("Fail to open the DB: %v, with errors: %v", setting.DatabaseSetting.Host, err)
	} else {
		logging.Info("Connect to DB successful.")
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConn)
	db.DB().SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConn)
}

func CloseDB(){
	defer db.Close()
}


