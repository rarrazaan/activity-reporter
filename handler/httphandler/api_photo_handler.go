package httphandler

import (
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) PostPhoto(c *gin.Context) {
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
	err = c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		return
	}
	dst := filepath.Base(req.Image.Filename)
	err = c.SaveUploadedFile(req.Image, dst)
	if err != nil {
		c.Error(err)
		return
	}
	resp, err := h.cld.Upload.Upload(c, dst, uploader.UploadParams{PublicID: req.Image.Filename})
	if err != nil {
		c.Error(err)
		return
	}
	res, err := h.photoUsecase.PostPhoto(c, dto.ConvPhotoReq(&req, resp.SecureURL, int64(userID)))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}
