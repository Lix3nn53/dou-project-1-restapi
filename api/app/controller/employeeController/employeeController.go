package employeeController

import (
	"dou-survey/app/service/employeeService"
	"dou-survey/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

//EmployeeControllerInterface define the employee controller interface methods
type EmployeeControllerInterface interface {
	Info(c *gin.Context)
}

// EmployeeController handles communication with the employee service
type EmployeeController struct {
	service employeeService.EmployeeServiceInterface
	logger  logger.Logger
}

// NewEmployeeController implements the employee controller interface.
func NewEmployeeController(service employeeService.EmployeeServiceInterface, logger logger.Logger) EmployeeControllerInterface {
	return &EmployeeController{
		service,
		logger,
	}
}

// Find implements the method to handle the service to find a employee by the primary key
func (uc *EmployeeController) Info(c *gin.Context) {
	// Employee was added to context in middleware
	employee := c.MustGet("employee")

	c.JSON(http.StatusOK, employee)
}
