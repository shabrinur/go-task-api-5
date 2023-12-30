package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"

	"github.com/gin-gonic/gin"
)

func SetLoginRouter(group *gin.RouterGroup, userMgmtRepo repository.UserMgmtRepository) {

	svc := service.NewLoginService(userMgmtRepo)
	ctrl := controller.NewLoginController(svc)

	group.POST("/login", ctrl.UserPassLogin)
}
