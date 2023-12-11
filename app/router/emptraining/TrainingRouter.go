package router

import (
	controller "idstar-idp/rest-api/app/controller/emptraining"
	repository "idstar-idp/rest-api/app/repository/emptraining"
	service "idstar-idp/rest-api/app/service/emptraining"

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
