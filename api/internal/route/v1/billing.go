package v1

import (
	"dou-survey/app/controller/billingController"

	"github.com/gin-gonic/gin"
)

func SetupBillingRoute(userBilling *gin.RouterGroup, c billingController.BillingControllerInterface) *gin.RouterGroup {
	userBilling.POST("", c.AddCustomer)

	return userBilling
}
