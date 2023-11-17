package middleware

import (
	"mini-socmed/internal/shared/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var mapErrorCode = map[helper.AppErrorCode]int{
	helper.StatusBadRequest:          http.StatusBadRequest,
	helper.StatusForbidden:           http.StatusForbidden,
	helper.StatusUnauthorized:        http.StatusUnauthorized,
	helper.StatusInternalServerError: http.StatusInternalServerError,
}

func GlobalErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *helper.AppError:
				c.AbortWithStatusJSON(mapErrorCode[e.Code], e.ToErrorDto())
			case validator.ValidationErrors:
				c.AbortWithStatusJSON(http.StatusBadRequest, helper.ValidationErrResponse(e))
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": e.Error(),
				})
			}
			c.Abort()
		}
	}
}
