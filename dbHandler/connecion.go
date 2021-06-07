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

var config = dbConfig{"localhost", 5432, "postgres", "gqlwithgo", "postgres_docker", "disable", "Asia/Shanghai"}

func ConnectDataBase(vender string, dbname string, isProd bool) {
	db, err := GetDatabase(dbname)

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

func GetDatabase(database string) (*gorm.DB, error) {

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
	var hashedPassword model.User
	DB.Where("email=?", email).First(&hashedPassword)
	return hashedPassword.Password
}

func UserByEmail(email string) (*model.User, error) {
	var user *model.User
	err := DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetImage(input string, vachil1 *model.Vachil) string {
	if len(input) != 0 {
		return vachil1.Images
	}
	return input
}
