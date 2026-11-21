package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	service *services.AssignService
}

func NewAssignmentHandler(s *services.AssignService) *AssignmentHandler {
	return &AssignmentHandler{service: s}
}

func (h *AssignmentHandler) GetAssignments(c *gin.Context) {
	assign:= h.service.GetAssignments()

	if assign == nil {
		c.JSON(http.StatusNotFound, gin.H{"messege":"data not found"})
		return
	}

	c.JSON(http.StatusOK, assign)
}
func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	var assignment models.Assignment

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json: " + err.Error(),
		})
		return
	}

	h.service.CreateAssignment(assignment)

	c.JSON(http.StatusCreated, gin.H{
		"success": "Assignment has been created",
	})
}

func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {

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

	var assignment models.Assignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	updatedAt:= time.Now().UTC()
	assignment.UpdateAt = updatedAt
	err = h.service.UpdateAssignment(id,assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Assignment updated",
	})
}

func (h *AssignmentHandler) DeleteAssignment(c *gin.Context) {

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

	err = h.service.DeleteAssignment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Assignment has been deleted",
	})
}