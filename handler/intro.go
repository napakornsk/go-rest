package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/orm/model"
)

func (h *PortfolioHandler) GetIntroHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId64, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		log.Printf("Failed to parse string query: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse string query: %s" + err.Error(),
		})
		return
	}

	userId := uint(userId64)

	data, err := h.srv.GetIntroById(&userId)
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
	err := c.BindJSON(req)
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
