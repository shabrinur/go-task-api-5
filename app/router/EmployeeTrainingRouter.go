package router

import (
	"idstar-idp/rest-api/app/controller"
	"idstar-idp/rest-api/app/repository"
	"idstar-idp/rest-api/app/service"

	"github.com/gin-gonic/gin"
)

func SetEmployeeTrainingRouter(group *gin.RouterGroup) {

	repo := repository.NewEmployeeTrainingRepository()
	svc := service.NewEmployeeTrainingService(*repo)
	ctrl := controller.NewEmployeeTrainingController(svc)

	group.POST("/save", ctrl.CreateEmployeeTraining)
	group.PUT("/update", ctrl.UpdateEmployeeTraining)
	group.GET("/:id", ctrl.GetEmployeeTrainingById)
	group.GET("/list", ctrl.GetEmployeeTrainingList)
	group.DELETE("/delete", ctrl.DeleteEmployeeTraining)
}
