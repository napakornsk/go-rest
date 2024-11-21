package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/orm/model"
	"github.com/napakornsk/go-rest/service"
)

type PortfolioHandler struct {
	srv *service.PortfolioSrv
}

func InitPortfolioHandler(srv *service.PortfolioSrv) *PortfolioHandler {
	return &PortfolioHandler{
		srv: srv,
	}
}

func (h *PortfolioHandler) GetIntroHandler(c *gin.Context)  {
	intro := new(model.Intro)
	err := c.ShouldBindJSON(intro)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON: %s" + err.Error(),
		})
		return
	}

	data, err := h.srv.GetIntro(intro.UserID)
	if err != nil {
		log.Printf("Failed to receive data from service: %v\n", err)
		c.IndentedJSON(
			http.StatusOK,
			gin.H{
				"message": "Failed to receive data from service: " + err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "successful",
			"data":    data,
		},
	)
}

func (h *PortfolioHandler) GetAllIntroHandler(c *gin.Context) {
	data, err := h.srv.GetAllIntro()
	if err != nil {
		log.Printf("Failed to receive data from service: %v\n", err)
		c.IndentedJSON(
			http.StatusOK,
			gin.H{
				"message": "Failed to receive data from service: " + err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "successful",
			"data":    data,
		},
	)
}

func (h *PortfolioHandler) CreateIntroHandler(c *gin.Context) {
	intro := new(model.Intro)
	err := c.ShouldBindJSON(intro)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON: %s" + err.Error(),
		})
		return
	}

	data, err := h.srv.CreateIntro(intro)
	if err != nil {
		log.Printf("Failed to receive data from service: %v\n", err)
		c.IndentedJSON(
			http.StatusOK,
			gin.H{
				"message": "Failed to receive data from service: " + err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "successful",
			"data":    data,
		},
	)
}
