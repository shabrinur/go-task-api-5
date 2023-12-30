package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/service/usermgmt/helper"

	"github.com/gin-gonic/gin"
)

func SetRegistrationRouter(group *gin.RouterGroup, userHelper helper.UserHelper) {

	svc := service.NewRegistrationService(userHelper)
	ctrl := controller.NewRegistrationController(svc)

	group.POST("", ctrl.RegisterUser)
	group.POST("/send-link", ctrl.GetActivationLink)
	group.GET("/activate", ctrl.ActivateByLink)
	group.POST("/activate", ctrl.ActivateByCode)
}
