package middleware

import (
	"github.com/gin-gonic/gin"
)

// LazyResponse ...
func LazyResponse(c *gin.Context) {
	defer func() {
		ret := make(map[string]interface{})
		if v, ok := c.Get("ret"); ok {
			ret = v.(map[string]interface{})
		}
		ret["request_id"] = c.MustGet("requestID")
		c.JSON(200, ret)
	}()
	c.Next()
}
