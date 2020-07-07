package middleware

import (
	"github.com/gin-gonic/gin"
)

// LazyResponse ...
func LazyResponse(c *gin.Context) {
	defer func() {
		code := int(0)
		data, _ := c.Get("data")
		if codeInter, ok := c.Get("code"); ok {
			code = codeInter.(int)
		}
		if clientErr, clientErrExist := c.Get("client_err"); clientErrExist {
			c.JSON(200, map[string]interface{}{
				"data":       nil,
				"error_no":   code,
				"error_msg":  clientErr.(error).Error(),
				"request_id": c.MustGet("requestID"),
			})
			return
		}

		if serverErr, serverErrExist := c.Get("server_err"); serverErrExist {
			c.JSON(200, map[string]interface{}{
				"data":       nil,
				"error_no":   code,
				"error_msg":  serverErr.(error).Error(),
				"request_id": c.MustGet("requestID"),
			})
			return
		}

		c.JSON(200, map[string]interface{}{
			"data":       data,
			"error_no":   0,
			"error_msg":  ``,
			"request_id": c.MustGet("requestID"),
		})

	}()
	c.Next()
}
