package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/orm/model"
)

func (h *PortfolioHandler) GetWorkExperienceHandler(c *gin.Context) {
	req := new(model.GetByUserId)
	err := c.BindJSON(req)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON: %s" + err.Error(),
		})
		return
	}

	res, err := h.srv.GetWorkExperience(req.UserId)
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

func (h *PortfolioHandler) CreateWorkExperienceHandler(c *gin.Context) {
	req := new(model.WorkExperience)
	err := c.BindJSON(req)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON: %s" + err.Error(),
		})
		return
	}

	res, err := h.srv.CreateWorkExperience(req)
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

func (h *PortfolioHandler) UpdateWorkExperienceHandler(c *gin.Context) {
	var req model.WorkExperience

	// Bind and debug the raw JSON
	err := c.BindJSON(&req)
	if err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON: " + err.Error(),
		})
		return
	}

	// Debug the bound struct
	log.Printf("Bound Struct: %+v\n", req)

	// Pass the ID and struct to the service layer
	res, err := h.srv.UpdateWorkExperience(req.UserId, &req)
	if err != nil {
		log.Printf("Service error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update work experience: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update successful",
		"data":    res,
	})
}
