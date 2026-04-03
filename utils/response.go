package utils

import(
	"github.com/gin-gonic/gin"
)
type APIResponse struct{
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, code int, status string, message string, data interface{})   {
	c.JSON(code,APIResponse{
		Status: status,
		Message: message,
		Data: data,
	})
}
func SendError(c *gin.Context, code int, message string)  {
	c.AbortWithStatusJSON(code,gin.H{
		"status":"error",
		"messege":message,
	})
}