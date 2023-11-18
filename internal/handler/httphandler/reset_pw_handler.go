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
	esuc   usecase.EmailSenderUsecase
}

func (h ResetPWHandler) ForgetPW(c *gin.Context) {
	var req dto.ForgetPWReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	err := h.esuc.SendEmail("mini-socmed", "test", "akaraiju@gmail.com")
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func (h ResetPWHandler) Route(r *gin.Engine) {
	r.POST("/forgot-password", h.ForgetPW)
}

func NewResetPWHandler(config dependency.Config, rpuc usecase.ResetPWUsecase, esuc usecase.EmailSenderUsecase) ResetPWHandler {
	return ResetPWHandler{
		config: config,
		rpuc:   rpuc,
		esuc:   esuc,
	}
}
