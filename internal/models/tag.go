package models

import (
  	"time"

  	"github.com/google/uuid"
  	"gorm.io/gorm"
  )

// Tag represents a label that can be attached to multiple tasks.
type Tag struct {
  	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
  	Name      string         `gorm:"uniqueIndex;not null" json:"name"`
  	Color     string         `gorm:"default:'#6366f1'" json:"color"`
  	CreatedAt time.Time      `json:"created_at"`
  	UpdatedAt time.Time      `json:"updated_at"`
  	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

  	Tasks []Task `gorm:"many2many:task_tags;" json:"-"`
  }

// BeforeCreate sets a UUID before inserting a new tag record.
func (t *Tag) BeforeCreate(tx *gorm.DB) error {
  	if t.ID == uuid.Nil {
      		t.ID = uuid.New()
      	}
  	return nil
  }

// TableName overrides the default GORM table name.
func (Tag) TableName() string {
  	return "tags"
  }
