package httphandler

import (
	"activity-reporter/usecase"
)

type HttpHandler struct {
	userUseCase    usecase.UserUsecase
	photoUsecase   usecase.PhotoUsecase
	resetPWUsecase usecase.ResetPWUsecase
}

func NewHttpHandler(userUseCase usecase.UserUsecase, photoUsecase usecase.PhotoUsecase, resetPWUsecase usecase.ResetPWUsecase) *HttpHandler {
	return &HttpHandler{
		userUseCase:    userUseCase,
		photoUsecase:   photoUsecase,
		resetPWUsecase: resetPWUsecase,
	}
}
