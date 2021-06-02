package db

import (
	"fmt"

	"github.com/Nikhil12894/gqlwithgo/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
	sslmode  string
	TimeZone string
}

var config = dbConfig{"localhost", 5432, "postgres", "gqlwithgo", "root", "disable", "Asia/Shanghai"}

func ConnectDataBase(vender string, dbname string, isProd bool) {
	db, err := GetDatabase()

	if err != nil {
		panic("Failed to connect to database!")
	}
	if !isProd {
		RunMigrations(db)
	}
	DB = db
}

func getDatabaseUrl() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.host, config.user, config.password, config.dbname, config.port, config.sslmode, config.TimeZone)
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDatabaseUrl()), &gorm.Config{})
	return db, err
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Vachil{}, &model.User{}, &model.TotalPrice{}, &model.Booking{})
}

func FindUserById(id int) *model.User {
	var user *model.User
	DB.Where("id = ?", id).First(&user)
	return user
}
func FindVachilById(id int) *model.Vachil {
	var vachil *model.Vachil
	DB.Where("id = ?", id).First(&vachil)
	return vachil
}

func UserPasswordByName(email string) string {
	var hashedPassword string
	DB.Select("SELECT password FROM user").Where("email=?", email).First(&hashedPassword)
	return hashedPassword
}

func UserByEmail(email string) (*model.User, error) {
	var user *model.User
	err := DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
