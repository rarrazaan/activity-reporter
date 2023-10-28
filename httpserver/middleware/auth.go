package middleware

import (
	"activity-reporter/shared/helper"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	helper.LoadEnv()
	return func(c *gin.Context) {
		if os.Getenv("APP_ENVIRONMENT") == "testing" {
			c.Set("user_id", int64(1))
			c.Next()
			return
		}

		header := c.GetHeader("Authorization")
		splittedHeader := strings.Split(header, " ")
		if len(splittedHeader) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidAuthHeader.ToErrorDto())
			return
		}

		token, err := helper.ValidateJWT(splittedHeader[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidJWTToken.ToErrorDto())
			return
		}

		claims, ok := token.Claims.(*helper.MyClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidJWTToken.ToErrorDto())
			return
		}

		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
