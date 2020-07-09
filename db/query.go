package db

import (
	"reflect"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
)

// Query ...
func Query(db *gorm.DB, active sq.SelectBuilder) (ret []map[string]interface{}, err error) {
	ret = make([]map[string]interface{}, 0)
	sql, args, err := active.ToSql()
	if err != nil {
		return ret, err
	}

	rows, sqlErr := db.Raw(sql, args...).Rows()

	defer rows.Close()
	if sqlErr != nil {
		return ret, sqlErr
	}

	columns, err := rows.Columns()
	if err != nil {
		return ret, err
	}
	length := len(columns)
	for rows.Next() {
		current := makeResultReceiver(length)
		if err := rows.Scan(current...); err != nil {
			return ret, err
		}
		value := make(map[string]interface{})
		for i := 0; i < length; i++ {
			k := columns[i]
			val := *(current[i]).(*interface{})
			if val == nil {
				value[k] = nil
				continue
			}
			vType := reflect.TypeOf(val)
			switch vType.String() {
			case "uint8":
				value[k] = val.(int8)
			case "uint16":
				value[k] = val.(int16)
			case "uint32":
				value[k] = val.(int32)
			case "uint64":
				value[k] = val.(int64)
			case "int8":
				value[k] = val.(int8)
			case "int16":
				value[k] = val.(int16)
			case "int32":
				value[k] = val.(int32)
			case "int64":
				value[k] = val.(int64)
			case "bool":
				value[k] = val.(bool)
			case "string":
				value[k] = val.(string)
			case "time.Time":
				value[k] = val.(time.Time)
			case "[]uint8":
				value[k] = string(val.([]uint8))
			default:
				// logrus.Warnf("unsupport data type '%s' now\n", vType)
			}
		}
		ret = append(ret, value)
	}

	return ret, nil

}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}
