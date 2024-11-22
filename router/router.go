package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/handler"
)

type PortfolioRouter struct {
	h *handler.PortfolioHandler
}

func InitPortfolioRouter(handler *handler.PortfolioHandler) *PortfolioRouter {
	return &PortfolioRouter{
		h: handler,
	}
}

func (r PortfolioRouter) SetupRouter(g *gin.Engine) {
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			},
		)
	})

	g.GET("/all-intro", r.h.GetAllIntroHandler)
	g.POST("/intro-by-id", r.h.GetIntroHandler)
	g.POST("/intro", r.h.CreateIntroHandler)
	g.POST("/exp", r.h.CreateWorkExperienceHandler)
	g.POST("/exp-by-id", r.h.GetWorkExperienceHandler)
	g.PATCH("/exp-by-id", r.h.UpdateWorkExperienceHandler)
}
