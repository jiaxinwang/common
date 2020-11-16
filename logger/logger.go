package logger

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/sjson"
)

// Pretty ...
func Pretty(val interface{}, ignore ...string) string {
	if b, err := jsoniter.MarshalIndent(val, "", " "); err == nil {
		ret := string(b)
		for _, v := range ignore {
			ret, _ = sjson.Delete(ret, v)
		}
		return ret
	} else {
		logrus.WithError(err).Error()
		return ""
	}
}
