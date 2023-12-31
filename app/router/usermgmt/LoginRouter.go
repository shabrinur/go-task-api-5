package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/service/usermgmt/helper"

	"github.com/gin-gonic/gin"
)

func SetLoginRouter(group *gin.RouterGroup, userHelper helper.UserHelper) {

	svc := service.NewLoginService(userHelper)
	ctrl := controller.NewLoginController(svc)

	group.POST("/login", ctrl.UserPassLogin)
	group.GET("/oauth/:provider", ctrl.OauthLogin)
}
