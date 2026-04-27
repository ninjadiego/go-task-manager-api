																																																																																package server

import (
  	"context"
  	"net/http"

  	"github.com/gin-gonic/gin"
  	"github.com/ninjadiego/go-task-manager-api/internal/config"
  	"github.com/ninjadiego/go-task-manager-api/internal/handlers"
  	"github.com/ninjadiego/go-task-manager-api/internal/middleware"
  	"github.com/ninjadiego/go-task-manager-api/internal/services"
  	"github.com/ninjadiego/go-task-manager-api/pkg/logger"
  	"gorm.io/gorm"
  )

type Server struct {
  	cfg    *config.Config
  	db     *gorm.DB
  	log    *logger.Logger
  	engine *gin.Engine
  	httpd  *http.Server
  }

func New(cfg *config.Config, db *gorm.DB, log *logger.Logger) *Server {
  	if cfg.Env == "production" {
      		gin.SetMode(gin.ReleaseMode)
      	}

  	engine := gin.New()
  	engine.Use(gin.Recovery())
  	engine.Use(middleware.CORS(nil))

  	authSvc := services.NewAuthService(db, cfg.JWT)
  	authHandler := handlers.NewAuthHandler(authSvc)
  	healthHandler := handlers.NewHealthHandler(db, "1.0.0")

  	engine.GET("/health", healthHandler.Healthz)

  	v1 := engine.Group("/api/v1")
  	auth := v1.Group("/auth")
  	auth.POST("/register", authHandler.Register)
  	auth.POST("/login", authHandler.Login)

  	protected := v1.Group("")
  	protected.Use(middleware.JWTAuth(authSvc))
  	protected.GET("/users/me", func(c *gin.Context) {
      		userID, _ := c.Get(middleware.ContextUserIDKey)
      		c.JSON(http.StatusOK, gin.H{"user_id": userID})
      	})

  	httpd := &http.Server{
      		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
      		Handler:      engine,
      		ReadTimeout:  cfg.Server.ReadTimeout,
      		WriteTimeout: cfg.Server.WriteTimeout,
      	}

  	return &Server{cfg: cfg, db: db, log: log, engine: engine, httpd: httpd}
  }

func (s *Server) Start() error {
  	s.log.Info("http server listening", "addr", s.httpd.Addr)
  	if err := s.httpd.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      		return err
      	}
  	return nil
  }

func (s *Server) Shutdown(ctx context.Context) error {
  	return s.httpd.Shutdown(ctx)
  }
