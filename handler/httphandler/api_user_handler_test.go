package httphandler_test

import (
	"activity-reporter/handler/httphandler"
	"activity-reporter/mocks"
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUserUsecase *mocks.UserUsecase
)

func Setup() {
	mockUserUsecase = new(mocks.UserUsecase)
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func TestHttpHandler_Register(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return 200 when successfully registered", func(t *testing.T) {
		Setup()
		h := httphandler.NewHttpHandler(mockUserUsecase)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := dto.RegisterReq{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}
		userRes := dto.UserRes{
			ID:    1,
			Name:  "Chitanda",
			Email: "chitanda@example.com",
		}
		mockUserUsecase.On("Register", mock.Anything, dto.ConvURegisToModel(req)).Return(userRes, nil)
		expectedResp, _ := json.Marshal(
			dto.JSONResponse{
				Data: dto.UserRes{
					ID:    1,
					Name:  "Chitanda",
					Email: "chitanda@example.com",
				},
			},
		)

		httpReq := httptest.NewRequest("POST", "/register", MakeRequestBody(req))
		c.Request = httpReq
		h.Register(c)

		assert.Equal(http.StatusCreated, w.Result().StatusCode)
		str := strings.TrimSpace(w.Body.String())
		assert.Equal(string(expectedResp), str)
	})

	t.Run("should return 500 when failed to create user", func(t *testing.T) {
		Setup()
		h := httphandler.NewHttpHandler(mockUserUsecase)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := dto.RegisterReq{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}
		userRes := dto.UserRes{}
		mockUserUsecase.On("Register", mock.Anything, dto.ConvURegisToModel(req)).Return(userRes, helper.ErrInternalServer)
		expectedResp, _ := json.Marshal(
			dto.JSONResponse{
				Message: "internal server error",
			},
		)
		httpReq := httptest.NewRequest("POST", "/register", MakeRequestBody(req))
		c.Request = httpReq
		w.Result().StatusCode = http.StatusInternalServerError
		w.Body.WriteString("{\"message\":\"internal server error\"}")
		h.Register(c)

		assert.Equal(http.StatusInternalServerError, w.Result().StatusCode)
		str := strings.TrimSpace(w.Body.String())
		assert.Equal(string(expectedResp), str)
	})
}
