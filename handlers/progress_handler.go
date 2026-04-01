package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
)

type ProgressHandler struct {
	service *services.ProgressService
}

func NewProgressHandler(s *services.ProgressService) *ProgressHandler {
	return &ProgressHandler{service: s}
}

func (h *ProgressHandler) GetProgresss(c *gin.Context) {
	var assignment models.Progress
	
	c.JSON(http.StatusOK, assignment)
}
func (h *ProgressHandler) CreateProgress(c *gin.Context) {
	var assignment models.Progress

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json: " + err.Error(),
		})
		return
	}

	h.service.CreateProgress(assignment)

	c.JSON(http.StatusCreated, gin.H{
		"success": "Progress has been created",
	})
}

func (h *ProgressHandler) UpdateProgress(c *gin.Context) {

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

	var assignment models.Progress
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	err = h.service.UpdateProgress(id,assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Progress updated",
	})
}

func (h *ProgressHandler) DeleteProgress(c *gin.Context) {

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

	err = h.service.DeleteProgress(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Progress has been deleted",
	})
}