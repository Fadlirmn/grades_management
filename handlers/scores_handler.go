package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
)

type ScoreHandler struct {
	service *services.ScoresService
}

func NewScoreHandler(s *services.ScoresService) *ScoreHandler {
	return &ScoreHandler{service: s}
}

func (h *ScoreHandler) GetScores(c *gin.Context) {
	scores:= h.service.GetScores()

	if scores == nil {
		c.JSON(http.StatusNotFound, gin.H{"messege":"data not found"})
		return
	}

	c.JSON(http.StatusOK, scores)
}
func (h *ScoreHandler) CreateScore(c *gin.Context) {
	var assignment models.Scores

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json: " + err.Error(),
		})
		return
	}

	h.service.CreateScores(assignment)

	c.JSON(http.StatusCreated, gin.H{
		"success": "Score has been created",
	})
}

func (h *ScoreHandler) UpdateScore(c *gin.Context) {

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

	var assignment models.Scores
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	err = h.service.UpdateScores(id,assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Score updated",
	})
}

func (h *ScoreHandler) DeleteScore(c *gin.Context) {

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

	err = h.service.DeleteScores(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Score has been deleted",
	})
}