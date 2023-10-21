package httphandler

import "activity-reporter/usecase"

type HttpHandler struct {
	userUseCase usecase.UserUsecase
}

func NewHttpHandler(userUseCase usecase.UserUsecase) *HttpHandler {
	return &HttpHandler{
		userUseCase: userUseCase,
	}
}
