																																				package models

import (
  	"testing"
  	"time"

  	"github.com/stretchr/testify/assert"
  )

func TestTask_IsOverdue(t *testing.T) {
  	past := time.Now().Add(-1 * time.Hour)
  	future := time.Now().Add(1 * time.Hour)

  	tests := []struct {
      		name     string
      		task     Task
      		expected bool
      	}{
      		{"no due date", Task{Status: StatusPending}, false},
      		{"future due date", Task{Status: StatusPending, DueDate: &future}, false},
      		{"past completed", Task{Status: StatusCompleted, DueDate: &past}, false},
      		{"past pending", Task{Status: StatusPending, DueDate: &past}, true},
      	}

  	for _, tc := range tests {
      		t.Run(tc.name, func(t *testing.T) {
            			assert.Equal(t, tc.expected, tc.task.IsOverdue())
            		})
      	}
  }

func TestTask_MarkCompleted(t *testing.T) {
  	task := Task{Status: StatusPending}
  	task.MarkCompleted()
  	assert.Equal(t, StatusCompleted, task.Status)
  	assert.NotNil(t, task.CompletedAt)
  }
