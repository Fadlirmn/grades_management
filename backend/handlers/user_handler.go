package handlers

import (
	"grades-management/models"
	"grades-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	service *services.UserService
}

func NewUserHandler(s *services.UserService)*UserHandler  {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.service.GetUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler)CreateUser(c *gin.Context)  {
	var user models.User
	if err:= c.ShouldBindJSON(&user);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid json"+err.Error()})
		return
	}
	h.service.CreateUser(user)
	c.JSON(http.StatusOK,gin.H{"success":"User has been created"})
}

func (h *UserHandler)UpdateUser(c *gin.Context)  {
	id:= c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid Id"})
		return
	}
	var user models.User
	err:= c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"invalid json"})
	}
	err = h.service.UpdateUser(id,user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
}
func (h *UserHandler)DeleteUser(c *gin.Context)  {
	id := c.Query("id")

	err:= h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"ID not found"})
	}
	c.JSON(http.StatusOK, gin.H{"success":"User has been deleted"})
}