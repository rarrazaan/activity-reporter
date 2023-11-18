package middleware

import (
	"mini-socmed/internal/constant"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/shared/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(config dependency.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.App.EnvMode == "testing" {
			c.Set(constant.CtxUserId, int64(1))
			c.Next()
			return
		}

		accessTokenStr, err := c.Cookie(constant.AccessTokenCookieName)
		if err != nil {
			e := helper.ErrAccessTokenExpired
			c.AbortWithStatusJSON(mapErrorCode[e.Code], e.ToErrorDto())
			return
		}

		token, err := helper.ValidateAccessToken(accessTokenStr, config)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidJWTToken.ToErrorDto())
			return
		}

		claims, ok := token.Claims.(*helper.AccessJWTClaim)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrInvalidJWTToken.ToErrorDto())
			return
		}

		c.Set(constant.CtxUserId, claims.UserId)

		c.Next()
	}
}
