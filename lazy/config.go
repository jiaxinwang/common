package lazy

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Configuration configs lazy values and actions
type Configuration struct {
	DB     *gorm.DB
	Table  string
	Columm string
	Model  interface{}
	Target []interface{}
	Before *ActionConfiguration
	After  *ActionConfiguration

	// BeforeTables    string
	// BeforeColumm    string
	// BeforeStruct    interface{}
	// BeforeResultMap map[string]string
	// BeforeAction    func(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error)
	// AfterTables     string
	// AfterStruct     interface{}
	// AfterResultMap  map[string]string
	// AfterAction     func(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error)
}

// ActionConfiguration configs action, before-action, after-action values and actions
type ActionConfiguration struct {
	Table     string
	Columms   string
	Model     interface{}
	ResultMap map[string]string
	Action    func(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error)
	Target    []interface{}
}
