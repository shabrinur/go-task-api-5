package config

import (
	model "idstar-idp/rest-api/app/model/usermgmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var UserMgmtDB *gorm.DB

func InitUserMgmtDB() {
	host := getConfigValue("postgres.usermanagement.host")
	port := getConfigValue("postgres.usermanagement.port")
	dbname := getConfigValue("postgres.usermanagement.dbName")
	user := getConfigValue("postgres.usermanagement.user")
	password := getConfigValue("postgres.usermanagement.password")

	// create postgres connection
	conn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"
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

	db.AutoMigrate(&model.RoleModel{})
	db.AutoMigrate(&model.ModuleModel{})
	db.AutoMigrate(&model.UserModel{})
	db.AutoMigrate(&model.RoleModuleModel{})

	UserMgmtDB = db
}

func GetUserMgmtDB() *gorm.DB {
	return UserMgmtDB
}
