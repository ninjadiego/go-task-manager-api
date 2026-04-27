// Package config provides centralized application configuration management.
// It loads values from environment variables and configuration files using Viper.
package config

import (
  	"fmt"
  	"strings"
  	"time"

  	"github.com/spf13/viper"
  )

// Config holds all application configuration values.
type Config struct {
  	Env      string         `mapstructure:"ENV"`
  	Server   ServerConfig   `mapstructure:",squash"`
  	Database DatabaseConfig `mapstructure:",squash"`
  	JWT      JWTConfig      `mapstructure:",squash"`
  }

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
  	Host            string        `mapstructure:"SERVER_HOST"`
  	Port            string        `mapstructure:"SERVER_PORT"`
  	ReadTimeout     time.Duration `mapstructure:"SERVER_READ_TIMEOUT"`
  	WriteTimeout    time.Duration `mapstructure:"SERVER_WRITE_TIMEOUT"`
  	ShutdownTimeout time.Duration `mapstructure:"SERVER_SHUTDOWN_TIMEOUT"`
  }

// DatabaseConfig holds database connection configuration.
type DatabaseConfig struct {
  	Host     string `mapstructure:"DB_HOST"`
  	Port     string `mapstructure:"DB_PORT"`
  	User     string `mapstructure:"DB_USER"`
  	Password string `mapstructure:"DB_PASSWORD"`
  	Name     string `mapstructure:"DB_NAME"`
  	SSLMode  string `mapstructure:"DB_SSL_MODE"`
  }

// JWTConfig holds JWT-related configuration.
type JWTConfig struct {
  	Secret           string        `mapstructure:"JWT_SECRET"`
  	AccessTokenTTL   time.Duration `mapstructure:"JWT_ACCESS_TTL"`
  	RefreshTokenTTL  time.Duration `mapstructure:"JWT_REFRESH_TTL"`
  	Issuer           string        `mapstructure:"JWT_ISSUER"`
  }

// Load reads configuration from environment variables and .env file.
func Load() (*Config, error) {
  	viper.SetConfigFile(".env")
  	viper.SetConfigType("env")
  	viper.AutomaticEnv()
  	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

  	setDefaults()

  	_ = viper.ReadInConfig() // .env is optional

  	var cfg Config
  	if err := viper.Unmarshal(&cfg); err != nil {
      		return nil, fmt.Errorf("unmarshal config: %w", err)
      	}

  	if err := cfg.validate(); err != nil {
      		return nil, err
      	}

  	return &cfg, nil
  }

func setDefaults() {
  	viper.SetDefault("ENV", "development")
  	viper.SetDefault("SERVER_HOST", "0.0.0.0")
  	viper.SetDefault("SERVER_PORT", "8080")
  	viper.SetDefault("SERVER_READ_TIMEOUT", "15s")
  	viper.SetDefault("SERVER_WRITE_TIMEOUT", "15s")
  	viper.SetDefault("SERVER_SHUTDOWN_TIMEOUT", "10s")
  	viper.SetDefault("DB_SSL_MODE", "disable")
  	viper.SetDefault("JWT_ACCESS_TTL", "15m")
  	viper.SetDefault("JWT_REFRESH_TTL", "168h")
  	viper.SetDefault("JWT_ISSUER", "go-task-manager-api")
  }

func (c *Config) validate() error {
  	if c.JWT.Secret == "" {
      		return fmt.Errorf("JWT_SECRET is required")
      	}
  	if c.Database.Host == "" {
      		return fmt.Errorf("DB_HOST is required")
      	}
  	return nil
  }

// DSN returns the PostgreSQL connection string.
func (d DatabaseConfig) DSN() string {
  	return fmt.Sprintf(
      		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
      		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
      	)
  }
