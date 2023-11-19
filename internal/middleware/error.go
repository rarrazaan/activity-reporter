package middleware

import (
	"context"
	"errors"
	"log"
	"mini-socmed/internal/shared/dto"
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
		log.Println(err)
		if err != nil {
			switch e := err.Err.(type) {
			case *helper.AppError:
				c.AbortWithStatusJSON(mapErrorCode[e.Code], e.ToErrorDto())
			case validator.ValidationErrors:
				c.AbortWithStatusJSON(http.StatusBadRequest, helper.ValidationErrResponse(e))
			default:
				if errors.Is(err, context.DeadlineExceeded) {
					c.AbortWithStatus(http.StatusRequestTimeout)
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, dto.JSONResponse{
						Message: "internal server error",
					})
				}
			}
			c.Abort()
		}
	}
}
