// Package database wires up the GORM connection and migrations.
package database

import (
  	"fmt"
  	"time"

  	"github.com/ninjadiego/go-task-manager-api/internal/config"
  	"github.com/ninjadiego/go-task-manager-api/internal/models"
  	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
  	"gorm.io/gorm/logger"
  )

// Connect opens a new database connection using the provided configuration.
func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
  	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
      		Logger: logger.Default.LogMode(logger.Warn),// Package database wires up the GORM connection and migrations.
      package database

      import (
        	"fmt"
        	"time"

        	"github.com/ninjadiego/go-task-manager-api/internal/config"
        	"github.com/ninjadiego/go-task-manager-api/internal/models"
        	"gorm.io/driver/postgres"
        	"gorm.io/gorm"
        	"gorm.io/gorm/logger"
        )

      // Connect opens a new database connection using the provided configuration.
      func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
        	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
            		Logger: logger.Default.LogMode(logger.Warn),
            	})
        	if err != nil {
            		return nil, fmt.Errorf("open db: %w", err)
            	}

        	sqlDB, err := db.DB()
        	if err != nil {
            		return nil, fmt.Errorf("get sql.DB: %w", err)
            	}

        	sqlDB.SetMaxIdleConns(10)
        	sqlDB.SetMaxOpenConns(100)
        	sqlDB.SetConnMaxLifetime(time.Hour)

        	return db, nil
        }

      // Migrate runs the auto-migrations for all registered domain models.
      func Migrate(db *gorm.DB) error {
        	return db.AutoMigrate(
            		&models.User{},
            		&models.Task{},
            		&models.Tag{},
            	)
        }
      
      	})
  	if err != nil {
      		return nil, fmt.Errorf("open db: %w", err)
      	}

  	sqlDB, err := db.DB()
  	if err != nil {
      		return nil, fmt.Errorf("get sql.DB: %w", err)
      	}

  	sqlDB.SetMaxIdleConns(10)
  	sqlDB.SetMaxOpenConns(100)
  	sqlDB.SetConnMaxLifetime(time.Hour)

  	return db, nil
  }

// Migrate runs the auto-migr
