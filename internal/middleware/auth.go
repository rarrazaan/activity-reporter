package middleware

import (
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/shared/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(config dependency.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.App.EnvMode == "testing" {
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

		token, err := helper.ValidateAccessToken(splittedHeader[1], config)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidJWTToken.ToErrorDto())
			return
		}

		claims, ok := token.Claims.(*helper.AccessJWTClaim)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidJWTToken.ToErrorDto())
			return
		}

		c.Set("user_id", claims.UserId)

		c.Next()
	}
}
