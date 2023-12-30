package main

import (
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/middleware"
	"idstar-idp/rest-api/app/migration"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	empTraining "idstar-idp/rest-api/app/router/emptraining"
	fileUpload "idstar-idp/rest-api/app/router/fileupload"
	userMgmt "idstar-idp/rest-api/app/router/usermgmt"
	"idstar-idp/rest-api/app/util"
	"sync"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var once sync.Once

//	@termsOfService				http://swagger.io/terms/
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	//gin.SetMode(gin.ReleaseMode)

	initApp()

	// declare composite util
	userMgmtUtil := util.InitUserMgmtUtil()

	if err := migration.Exec(); err != nil {
		panic(err)
	}

	r := gin.Default()

	// declare middleware
	logger := middleware.LoggerMiddleware{}
	auth := middleware.NewAuthMiddleware()

	r.Use(logger.Logger())

	// declare composite repository
	userMgmtRepo := repository.NewUserMgmtRepository()

	// group routing /user-register
	register := r.Group("/v1/user-register")
	{
		userMgmt.SetRegistrationRouter(register, *userMgmtRepo, *userMgmtUtil)
	}

	// group routing /user-login
	login := r.Group("/v1/user-login")
	{
		userMgmt.SetLoginRouter(login, *userMgmtRepo)
	}

	// group routing /forget-password
	changePwd := r.Group("/v1/forget-password")
	{
		userMgmt.SetChangePasswordRouter(changePwd, userMgmtRepo.UserRepository, *userMgmtUtil)
	}

	// group routing /v1/idstar
	idstar := r.Group("/v1/idstar", auth.Authenticate())
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
	}

	// group routing /v1/file
	file := r.Group("/v1/file", auth.Authenticate())
	{
		fileUpload.SetFileUploadRouter(file)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func initApp() {
	once.Do(func() {
		config.LoadConfigFile()
		config.InitTrainingDB()
		config.InitUserMgmtDB()
		config.InitSwagger()
	})
}
