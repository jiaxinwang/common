package lazy

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DefaultAfterAction ...
func DefaultAfterAction(c *gin.Context, gormDB *gorm.DB, config Configuration, payload interface{}) (result interface{}, reduce map[string][]string, err error) {
	// afterConfig := config.After
	// last, exits := c.Get("lazy-data")
	// if !exits {
	// 	return nil, nil, errors.New("last data not exist")
	// }
	// eq := make(map[string][]interface{})

	// sel := sq.Select(afterConfig.Columms).From(afterConfig.Table)
	// sel = SelectBuilder(sel, eq, nil, nil, nil, nil)
	// data, err = ExecSelect(config.DB, sel)

	return
}
