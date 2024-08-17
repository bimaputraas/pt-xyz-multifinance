package routes

import (
	"xyz-multifinance/internal/api/middleware"
	"xyz-multifinance/internal/logic"

	"github.com/gin-gonic/gin"
)

func New(middleware *middleware.Middleware, logic *logic.Logic) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.Use(middleware.Cors())

	v1 := router.Group("/api/v1")
	v1.GET("user/register")
	v1.GET("user/login")
	v1.PUT("transaction")

	return router
}
