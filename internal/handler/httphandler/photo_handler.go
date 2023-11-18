package httphandler

import (
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/middleware"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/shared/helper"
	"mini-socmed/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	config dependency.Config
	puc    usecase.PhotoUsecase
}

func (h PostHandler) PostPhoto(c *gin.Context) {
	var req dto.PhotoReq
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.Error(err)
		return
	}
	userIDJWT := c.GetInt64("user_id")
	if int64(userID) != userIDJWT {
		c.Error(helper.ErrUnauthorizedUser)
		return
	}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}
	res, err := h.puc.PostPhoto(c, dto.ConvPhotoReq(&req, int64(userID)))
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

func NewPostHandler(config dependency.Config, puc usecase.PhotoUsecase) PostHandler {
	return PostHandler{
		config: config,
		puc:    puc,
	}
}
