package lazy

import (
	"reflect"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jiaxinwang/common"
)

// Handle ...
func Handle(c *gin.Context) (data []map[string]interface{}, err error) {
	var config *Configuration
	if v, ok := c.Get("lazy-configuration"); ok {
		config = v.(*Configuration)
		_, _, errBefore := config.BeforeAction(c, config.DB, *config, nil)
		if errBefore != nil {
			return nil, errBefore
		}
	}

	param := struct {
		Offset uint64 `form:"offset" binding:"gte=0"`
		Page   uint64 `form:"page" binding:"gte=0"`
		Size   uint64 `form:"size" binding:"gte=1,lte=1000"`
	}{
		Size: 1000,
	}

	if err = c.ShouldBindQuery(&param); err != nil {
		return nil, err
	}

	var merged map[string][]string
	additional, ok := c.Get("_additional_values")
	if ok {
		merged = mergeValues(c.Request.URL.Query(), additional.(map[string][]string))
	}

	eq, gt, lt, gte, lte := LazyURLValues(config.Struct, merged)

	sel := sq.Select(config.Columm).From(config.Table).Limit(param.Size).Offset(param.Size * param.Page)
	sel = common.SelectBuilder(sel, eq, gt, lt, gte, lte)
	data, err = Query(config.DB, sel)

	for _, v := range data {
		MapStruct(v, config.Struct)
		tmp := clone(config.Struct)
		config.Targets = append(config.Targets, tmp)
	}
	return
}

func clone(inter interface{}) interface{} {
	newInter := reflect.New(reflect.TypeOf(inter).Elem())

	val := reflect.ValueOf(inter).Elem()
	taggetVal := newInter.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := taggetVal.Field(i)
		field.Set(val.Field(i))
	}
	return newInter.Interface()
}
