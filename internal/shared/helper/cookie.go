package helper

import (
	"mini-socmed/internal/constant"
	"mini-socmed/internal/dependency"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookieAuthToken(c *gin.Context, config dependency.Config, accessToken, refreshToken string) {
	accessTokenCookieExp := int(config.Jwt.AccessTokenExpiration) * 60
	refreshTokenCookieExp := int(config.Jwt.RefreshTokenExpiration) * 60

	if config.App.Domain == "localhost" {
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie(constant.AccessTokenCookieName, accessToken, accessTokenCookieExp, "/", config.App.Domain, false, true)
		c.SetCookie(constant.RefreshTokenCookieName, refreshToken, refreshTokenCookieExp, "/", config.App.Domain, false, true)
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(constant.AccessTokenCookieName, accessToken, accessTokenCookieExp, "/", config.App.Domain, true, true)
	c.SetCookie(constant.RefreshTokenCookieName, refreshToken, refreshTokenCookieExp, "/", config.App.Domain, true, true)
}
