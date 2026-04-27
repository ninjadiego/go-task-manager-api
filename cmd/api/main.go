// Package main is the entry point for the Go Task Manager API.
//
// @title           Go Task Manager API
// @version         1.0
// @description     A production-ready REST API for managing tasks with JWT authentication.
// @termsOfService  http://swagger.io/terms/
//
// @contact.name   ninjadiego
// @contact.url    https://github.com/ninjadiego
//
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
//
// @host      localhost:8080
// @BasePath  /api/v1
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
  	"context"
  	"log"
  	"os"
  	"os/signal"
  	"syscall"
  	"time"

  	"github.com/ninjadiego/go-task-manager-api/internal/config"
  	"github.com/ninjadiego/go-task-manager-api/internal/database"
  	"github.com/ninjadiego/go-task-manager-api/internal/server"
  	"github.com/ninjadiego/go-task-manager-api/pkg/logger"
  )

func main() {
  	// Load configuration
  	cfg, err := config.Load()
  	if err != nil {
      		log.Fatalf("failed to load config: %v", err)
      	}

  	// Initialize logger
  	log := logger.New(cfg.Env)
  	defer log.Sync()

  	log.Info("starting go-task-manager-api", "version", "1.0.0", "env", cfg.Env)

  	// Connect to database
  	db, err := database.Connect(cfg.Database)
  	if err != nil {
      		log.Fatal("failed to connect to database", "error", err)
      	}

  	// Run migrations
  	if err := database.Migrate(db); err != nil {
      		log.Fatal("failed to run migrations", "error", err)
      	}

  	// Build and start the server
  	srv := server.New(cfg, db, log)

  	go func() {
      		if err := srv.Start(); err != nil {
            			log.Fatal("server error", "error", err)
            		}
      	}()

  	// Wait for interrupt signal
  	quit := make(chan os.Signal, 1)
  	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  	<-quit

  	log.Info("shutting down server gracefully...")

  	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  	defer cancel()

  	if err := srv.Shutdown(ctx); err != nil {
      		log.Error("server forced to shutdown", "error", err)
      	}

  	log.Info("server exited cleanly")
  }
