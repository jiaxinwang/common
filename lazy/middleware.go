package lazy

import (
	"github.com/gin-gonic/gin"
)

// Middleware run the query
func Middleware(c *gin.Context) {
	defer func() {
		if _, err := Handle(c); err != nil {
			c.Set("error_msg", err.Error())
			return
		}
		if data, exist := c.Get("lazy-results"); exist {
			c.Set("ret", map[string]interface{}{"data": data})
		}
		return
	}()
	c.Next()
}
