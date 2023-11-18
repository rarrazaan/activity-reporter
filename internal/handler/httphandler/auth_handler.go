package httphandler

import (
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uuc    usecase.UserUsecase
	config dependency.Config
}

func (h AuthHandler) register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.uuc.Register(c, dto.ConvURegisToModel(&req))
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
	res, err := h.uuc.Login(c, dto.ConvULoginToModel(&req))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}

func (h AuthHandler) Route(r *gin.Engine) {
	r.
		Group("/auth").
		POST("/register", h.register).
		POST("/login", h.login)
}

func NewAuthHandler(uuc usecase.UserUsecase, config dependency.Config) *AuthHandler {
	return &AuthHandler{
		uuc:    uuc,
		config: config,
	}
}
