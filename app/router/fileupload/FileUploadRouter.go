package router

import (
	controller "idstar-idp/rest-api/app/controller/fileupload"

	"github.com/gin-gonic/gin"
)

func SetFileUploadRouter(group *gin.RouterGroup) {

	ctrl := controller.NewFileUploadController()
	group.POST("/upload", ctrl.UploadFile)
	group.GET("/show/:filename", ctrl.ShowFile)
	group.DELETE("/delete/:filename", ctrl.DeleteFile)
}
