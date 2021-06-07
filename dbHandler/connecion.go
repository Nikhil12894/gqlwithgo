package db

import (
	"fmt"
	"math"
	"time"

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

var config = dbConfig{"ec2-184-72-236-57.compute-1.amazonaws.com", 5432, "xfqhokcshowjkd", "dbfc7n6f9pt7uk", "b0e7528edabae5cd03687bdd17c70b885828ab249cdbcdbcb9cdb449de3cac84", "disable", "Asia/Shanghai"}

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

	db, err := gorm.Open(postgres.Open("postgres://xfqhokcshowjkd:b0e7528edabae5cd03687bdd17c70b885828ab249cdbcdbcb9cdb449de3cac84@ec2-184-72-236-57.compute-1.amazonaws.com:5432/dbfc7n6f9pt7uk"), &gorm.Config{})
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
		return input
	}
	return vachil1.Images
}

func CreateTotalPrice(startDate time.Time, endDate time.Time, vachilId int) int {
	vachil := FindVachilById(vachilId)
	serviceCharge := 99.49
	if vachil != nil {
		totalHr := endDate.Sub(startDate).Hours()
		totalDays := math.Ceil(totalHr / 24)
		totalCost := (vachil.UnitPrice * totalDays) + serviceCharge
		totalPrice := &model.TotalPrice{
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ServiceCharge: serviceCharge,
			TTLDays:       int(totalDays),
			UnitPrice:     vachil.UnitPrice,
			Price:         totalCost,
		}
		DB.Create(&totalPrice)
		return totalPrice.ID
	}
	return 0
}
