package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/orm/model"
)

func (h *PortfolioHandler) GetIntroHandler(c *gin.Context) {
	req := new(model.GetByUserId)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON: %s" + err.Error(),
		})
		return
	}

	data, err := h.srv.GetIntro(req.UserId)
	if err != nil {
		log.Printf("Failed to receive data from service: %v\n", err)
		c.IndentedJSON(
			http.StatusInternalServerError,
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
	res, err := h.srv.GetAllIntro()
	if err != nil {
		log.Printf("Failed to receive data from service: %v\n", err)
		c.IndentedJSON(
			http.StatusInternalServerError,
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
			"data":    res,
		},
	)
}

func (h *PortfolioHandler) CreateIntroHandler(c *gin.Context) {
	req := new(model.Intro)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON: %s" + err.Error(),
		})
		return
	}

	res, err := h.srv.CreateIntro(req)
	if err != nil {
		log.Printf("Failed to receive data from service: %v\n", err)
		c.IndentedJSON(
			http.StatusInternalServerError,
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
			"data":    res,
		},
	)
}
