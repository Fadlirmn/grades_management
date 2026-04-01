package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
)

type SubjectsHandler struct {
	service *services.SubjectService
}

func NewSubjectsHandler(s *services.SubjectService) *SubjectsHandler {
	return &SubjectsHandler{service: s}
}

func (h *SubjectsHandler) GetSubjectss(c *gin.Context) {
	var assignment models.Subject
	
	c.JSON(http.StatusOK, assignment)
}
func (h *SubjectsHandler) CreateSubjects(c *gin.Context) {
	var assignment models.Subject

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json: " + err.Error(),
		})
		return
	}

	h.service.CreateSubject(assignment)

	c.JSON(http.StatusCreated, gin.H{
		"success": "Subjects has been created",
	})
}

func (h *SubjectsHandler) UpdateSubjects(c *gin.Context) {

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

	var assignment models.Subject
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	err = h.service.UpdateSubject(id,assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Subjects updated",
	})
}

func (h *SubjectsHandler) DeleteSubjects(c *gin.Context) {

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

	err = h.service.DeleteSubject(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Subjects has been deleted",
	})
}