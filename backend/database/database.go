package database

import (
	"backend/entity"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
	once sync.Once
)

// creating a connection to database
func Connect() *gorm.DB {
	once.Do(func() {
		// signing env items to variables
		// usr := "root"
		// pwd := ""
		// hst := "localhost"
		// dbs := "bakery"
		usr := os.Getenv("MYSQL_USER")
		pwd := os.Getenv("MYSQL_PASSWORD")
		hst := os.Getenv("DB_HOST")
		dbs := os.Getenv("MYSQL_DATABASE")
		// connnecting to database
		lnk := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", usr, pwd, hst, dbs)
		var err error
		conn, err = gorm.Open(mysql.Open(lnk), &gorm.Config{})
		if err != nil {
			panic("Failed to create connection to database")
		}
		// auto migrate schema
		conn.AutoMigrate(&entity.Type{}, &entity.Event{}, &entity.Product{}, &entity.News{})
	})
	return conn
}

// closing a connection to database
func Close() {
	sql, err := conn.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}
	sql.Close()
}
