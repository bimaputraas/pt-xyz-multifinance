package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xyz-multifinance/internal/logic"
	"xyz-multifinance/internal/model"
)

type (
	Controller struct {
		logic *logic.Logic
	}
)

func New(logic *logic.Logic) (*Controller, error) {
	return &Controller{
		logic: logic,
	}, nil
}

func (ctr *Controller) Register(ctx *gin.Context) {
	var (
		payload = model.User{}
	)
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.logic.Register(payload); err != nil {
		code, resp := errLogicHandler(err)
		ctx.JSON(code, resp)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"code":    201,
	})
}

func (ctr *Controller) Login(ctx *gin.Context) {
	var (
		payload = model.User{}
	)
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := ctr.logic.Login(payload)
	if err != nil {
		code, resp := errLogicHandler(err)
		ctx.JSON(code, resp)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    200,
		"token":   jwt.Token,
	})
}

func (ctr *Controller) RegisterInstallment(ctx *gin.Context) {
	var (
		payload    = model.UserDetail{}
		userAny, _ = ctx.Get("user")
	)

	user, ok := userAny.(model.User)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"status":  http.StatusUnauthorized,
			"detail":  "undefined user",
		})
		return
	}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.logic.RegisterInstallment(int(user.ID), payload); err != nil {
		code, resp := errLogicHandler(err)
		ctx.JSON(code, resp)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"code":    201,
	})
}

func (ctr *Controller) NewTransaction(ctx *gin.Context) {
	var (
		payload    = model.Transaction{}
		userAny, _ = ctx.Get("user")
	)

	user, ok := userAny.(model.User)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"status":  http.StatusUnauthorized,
			"detail":  "undefined user",
		})
		return
	}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.logic.NewTransaction(int(user.ID), payload); err != nil {
		code, resp := errLogicHandler(err)
		ctx.JSON(code, resp)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"code":    201,
	})
}

// return http code and json
func errLogicHandler(err error) (int, gin.H) {
	errLogic, ok := logic.ParseError(err)
	if !ok {
		return 500, gin.H{
			"message": "Internal Server Error",
			"details": err.Error(),
			"code":    500,
		}
	}
	switch errLogic.Code() {
	case logic.InvalidArgument:
		return 400, gin.H{
			"message": "Bad Request",
			"details": err.Error(),
			"code":    400,
		}
	case logic.NotFound:
		return 404, gin.H{
			"message": "Not Found",
			"details": err.Error(),
			"code":    404,
		}
	case logic.Illegal:
		return 401, gin.H{
			"message": "Unauthorized",
			"details": err.Error(),
			"code":    401,
		}
	default:
		return 500, gin.H{
			"message": "Internal Server Error",
			"details": err.Error(),
			"code":    500,
		}
	}
}
