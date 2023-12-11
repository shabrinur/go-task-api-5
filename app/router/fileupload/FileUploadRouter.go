package router

import (
	controller "idstar-idp/rest-api/app/controller/fileupload"

	"github.com/gin-gonic/gin"
)

func SetFileUploadRouter(group *gin.RouterGroup) {

	group.POST("/upload", controller.UploadFile)
	group.GET("/show/:filename", controller.ShowFile)
	group.DELETE("/delete/:filename", controller.DeleteFile)
}
