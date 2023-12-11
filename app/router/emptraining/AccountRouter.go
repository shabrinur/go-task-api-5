package router

import (
	controller "idstar-idp/rest-api/app/controller/emptraining"
	repository "idstar-idp/rest-api/app/repository/emptraining"
	service "idstar-idp/rest-api/app/service/emptraining"

	"github.com/gin-gonic/gin"
)

func SetAccountRouter(group *gin.RouterGroup) {

	repo := repository.NewAccountRepository()
	svc := service.NewAccountService(*repo)
	ctrl := controller.NewAccountController(svc)

	group.POST("/save", ctrl.CreateAccount)
	group.PUT("/update", ctrl.UpdateAccount)
	group.GET("/:id", ctrl.GetAccountById)
	group.GET("/list", ctrl.GetAccountList)
	group.DELETE("/delete", ctrl.DeleteAccount)
}
