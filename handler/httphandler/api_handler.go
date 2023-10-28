package httphandler

import "activity-reporter/usecase"

type HttpHandler struct {
	userUseCase  usecase.UserUsecase
	photoUsecase usecase.PhotoUsecase
}

func NewHttpHandler(userUseCase usecase.UserUsecase, photoUsecase usecase.PhotoUsecase) *HttpHandler {
	return &HttpHandler{
		userUseCase:  userUseCase,
		photoUsecase: photoUsecase,
	}
}
