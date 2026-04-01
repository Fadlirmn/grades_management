package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
)

type ObjectiveHandler struct {
	service *services.ObjectiveService
}

func NewObjectiveHandler(s *services.ObjectiveService) *ObjectiveHandler {
	return &ObjectiveHandler{service: s}
}

func (h *ObjectiveHandler) GetObjectives(c *gin.Context) {
	var assignment models.Objective
	
	c.JSON(http.StatusOK, assignment)
}
func (h *ObjectiveHandler) CreateObjective(c *gin.Context) {
	var assignment models.Objective

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json: " + err.Error(),
		})
		return
	}

	h.service.CreateObjective(assignment)

	c.JSON(http.StatusCreated, gin.H{
		"success": "Objective has been created",
	})
}

func (h *ObjectiveHandler) UpdateObjective(c *gin.Context) {

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

	var assignment models.Objective
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	err = h.service.UpdateObjective(id,assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Objective updated",
	})
}

func (h *ObjectiveHandler) DeleteObjective(c *gin.Context) {

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

	err = h.service.DeleteObjective(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Objective has been deleted",
	})
}