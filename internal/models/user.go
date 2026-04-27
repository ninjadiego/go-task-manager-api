// Package models defines the domain entities used across the application.
package models

import (
  	"time"

  	"github.com/google/uuid"
  	"gorm.io/gorm"
  )

// User represents an authenticated user of the system.
type User struct {
  	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
  	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
  	Username     string         `gorm:"uniqueIndex;not null" json:"username"`
  	PasswordHash string         `gorm:"not null" json:"-"`
  	FullName     string         `json:"full_name"`
  	AvatarURL    string         `json:"avatar_url,omitempty"`
  	IsActive     bool           `gorm:"default:true" json:"is_active"`
  	IsVerified   bool           `gorm:"default:false" json:"is_verified"`
  	Role         string         `gorm:"default:'user'" json:"role"`
  	LastLoginAt  *time.Time     `json:"last_login_at,omitempty"`
  	CreatedAt    time.Time      `json:"created_at"`
  	UpdatedAt    time.Time      `json:"updated_at"`
  	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

  	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
  }

// BeforeCreate sets a UUID before inserting a new user record.
func (u *User) BeforeCreate(tx *gorm.DB) error {
  	if u.ID == uuid.Nil {
      		u.ID = uuid.New()
      	}
  	return nil
  }

// TableName overrides the default GORM table name.
func (User) TableName() string {
  	return "users"
  }

// PublicView returns a sanitized version of the user safe for API responses.
func (u *User) PublicView() map[string]interface{} {
  	return map[string]interface{}{
      		"id":         u.ID,
      		"email":      u.Email,
      		"username":   u.Username,
      		"full_name":  u.FullName,
      		"avatar_url": u.AvatarURL,
      		"role":       u.Role,
      		"created_at": u.CreatedAt,
      	}
  }
