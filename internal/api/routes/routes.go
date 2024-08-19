package routes

import (
	"github.com/gin-gonic/gin"
	"xyz-multifinance/internal/api/controller"
	"xyz-multifinance/internal/api/middleware"
)

func New(middleware *middleware.Middleware, controller *controller.Controller) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(middleware.Cors())

	v1 := router.Group("/api/v1")
	v1.POST("user/register", controller.Register)
	v1.POST("user/login", controller.Login)
	v1.POST("user/installment", middleware.Auth(), controller.RegisterInstallment)

	v1.PUT("transaction", middleware.Auth(), controller.NewTransaction)

	return router
}
