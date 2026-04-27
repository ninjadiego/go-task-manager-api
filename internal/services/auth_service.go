// Package services contains the application's core business logic.
package services

import (
  	"errors"
  	"fmt"
  	"time"

  	"github.com/golang-jwt/jwt/v5"
  	"github.com/google/uuid"
  	"github.com/ninjadiego/go-task-manager-api/internal/config"
  	"github.com/ninjadiego/go-task-manager-api/internal/models"
  	"golang.org/x/crypto/bcrypt"
  	"gorm.io/gorm"
  )

var (
  	ErrInvalidCredentials = errors.New("invalid email or password")
  	ErrUserAlreadyExists  = errors.New("user with that email already exists")
  	ErrInvalidToken       = errors.New("invalid or expired token")
  )

// AuthService handles registration, login and JWT issuance.
type AuthService struct {
  	db  *gorm.DB
  	cfg config.JWTConfig
  }

// NewAuthService constructs a new AuthService.
func NewAuthService(db *gorm.DB, cfg config.JWTConfig) *AuthService {
  	return &AuthService{db: db, cfg: cfg}
  }

// TokenPair represents an access and refresh token pair.
type TokenPair struct {
  	AccessToken  string `json:"access_token"`
  	RefreshToken string `json:"refresh_token"`
  	ExpiresIn    int64  `json:"expires_in"`
  }

// Register creates a new user with a hashed password.
func (s *AuthService) Register(email, username, password, fullName string) (*models.User, error) {
  	var existing models.User
  	if err := s.db.Where("email = ?", email).First(&existing).Error; err == nil {
      		return nil, ErrUserAlreadyExists
      	}

  	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  	if err != nil {
      		return nil, fmt.Errorf("hash password: %w", err)
      	}

  	user := &models.User{
      		Email:        email,
      		Username:     username,
      		PasswordHash: string(hash),
      		FullName:     fullName,
      		IsActive:     true,
      		Role:         "user",
      	}

  	if err := s.db.Create(user).Error; err != nil {
      		return nil, fmt.Errorf("create user: %w", err)
      	}
  	return user, nil
  }

// Login validates credentials and returns a token pair.
func (s *AuthService) Login(email, password string) (*models.User, *TokenPair, error) {
  	var user models.User
  	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
      		return nil, nil, ErrInvalidCredentials
      	}

  	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
      		return nil, nil, ErrInvalidCredentials
      	}

  	now := time.Now()
  	user.LastLoginAt = &now
  	s.db.Save(&user)

  	tokens, err := s.generateTokens(user.ID)
  	if err != nil {
      		return nil, nil, err
      	}
  	return &user, tokens, nil
  }

func (s *AuthService) generateTokens(userID uuid.UUID) (*TokenPair, error) {
  	now := time.Now()

  	accessClaims := jwt.MapClaims{
      		"sub": userID.String(),
      		"iss": s.cfg.Issuer,
      		"iat": now.Unix(),
      		"exp": now.Add(s.cfg.AccessTokenTTL).Unix(),
      		"typ": "access",
      	}

  	refreshClaims := jwt.MapClaims{
      		"sub": userID.String(),
      		"iss": s.cfg.Issuer,
      		"iat": now.Unix(),
      		"exp": now.Add(s.cfg.RefreshTokenTTL).Unix(),
      		"typ": "refresh",
      	}

  	secret := []byte(s.cfg.Secret)

  	access, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secret)
  	if err != nil {
      		return nil, err
      	}
  	refresh, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secret)
  	if err != nil {
      		return nil, err
      	}

  	return &TokenPair{
      		AccessToken:  access,
      		RefreshToken: refresh,
      		ExpiresIn:    int64(s.cfg.AccessTokenTTL.Seconds()),
      	}, nil
  }

// ValidateToken parses a JWT and returns the user ID stored in the subject claim.
func (s *AuthService) ValidateToken(tokenString string) (uuid.UUID, error) {
  	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
      		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
            		}
      		return []byte(s.cfg.Secret), nil
      	})
  	if err != nil || !token.Valid {
      		return uuid.Nil, ErrInvalidToken
      	}

  	claims, ok := token.Claims.(jwt.MapClaims)
  	if !ok {
      		return uuid.Nil, ErrInvalidToken
      	}

  	sub, ok := claims["sub"].(string)
  	if !ok {
      		return uuid.Nil, ErrInvalidToken
      	}

  	return uuid.Parse(sub)
  }
