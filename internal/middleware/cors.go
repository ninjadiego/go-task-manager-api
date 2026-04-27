																																																				package middleware

import (
  	"net/http"

  	"github.com/gin-gonic/gin"
  )

func CORS(allowedOrigins []string) gin.HandlerFunc {
  	originSet := make(map[string]struct{}, len(allowedOrigins))
  	for _, o := range allowedOrigins {
      		originSet[o] = struct{}{}
      	}
  	return func(c *gin.Context) {
      		origin := c.GetHeader("Origin")
      		if _, ok := originSet[origin]; ok || len(allowedOrigins) == 0 {
            			if origin != "" {
                    				c.Header("Access-Control-Allow-Origin", origin)
                    			} else {
                    				c.Header("Access-Control-Allow-Origin", "*")
                    			}
            		}
      		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
      		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
      		c.Header("Access-Control-Allow-Credentials", "true")
      		if c.Request.Method == http.MethodOptions {
            			c.AbortWithStatus(http.StatusNoContent)
            			return
            		}
      		c.Next()
      	}
  }
