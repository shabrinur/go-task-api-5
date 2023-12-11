package router

import (
	controller "idstar-idp/rest-api/app/controller/emptraining"
	repository "idstar-idp/rest-api/app/repository/emptraining"
	service "idstar-idp/rest-api/app/service/emptraining"

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
