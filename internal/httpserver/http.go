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
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type (
	server struct {
		r       *gin.Engine
		v       *validator.Validate
		cfg     dependency.Config
		crypto  helper.AppCrypto
		jwt     helper.JwtTokenizer
		rstring helper.RandomString
		uuidGen helper.UuidGenerator
		repos   repositories
		ucs     usecases
	}
	repositories struct {
		userRepo  repository.UserRepo
		photoRepo repository.PhotoRepo
		redisRepo repository.RedisRepo
	}
	usecases struct {
		authUsecase        usecase.AuthUsecase
		photoUsecase       usecase.PhotoUsecase
		resetPWUsecase     usecase.ResetPWUsecase
		emailSenderUsecase usecase.EmailSenderUsecase
	}
)

func (s *server) initRepository(db *gorm.DB, rd *redis.Client, mdb *mongo.Database, cfg dependency.Config) {
	s.repos.userRepo = repository.NewUserRepo(db)
	s.repos.photoRepo = repository.NewPhotoRepo(mdb)
	s.repos.redisRepo = repository.NewRedisRepo(s.cfg, rd)
}

func (s *server) initUsecase(rd *redis.Client) {
	s.ucs.emailSenderUsecase = usecase.NewEmailSenderUsecase(
		s.cfg.Email.SenderName,
		s.cfg.Email.SenderAddress,
		s.cfg.Email.SenderPassword,
	)
	s.ucs.authUsecase = usecase.NewUserUsecase(
		s.repos.userRepo,
		s.repos.redisRepo,
		s.crypto,
		s.jwt,
		s.cfg,
		s.uuidGen,
		s.ucs.emailSenderUsecase,
	)
	s.ucs.photoUsecase = usecase.NewPhotoUsecase(s.repos.photoRepo, s.repos.userRepo)
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
	httphandler.NewAuthHandler(s.ucs.authUsecase, s.cfg).Route(s.r)
	httphandler.NewPostHandler(s.cfg, s.ucs.photoUsecase, s.rstring).Route(s.r)
	httphandler.NewResetPWHandler(s.cfg, s.ucs.resetPWUsecase, s.ucs.emailSenderUsecase).Route(s.r)

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

func InitApp(
	db *gorm.DB,
	rc *redis.Client,
	mdb *mongo.Database,
	cfg dependency.Config,
	logger dependency.Logger,
	crypto helper.AppCrypto,
	jwt helper.JwtTokenizer,
	rstring helper.RandomString,
	uuidGen helper.UuidGenerator,
) {

	s := server{
		v:       validator.New(),
		cfg:     cfg,
		crypto:  crypto,
		jwt:     jwt,
		rstring: rstring,
		uuidGen: uuidGen,
	}

	s.initRepository(db, rc, mdb, cfg)
	s.initUsecase(rc)
	s.initHTTPHandler(logger, cfg)

	restSrv := s.startRESTServer(cfg)

	initGracefulShutdown(restSrv, cfg)
}
