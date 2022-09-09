package initialize

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"winterchen.com/patient-go/patient/global"
)

func InitDB() {
	db, err := gorm.Open("mysql", global.Configs.Mysql)
	if err != nil {
		global.Log.Error(err.Error())
		panic(err)
	}
	global.DB = db
}
