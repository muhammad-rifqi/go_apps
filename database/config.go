package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var MYSQL = "root:@tcp(127.0.0.1:3306)/dbgo?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() error {
	var err error
	Database, err = gorm.Open(mysql.Open(MYSQL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database!")

	return nil

}
