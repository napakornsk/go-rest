package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/orm/model"
	"github.com/napakornsk/go-rest/service"
)

type StudentHandler struct {
	srv *service.StudentSrv
}

func InitStudentHandler(srv *service.StudentSrv) *StudentHandler {
	return &StudentHandler{
		srv: srv,
	}
}

func (h *StudentHandler) GetAllStudentHandler(c *gin.Context) {
	data, err := h.srv.GetAllStudents()
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

func (h *StudentHandler) CreateStudentHandler(c *gin.Context) {
	model := new(model.CreateStudent)
	err := c.ShouldBindJSON(model)
	if err != nil {
		log.Printf("Failed to bind json: %v\n", err)
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Failed to bind json: " + err.Error(),
			},
		)
		return
	}

	data, err := h.srv.CreateStudents(model.Student)
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
