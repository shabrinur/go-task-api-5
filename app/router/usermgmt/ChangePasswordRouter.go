package router

import (
	controller "idstar-idp/rest-api/app/controller/usermgmt"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/service/usermgmt/helper"

	"github.com/gin-gonic/gin"
)

func SetChangePasswordRouter(group *gin.RouterGroup, otpHelper helper.OtpHelper) {

	svc := service.NewChangePasswordService(otpHelper)
	ctrl := controller.NewChangePasswordController(svc)

	group.POST("/send-otp", ctrl.GetChangePasswordOtp)
	group.POST("/validate-otp", ctrl.ValidateChangePasswordOtp)
	group.POST("/change-password", ctrl.ChangePassword)
}
