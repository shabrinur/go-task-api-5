package config

import (
	"idstar-idp/rest-api/app/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	// create postgres connection
	conn := "host=localhost user=training password=training dbname=training port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
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

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
