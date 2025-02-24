package database

import (
	"binvault/pkg/cfg"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	db := ObtainConnection()
	db.AutoMigrate(&File{})
	db.AutoMigrate(&Bucket{})
}

func OpenConnection() *gorm.DB {
	log.Default().Println("Opening connection to database")
	var sqlite = sqlite.Open(cfg.GetVars().DB_PATH)
	db, err := gorm.Open(sqlite)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func ObtainConnection() *gorm.DB {
	if database == nil {
		database = OpenConnection()
	}
	return database
}
