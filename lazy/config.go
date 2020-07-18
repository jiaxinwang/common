package lazy

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Configuration configs lazy values and actions
type Configuration struct {
	DB        *gorm.DB
	Table     string
	Columms   string
	Model     interface{}
	Results   []interface{}
	Before    *ActionConfiguration
	After     *ActionConfiguration
	NeedCount bool
}

// ActionConfiguration configs action, before-action, after-action values and actions
type ActionConfiguration struct {
	Table     string
	Model     interface{}
	ResultMap map[string]string
	Action    func(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error)
	// Target    []interface{}
	// Columms   string
}
