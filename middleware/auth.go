package middleware

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Protect () gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
				c.JSON(400, gin.H{
					"success": false,
					"message": "Missing Authorization header",
				})
				c.Abort()
				return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("SecretYouShouldHide"), nil
		})

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		if token.Valid {
			c.Set("user", token.Claims)
			c.Next()
		} else {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
	}
}