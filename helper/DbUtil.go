package helper

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//InitializeMySQL to OrderDB
func InitializeMySQL() {
	dBConnection, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:abcd1234@tcp(127.0.0.1:3306)/OrderDB?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	db = dBConnection
}

//GetConnection is get MySQL Connection
func GetConnection() *gorm.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}
