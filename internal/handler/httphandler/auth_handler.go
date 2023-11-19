package httphandler

import (
	"mini-socmed/internal/constant"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/shared/helper"
	"mini-socmed/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auc    usecase.AuthUsecase
	config dependency.Config
}

func (h AuthHandler) register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.auc.Register(c, dto.ConvURegisToModel(&req))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h AuthHandler) login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.auc.Login(c, dto.ConvULoginToModel(&req))
	if err != nil {
		c.Error(err)
		return
	}
	helper.SetCookieAuthToken(c, h.config, res.AToken, res.RToken)
	c.Status(http.StatusOK)
}

func (h AuthHandler) refreshToken(c *gin.Context) {
	rToken, err := c.Cookie(constant.RefreshTokenCookieName)
	if err != nil {
		c.Error(helper.ErrRefreshTokenExpired)
		return
	}

	aToken, err := h.auc.RefreshAccessToken(c, rToken)
	if err != nil {
		c.Error(err)
		return
	}

	aTokenCookieExp := int(h.config.Jwt.AccessTokenExpiration) * 60

	c.SetCookie(constant.AccessTokenCookieName, *aToken, aTokenCookieExp, "/", "", false, true)
	c.Status(http.StatusOK)
}

func (h AuthHandler) Route(r *gin.Engine) {
	r.
		Group("/auth").
		POST("/register", h.register).
		POST("/login", h.login).
		POST("/refresh-token", h.refreshToken)
}

func NewAuthHandler(auc usecase.AuthUsecase, config dependency.Config) *AuthHandler {
	return &AuthHandler{
		auc:    auc,
		config: config,
	}
}
