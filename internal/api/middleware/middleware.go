package middleware

import (
	"net/http"
	"strings"
	"time"
	"xyz-multifinance/internal/logic"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type (
	Middleware struct {
		logic *logic.Logic
	}
)

func New(logic *logic.Logic) *Middleware {
	return &Middleware{
		logic: logic,
	}
}

func (m Middleware) Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	})
}

func (m Middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", 1)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"details": "no token",
				"code":    401,
			})
			return
		}

		user, err := m.logic.AuthUser(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"details": err.Error(),
				"code":    401,
			})
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
