package httphandler

import (
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResetPWHandler struct {
	config dependency.Config
	rpuc   usecase.ResetPWUsecase
}

func (h ResetPWHandler) ForgetPW(c *gin.Context) {
	var req dto.ForgetPWReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.rpuc.ForgetPW(c, req.Email)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}

func (h ResetPWHandler) Route(r *gin.Engine) {
	r.POST("/forgot-password", h.ForgetPW)
}

func NewResetPWHandler(config dependency.Config, rpuc usecase.ResetPWUsecase) ResetPWHandler {
	return ResetPWHandler{
		config: config,
		rpuc:   rpuc,
	}
}
