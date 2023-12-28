package config

import (
	model "idstar-idp/rest-api/app/model/emptraining"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var TrainingDB *gorm.DB

func InitTrainingDB() {
	host := GetConfigValue("postgres.training.host")
	port := GetConfigValue("postgres.training.port")
	dbname := GetConfigValue("postgres.training.dbName")
	user := GetConfigValue("postgres.training.user")
	password := GetConfigValue("postgres.training.password")

	// create postgres connection
	conn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("Error when opening connection to database: ", err.Error())
		panic(err)
	}

	db.AutoMigrate(&model.TrainingModel{})
	db.AutoMigrate(&model.EmployeeDetailModel{})
	db.AutoMigrate(&model.EmployeeModel{})
	db.AutoMigrate(&model.AccountModel{})
	db.AutoMigrate(&model.EmployeeTrainingModel{})

	TrainingDB = db
}

func GetTrainingDB() *gorm.DB {
	return TrainingDB
}
