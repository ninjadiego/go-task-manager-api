// Package logger provides a thin wrapper around Uber's Zap logger.
package logger

import (
  	"go.uber.org/zap"
  	"go.uber.org/zap/zapcore"
  )

// Logger wraps zap.SugaredLogger to expose a simpler key/value interface.
type Logger struct {
  	*zap.SugaredLogger
  }

// New returns a new Logger configured for the provided environment.
// In production it uses JSON encoding; in development it uses console encoding.
func New(env string) *Logger {
  	var cfg zap.Config

  	if env == "production" {
      		cfg = zap.NewProductionConfig()
      		cfg.EncoderConfig.TimeKey = "timestamp"
      		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
      	} else {
      		cfg = zap.NewDevelopmentConfig()
      		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
      	}

  	base, err := cfg.Build(zap.AddCallerSkip(1))
  	if err != nil {
      		panic(err)
      	}

  	return &Logger{base.Sugar()}
  }

// WithRequestID returns a child logger with the request_id field attached.
func (l *Logger) WithRequestID(requestID string) *Logger {
  	return &Logger{l.With("request_id", requestID)}
  }

// WithUser returns a child logger with the user_id field attached.
func (l *Logger) WithUser(userID string) *Logger {
  	return &Logger{l.With("user_id", userID)}
  }
