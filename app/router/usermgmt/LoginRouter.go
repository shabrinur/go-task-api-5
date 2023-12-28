package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"

	"github.com/gin-gonic/gin"
)

func SetLoginRouter(group *gin.RouterGroup, pwdUtil util.PasswordUtil) {

	userRepo := repository.NewUserRepository()
	roleModuleRepo := repository.NewRoleModuleRepository()
	svc := service.NewLoginService(*userRepo, *roleModuleRepo, pwdUtil)
	ctrl := controller.NewLoginController(svc)

	group.POST("/login", ctrl.Login)
}
