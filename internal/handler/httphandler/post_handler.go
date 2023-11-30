package httphandler

import (
	"fmt"
	"mini-socmed/internal/cons"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/middleware"
	"mini-socmed/internal/model"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/shared/helper"
	"mini-socmed/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	config  dependency.Config
	rString helper.RandomString
	puc     usecase.PhotoUsecase
}

func (h PostHandler) PostPhoto(c *gin.Context) {
	req := new(dto.PhotoReq)
	userID := c.GetInt64("user_id")
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.Error(err)
		return
	}
	post := &model.Photo{
		ID:       fmt.Sprintf(cons.PhotoIDTempate, userID, h.rString.RandStringBytesMaskImprSrcSB()),
		ImageUrl: req.ImageUrl,
		Caption:  req.Caption,
		UserID:   userID,
	}
	res, err := h.puc.PostPhoto(c, post)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h PostHandler) Route(r *gin.Engine) {
	r.
		Group("/post", middleware.Auth(h.config)).
		POST("", h.PostPhoto)
}

func NewPostHandler(config dependency.Config, puc usecase.PhotoUsecase, rString helper.RandomString) PostHandler {
	return PostHandler{
		config:  config,
		puc:     puc,
		rString: rString,
	}
}
