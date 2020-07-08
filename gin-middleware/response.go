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
		if _, ok := ret["data"]; !ok {
			ret["data"] = nil
		}
		if _, ok := ret["error_no"]; !ok {
			if _, ok := ret["error_msg"]; ok {
				ret["error_no"] = 400
			} else {
				ret["error_no"] = 0
			}
		}
		if _, ok := ret["error_msg"]; !ok {
			ret["error_msg"] = ``
		}

		ret["request_id"] = c.MustGet("requestID")
		c.JSON(200, ret)
	}()
	c.Next()
}
