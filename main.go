package main

import (
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/middleware"
	"idstar-idp/rest-api/app/router"
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

	// group routing /v1/idstar
	idstar := r.Group("/v1/idstar")
	{
		// group routing /karyawan
		employee := idstar.Group("/karyawan")
		{
			router.SetEmployeeRouter(employee)
		}

		// group routing /rekening
		account := idstar.Group("/rekening")
		{
			router.SetAccountRouter(account)
		}

		// group routing /training
		training := idstar.Group("/training")
		{
			router.SetTrainingRouter(training)
		}

		// group routing /karyawan-training
		employeeTraining := idstar.Group("/karyawan-training")
		{
			router.SetEmployeeTrainingRouter(employeeTraining)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func initApp() {
	once.Do(func() {
		config.LoadConfigFile()
		config.InitDB()
		config.InitSwagger()
	})
}
