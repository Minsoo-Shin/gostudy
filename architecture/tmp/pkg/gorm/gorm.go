package gorm

import (
	"fmt"
	"github.com/Minsoo-Shin/go-boilerplate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reflect"
	"time"
)

func New(cfg config.Config) *gorm.DB {
	if reflect.ValueOf(cfg).IsZero() {
		cfg = config.Config{
			Mysql: config.Mysql{
				Host:     "127.0.0.1:3306",
				User:     "user",
				Password: "pass",
				DbName:   "dbname",
				Options: struct {
					MinConnections int `yaml:"minConnections"`
					MaxConnections int `yaml:"maxConnections"`
				}{
					MinConnections: 1,
					MaxConnections: 10,
				},
			},
		}
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.DbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm err: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("err: failed to connecting to mysql: %v", err)
	}
	sqlDB.SetMaxIdleConns(cfg.Options.MaxConnections)
	sqlDB.SetMaxOpenConns(cfg.Options.MaxConnections)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("err: failed to connecting to mysql: %v", err)
	}
	return db
}
