package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(s *services.StudentService) *StudentHandler {
	return &StudentHandler{service: s}
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	var assignment models.Student
	
	c.JSON(http.StatusOK, assignment)
}
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var assignment models.Student

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json: " + err.Error(),
		})
		return
	}

	h.service.CreateStudent(assignment)

	c.JSON(http.StatusCreated, gin.H{
		"success": "Student has been created",
	})
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {

	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be a number",
		})
		return
	}

	var assignment models.Student
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	err = h.service.UpdateStudent(id,assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Student updated",
	})
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {

	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be a number",
		})
		return
	}

	err = h.service.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Student has been deleted",
	})
}