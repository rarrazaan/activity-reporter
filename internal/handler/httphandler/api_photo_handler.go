package httphandler

import (
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) PostPhoto(c *gin.Context) {
	var req dto.PhotoReq
	userIDstr := c.Param("id")
	userID, _ := strconv.Atoi(userIDstr)
	userIDJWT := c.GetInt64("user_id")
	if int64(userID) != userIDJWT {
		c.Error(helper.ErrUnauthorizedUser)
		return
	}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Println("ERR", err)
	}

	res, err := h.photoUsecase.PostPhoto(c, dto.ConvPhotoReq(req, int64(userID)))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}
