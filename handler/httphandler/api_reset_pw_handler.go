package httphandler

import (
	"activity-reporter/shared/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) ForgetPW(c *gin.Context) {
	var req dto.ForgetPWReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.resetPWUsecase.ForgetPW(c, req.Email)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}
