package lazy

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"

	"github.com/tidwall/sjson"
)

// Handle executes actions and returns response
func Handle(c *gin.Context) (data []map[string]interface{}, err error) {
	var config *Configuration
	if v, ok := c.Get("lazy-configuration"); ok {
		config = v.(*Configuration)
	} else {
		return nil, errors.New("can't find lazy-configuration")
	}

	set := foreignOfModel((*config).Model)

	if config.Before != nil {
		_, _, errBefore := config.Before.Action(c, config.DB, *config, nil)
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

	eq, gt, lt, gte, lte := LazyURLValues(config.Model, merged)

	sel := sq.Select(config.Columms).From(config.Table).Limit(param.Size).Offset(param.Size * param.Page)
	sel = SelectBuilder(sel, eq, gt, lt, gte, lte)
	data, err = ExecSelect(config.DB, sel)
	if err != nil {
		return
	}

	for _, v := range data {
		if err := MapStruct(v, config.Model); err != nil {
			return nil, err
		}
		tmp := clone(config.Model)

		for _, v := range set {
			value := valueOfTag(tmp, v[ForeignOfModelID])
			eq := map[string][]interface{}{v[ForeignOfModelForeignID]: []interface{}{value}}
			data, err := SelectEq(config.DB, v[ForeignOfModelForeignTable], "*", eq)
			if err != nil {
				return nil, err
			}
			if len(data) == 1 {
				jbyte, _ := json.Marshal(tmp)
				assemble, _ := sjson.Set(string(jbyte), v[ForeignOfModelName], data[0])
				json.Unmarshal([]byte(assemble), tmp)
			}
		}

		// TODO: batch

		config.Results = append(config.Results, tmp)
	}

	count := int64(len(data))

	if config.NeedCount {
		sel := sq.Select(`count(1) as c`).From(config.Table).Limit(param.Size).Offset(param.Size * param.Page)
		sel = SelectBuilder(sel, eq, gt, lt, gte, lte)
		data, err = ExecSelect(config.DB, sel)
		if err != nil {
			return
		}
		if len(data) == 1 {
			iter, _ := data[0][`c`]
			count, err = strconv.ParseInt(fmt.Sprintf("%v", iter), 10, 64)
			if err != nil {
				return
			}
		}
	}
	c.Set("lazy-count", count)
	c.Set("lazy-data", config.Results)
	c.Set("lazy-results", map[string]interface{}{"count": count, "items": config.Results})
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
