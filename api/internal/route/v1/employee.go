package v1

import (
	"dou-survey/app/controller/employeeController"

	"github.com/gin-gonic/gin"
)

func SetupEmployeeRoute(employees *gin.RouterGroup, c employeeController.EmployeeControllerInterface) *gin.RouterGroup {
	employees.GET("/info", c.Info)

	return employees
}
