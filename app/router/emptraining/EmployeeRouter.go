package router

import (
	controller "idstar-idp/rest-api/app/controller/emptraining"
	repository "idstar-idp/rest-api/app/repository/emptraining"
	service "idstar-idp/rest-api/app/service/emptraining"

	"github.com/gin-gonic/gin"
)

func SetEmployeeRouter(group *gin.RouterGroup) {

	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(*repo)
	ctrl := controller.NewEmployeeController(svc)

	group.POST("/save", ctrl.CreateEmployee)
	group.PUT("/update", ctrl.UpdateEmployee)
	group.GET("/:id", ctrl.GetEmployeeById)
	group.GET("/list", ctrl.GetEmployeeList)
	group.DELETE("/delete", ctrl.DeleteEmployee)
}
