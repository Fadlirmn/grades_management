package middleware

import(
	"grades-management/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader =="" {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Authorization header required"})
			c.Abort()
			return 
		}
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer"{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token format"})
			c.Abort()
			return 
		}
		tokenStrings := parts[1]
		claims := &utils.Claims{}

		token, err:= jwt.ParseWithClaims(tokenStrings,claims, func(t *jwt.Token) (interface{}, error) {
			return utils.SECRET_KEY, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid or expired token"})
			c.Abort()
			return 
		}
		c.Set("user_id",claims.UserID)
		c.Set("role",claims.Role)

		c.Next()
	} 
	
	
}
