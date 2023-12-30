package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"

	"github.com/gin-gonic/gin"
)

func SetRegistrationRouter(group *gin.RouterGroup, userMgmtRepo repository.UserMgmtRepository, userMgmtUtil util.UserMgmtUtil) {

	svc := service.NewRegistrationService(userMgmtRepo, userMgmtUtil)
	ctrl := controller.NewRegistrationController(svc)

	group.POST("", ctrl.CreateUser)
	group.POST("/send-link", ctrl.GetActivationLink)
	group.GET("/activate", ctrl.ActivateByLink)
	group.POST("/activate", ctrl.ActivateByCode)
}
