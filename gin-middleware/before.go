package middleware

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// BeforeQueryAction ...
type BeforeQueryAction struct {
	Param  string
	Model  interface{}
	Result string
}

// BeforeAction ...
type BeforeAction struct {
	Table      string
	Values     url.Values
	TargetName string
}

// BeforeQuery ...
func BeforeQuery(c *gin.Context) {
	// from (table,conditions)
	// transfer
	// to (param id)

	// payload
	if v, ok := c.Get("before_action"); ok {
		payload := v.(BeforeAction)
		logrus.Print(payload)
	}

	c.Next()
}
