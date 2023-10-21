package httphandler

import (
	"activity-reporter/shared/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.userUseCase.Register(c, dto.ConvURegisToModel(req))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h *HttpHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.userUseCase.Login(c, dto.ConvULoginToModel(req))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}
