package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/handler"
)

type StudentRouter struct {
	h *handler.StudentHandler
}

func InitStudentRouter(handler *handler.StudentHandler) *StudentRouter {
	return &StudentRouter{
		h: handler,
	}
}

func (r StudentRouter) SetupRouter(g *gin.Engine) {
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			},
		)
	})

	g.GET("/student", r.h.GetAllStudentHandler)
	g.POST("/student", r.h.CreateStudentHandler)
}
