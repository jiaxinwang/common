package lazy

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Configuration configs lazy values and actions
type Configuration struct {
	DB              *gorm.DB
	BeforeTables    string
	BeforeColumm    string
	BeforeStruct    interface{}
	BeforeResultMap map[string]string
	BeforeAction    func(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error)
	AfterTables     string
	AfterStruct     interface{}
	AfterResultMap  map[string]string
	AfterAction     func(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error)
	Table           string
	Columm          string
	Struct          interface{}
	Targets         []interface{}
}
