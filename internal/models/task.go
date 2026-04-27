package models

import (
  	"time"

  	"github.com/google/uuid"
  	"gorm.io/gorm"
  )

// TaskStatus represents the workflow state of a task.
type TaskStatus string

const (
  	StatusPending    TaskStatus = "pending"
  	StatusInProgress TaskStatus = "in_progress"
  	StatusCompleted  TaskStatus = "completed"
  	StatusArchived   TaskStatus = "archived"
  )

// TaskPriority represents the urgency level of a task.
type TaskPriority string

const (
  	PriorityLow      TaskPriority = "low"
  	PriorityMedium   TaskPriority = "medium"
  	PriorityHigh     TaskPriority = "high"
  	PriorityCritical TaskPriority = "critical"
  )

// Task represents a user-created task or to-do item.
type Task struct {
  	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
  	UserID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
  	Title       string         `gorm:"not null" json:"title"`
  	Description string         `json:"description"`
  	Status      TaskStatus     `gorm:"default:'pending'" json:"status"`
  	Priority    TaskPriority   `gorm:"default:'medium'" json:"priority"`
  	DueDate     *time.Time     `json:"due_date,omitempty"`
  	CompletedAt *time.Time     `json:"completed_at,omitempty"`
  	Tags        []Tag          `gorm:"many2many:task_tags;" json:"tags,omitempty"`
  	CreatedAt   time.Time      `json:"created_at"`
  	UpdatedAt   time.Time      `json:"updated_at"`
  	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
  }

// BeforeCreate ensures a UUID is generated before persisting the task.
func (t *Task) BeforeCreate(tx *gorm.DB) error {
  	if t.ID == uuid.Nil {
      		t.ID = uuid.New()
      	}
  	return nil
  }

// IsOverdue reports whether the task is past its due date and not completed.
func (t *Task) IsOverdue() bool {
  	if t.DueDate == nil || t.Status == StatusCompleted {
      		return false
      	}
  	return time.Now().After(*t.DueDate)
  }

// MarkCompleted updates the task status and timestamp.
func (t *Task) MarkCompleted() {
  	now := time.Now()
  	t.Status = StatusCompleted
  	t.CompletedAt = &now
  }

// TableName overrides the default GORM table name.
func (Task) TableName() string {
  	return "tasks"
  }
