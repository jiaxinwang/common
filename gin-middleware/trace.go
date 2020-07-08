package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Trace ...
func Trace(c *gin.Context) {
	c.Set("requestID", uuid.NewV4().String())
	c.Next()
}
