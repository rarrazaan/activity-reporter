package httpserver

import (
	"context"
	"fmt"
	"log"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/handler/httphandler"
	"mini-socmed/internal/middleware"
	"mini-socmed/internal/repository"
	"mini-socmed/internal/shared/helper"
	"mini-socmed/internal/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	server struct {
		r      *gin.Engine
		v      *validator.Validate
		cfg    dependency.Config
		crypto helper.AppCrypto
		jwt    helper.JwtTokenizer
		repos  repositories
		ucs    usecases
	}
	repositories struct {
		userRepo      repository.UserRepo
		photoRepo     repository.PhotoRepo
		userPhotoRepo repository.UserPhotoRepo
	}
	usecases struct {
		userUsecase    usecase.UserUsecase
		photoUsecase   usecase.PhotoUsecase
		resetPWUsecase usecase.ResetPWUsecase
	}
)

func (s *server) initRepository(db *gorm.DB, re *redis.Client, cfg dependency.Config) {
	s.repos.userRepo = repository.NewUserRepo(db)
	s.repos.photoRepo = repository.NewPhotoRepo(db)
	s.repos.userPhotoRepo = repository.NewUserPhotoRepo(db)
}

func (s *server) initUsecase(rd *redis.Client) {
	s.ucs.userUsecase = usecase.NewUserUsecase(
		s.repos.userRepo,
		s.crypto,
		s.jwt,
	)
	s.ucs.photoUsecase = usecase.NewPhotoUsecase(s.repos.photoRepo)
	s.ucs.resetPWUsecase = usecase.NewResetPWUsecase(rd, s.repos.userRepo)
}

func (s *server) initHTTPHandler(logger dependency.Logger, config dependency.Config) {
	s.r = gin.New()
	s.r.Use(
		gin.Recovery(),
		requestid.New(),
		middleware.Logger(logger),
		middleware.WithTimeout,
		middleware.GlobalErrorMiddleware(),
	)
	httphandler.NewAuthHandler(s.ucs.userUsecase, s.cfg).Route(s.r)
	httphandler.NewPostHandler(s.cfg, s.ucs.photoUsecase).Route(s.r)
	httphandler.NewResetPWHandler(s.cfg, s.ucs.resetPWUsecase).Route(s.r)

	s.r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})
}

func (s *server) startRESTServer(cfg dependency.Config) *http.Server {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Rest.Port),
		Handler: s.r,
	}

	go func() {
		log.Printf("REST server is running on port %d", cfg.Rest.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return &srv
}

func initGracefulShutdown(restSrv *http.Server, cfg dependency.Config) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.App.GracefulTimeout)*time.Second)
	defer cancel()

	// stop resthandler server
	if err := restSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exiting")
}

func InitApp(db *gorm.DB, rc *redis.Client, cfg dependency.Config, logger dependency.Logger) {

	s := server{
		v:   validator.New(),
		cfg: cfg,
	}

	s.initRepository(db, rc, cfg)
	s.initUsecase(rc)
	s.initHTTPHandler(logger, cfg)

	restSrv := s.startRESTServer(cfg)

	initGracefulShutdown(restSrv, cfg)
}
