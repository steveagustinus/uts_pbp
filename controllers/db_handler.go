package controllers

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:M4D3ENA@tcp(localhost:3306)/mp_games")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func connectGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:M4D3ENA@tcp(localhost:3306)/mp_games?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
