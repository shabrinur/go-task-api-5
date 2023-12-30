package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"

	"github.com/gin-gonic/gin"
)

func SetChangePasswordRouter(group *gin.RouterGroup, userRepo repository.UserRepository, userMgmtUtil util.UserMgmtUtil) {

	svc := service.NewChangePasswordService(userRepo, userMgmtUtil)
	ctrl := controller.NewChangePasswordController(svc)

	group.POST("/send-otp", ctrl.GetChangePasswordOtp)
	group.POST("/validate-otp", ctrl.ValidateChangePasswordOtp)
	group.POST("/change-password", ctrl.ChangePassword)
}
