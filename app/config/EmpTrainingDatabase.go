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
	host := getConfigValue("postgres.training.host")
	port := getConfigValue("postgres.training.port")
	dbname := getConfigValue("postgres.training.dbName")
	user := getConfigValue("postgres.training.user")
	password := getConfigValue("postgres.training.password")

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
