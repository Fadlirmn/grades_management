package handlers

import (
	"grades-management/models"
	"grades-management/services"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler  {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	
	var user models.User

	if err:= c.ShouldBindJSON(&user) ;err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	err := h.service.Register(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"user registred",
	})
}

func (h *AuthHandler) Login(c *gin.Context)  {
	var req struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	

	if err:=c.ShouldBindJSON(&req) ; err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	
	accessToken,refreshToken,err := h.service.Login(req.Username,req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":"invalid login",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"access_token":accessToken,
		"refresh_token":refreshToken,
	})

}
func (h *AuthHandler)Refresh(c *gin.Context)  {
	var payload struct{
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err:=c.ShouldBindJSON(&payload);err != nil {
		c.JSON(400,gin.H{"error":"refresh_token required"})
		return 
	}

	accessToken, err := h.service.RefreshToken(payload.RefreshToken)
	if err != nil {
		c.JSON(401,gin.H{"error":err.Error()})
		return 
	}
	c.JSON(200,gin.H{
		"access_token":accessToken,
	})
} 	