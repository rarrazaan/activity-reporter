package main

import (
	"activity-reporter/dependency"
	"activity-reporter/handler/httphandler"
	"activity-reporter/httpserver/middleware"
	"activity-reporter/logger"
	"activity-reporter/repository"
	"activity-reporter/shared/helper"
	"activity-reporter/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewRouter(h *httphandler.HttpHandler) *gin.Engine {
	r := gin.New()
	logger := logger.NewLogger()

	r.Use(gin.Recovery())
	r.Use(requestid.New())
	r.Use(middleware.Logger(logger), middleware.WithTimeout, middleware.GlobalErrorMiddleware())

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	user := r.Group("/users", middleware.Auth())
	user.POST("/:id/post", h.PostPhoto)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})

	return r
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
	db := dependency.ConnectDB()

	crypto := helper.NewAppCrypto()
	jwt := helper.NewJwtTokenizer()

	ur := repository.NewUserRepo(db)
	pr := repository.NewPhotoRepo(db)

	uu := usecase.NewUserUsecase(ur, crypto, jwt)
	pu := usecase.NewPhotoUsecase(pr)

	h := httphandler.NewHttpHandler(uu, pu)
	router := NewRouter(h)
	router.ContextWithFallback = true

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("API_PORT")),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
