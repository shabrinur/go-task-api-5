package router

import (
	"idstar-idp/rest-api/app/controller"
	"idstar-idp/rest-api/app/repository"
	"idstar-idp/rest-api/app/service"

	"github.com/gin-gonic/gin"
)

func SetTrainingRouter(group *gin.RouterGroup) {

	repo := repository.NewTrainingRepository()
	svc := service.NewTrainingService(*repo)
	ctrl := controller.NewTrainingController(svc)

	group.POST("/save", ctrl.CreateTraining)
	group.PUT("/update", ctrl.UpdateTraining)
	group.GET("/:id", ctrl.GetTrainingById)
	group.GET("/list", ctrl.GetTrainingList)
	group.DELETE("/delete", ctrl.DeleteTraining)
}
