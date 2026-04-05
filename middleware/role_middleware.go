package middleware

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc  {
	return  func (c *gin.Context)  {
		role,exist := c.Get("role")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorization:Role Invalid"})
			c.Abort()
			return 
		}
		userRole := role.(string)
		// semua fitur
		if userRole == "admin"{
			c.Next()
			return 
		}
		isAllowed := false
		for _, r:= range allowedRoles{
			if r == userRole {
				isAllowed=true
				break
			}
		}
		if !isAllowed{
			c.JSON(http.StatusForbidden,gin.H{
				"error":"Forbidden : Access invalid, cant use this features",
			})
			c.Abort()
			return 
		}
		c.Next()
	}

}