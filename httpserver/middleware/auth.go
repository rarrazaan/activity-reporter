package middleware

import (
	"activity-reporter/shared/helper"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
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
