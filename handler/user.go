package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/napakornsk/go-rest/orm/model"
)

func (h *PortfolioHandler) SignupUserHandler(c *gin.Context) {
	// get email and pass off req body
	req := new(model.User)

	if err := c.BindJSON(req); err != nil {
		log.Printf("Error while binding json: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind json from request: " + err.Error(),
		})
		return
	}

	if err := h.v.Validator.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = h.v.GetErrorMessage(fieldError)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  errorMessages,
		})
		return
	}

	// (service) create user
	res, err := h.srv.SignupUser(req)
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

	// respond
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "successful",
			"data":    res,
		},
	)
}

func (h *PortfolioHandler) SigninUserHandler(c *gin.Context) {
	// get email and pass off req body
	req := new(model.User)

	if err := c.BindJSON(req); err != nil {
		log.Printf("Error while binding json: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind json from request: " + err.Error(),
		})
		return
	}

	if err := h.v.Validator.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = h.v.GetErrorMessage(fieldError)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  errorMessages,
		})
		return
	}

	// (service) create user
	res, err := h.srv.SigninUser(req)
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

	// respond
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "successful",
			"data":    res,
		},
	)
}
