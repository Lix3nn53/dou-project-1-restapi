package v1

import (
	"dou-survey/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupEmployeeRoute(employees *gin.RouterGroup, c controller.EmployeeControllerInterface) *gin.RouterGroup {
	employees.GET("/info", c.Info)

	return employees
}
