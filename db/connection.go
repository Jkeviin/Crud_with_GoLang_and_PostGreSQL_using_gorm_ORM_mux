package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=127.0.0.1 user=postgres password=Jkeviin__2130 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func DBConnect() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected")
	}

}
