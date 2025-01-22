package middleware

import (
	"fmt"
	"net/http"
	"project_restaurant/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" || &clientToken == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No Authentication token provided")})
			c.Abort()
			return
		}
		claims, err := helpers.ValiateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
