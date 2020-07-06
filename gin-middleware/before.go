package middleware

import (
	"github.com/gin-gonic/gin"
)

// BeforeQueryAction ...
type BeforeQueryAction struct {
	Param  string
	Model  interface{}
	Result string
}

// BeforeQuery ...
func BeforeQuery(c *gin.Context) {
	// inter, ok := c.Get("lazy_before_query")
	// if !ok {
	// 	return
	// }
	// logrus.Print(action)
}
