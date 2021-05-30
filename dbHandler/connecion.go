package db

import (
	"github.com/Nikhil12894/gqlwithgo/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase(vender string, dbname string) {
	database, err := gorm.Open(vender, dbname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	database.AutoMigrate(&model.Vachil{})

	DB = database
}
