package gorm

import (
	mysqlconfig "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func New() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/drive?charset=utf8&parseTime=True&loc=Local", // data source name
		DSNConfig: &mysqlconfig.Config{
			DBName: "drive",
		},
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("db load err: %v", err)
	}

	return db
}
