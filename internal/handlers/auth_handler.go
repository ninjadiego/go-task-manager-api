package handlers

import (
  	"errors"
  	"net/http"

  	"github.com/gin-gonic/gin"
  	"github.com/ninjadiego/go-task-manager-api/internal/services"
  )

type AuthHandler struct {
  	authSvc *services.AuthService
  }

func NewAuthHandler(authSvc *services.AuthService) *AuthHandler {
  	return &AuthHandler{authSvc: authSvc}
  }

type registerRequest struct {
  	Email    string `json:"email" binding:"required,email"`
  	Username string `json:"username" binding:"required,min=3,max=32"`
  	Password string `json:"password" binding:"required,min=8"`
  	FullName string `json:"full_name"`
  }

type loginRequest struct {
  	Email    string `json:"email" binding:"required,email"`
  	Password string `json:"password" binding:"required"`
  }

func (h *AuthHandler) Register(c *gin.Context) {
  	var req registerRequest
  	if err := c.ShouldBindJSON(&req); err != nil {
      		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      		return
      	}
  	user, err := h.authSvc.Register(req.Email, req.Username, req.Password, req.FullName)
  	if err != nil {
      		if errors.Is(err, services.ErrUserAlreadyExists) {
            			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            			return
            		}
      		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      		return
      	}
  	c.JSON(http.StatusCreated, gin.H{"user": user.PublicView()})
  }

func (h *AuthHandler) Login(c *gin.Context) {
  	var req loginRequest
  	if err := c.ShouldBindJSON(&req); err != nil {
      		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      		return
      	}
  	user, tokens, err := h.authSvc.Login(req.Email, req.Password)
  	if err != nil {
      		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
      		return
      	}
  	c.JSON(http.StatusOK, gin.H{"user": user.PublicView(), "tokens": tokens})
  }
