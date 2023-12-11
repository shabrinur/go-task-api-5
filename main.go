package main

import (
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/middleware"
	empTraining "idstar-idp/rest-api/app/router/emptraining"
	fileUpload "idstar-idp/rest-api/app/router/fileupload"
	"sync"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var once sync.Once

func main() {
	gin.SetMode(gin.ReleaseMode)

	initApp()

	r := gin.Default()

	// add middleware
	logger := middleware.LoggerMiddleware{}
	r.Use(logger.Logger())

	// group routing /user-register

	// group routing /user-login

	// group routing /forget-password

	// group routing /v1/idstar
	idstar := r.Group("/v1/idstar")
	{
		// group routing /karyawan
		employee := idstar.Group("/karyawan")
		{
			empTraining.SetEmployeeRouter(employee)
		}

		// group routing /rekening
		account := idstar.Group("/rekening")
		{
			empTraining.SetAccountRouter(account)
		}

		// group routing /training
		training := idstar.Group("/training")
		{
			empTraining.SetTrainingRouter(training)
		}

		// group routing /karyawan-training
		employeeTraining := idstar.Group("/karyawan-training")
		{
			empTraining.SetEmployeeTrainingRouter(employeeTraining)
		}

		// group routing /file
		file := idstar.Group("/file")
		{
			fileUpload.SetFileUploadRouter(file)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func initApp() {
	once.Do(func() {
		config.LoadConfigFile()
		config.InitTrainingDB()
		config.InitSwagger()
	})
}
