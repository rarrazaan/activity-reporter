package httphandler

import (
	"activity-reporter/usecase"

	"github.com/cloudinary/cloudinary-go/v2"
)

type HttpHandler struct {
	userUseCase    usecase.UserUsecase
	photoUsecase   usecase.PhotoUsecase
	resetPWUsecase usecase.ResetPWUsecase
	cld            *cloudinary.Cloudinary
}

func NewHttpHandler(userUseCase usecase.UserUsecase, photoUsecase usecase.PhotoUsecase, resetPWUsecase usecase.ResetPWUsecase, cld *cloudinary.Cloudinary) *HttpHandler {
	return &HttpHandler{
		userUseCase:    userUseCase,
		photoUsecase:   photoUsecase,
		resetPWUsecase: resetPWUsecase,
		cld:            cld,
	}
}
