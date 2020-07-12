package lazy

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gm "github.com/jiaxinwang/common/gin-middleware"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func router() *gin.Engine {
	r := gin.Default()
	r.Use(gm.Trace)
	r.Use(gm.LazyResponse)

	return r
}

func TestActionHandle(t *testing.T) {
	r := router()
	r.GET("/dogs", func(c *gin.Context) {
		config := Configuration{
			DB:        gormDB,
			Table:     "dogs",
			Columms:   "*",
			Model:     &Dog{},
			Results:   []interface{}{},
			NeedCount: true,
		}
		c.Set("lazy-configuration", &config)
		if _, err := Handle(c); err != nil {
			c.Set("error_msg", err.Error())
			return
		}
		if v, exist := c.Get("lazy-results"); exist {
			c.Set("ret", map[string]interface{}{"data": v})
		}
		return
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dogs", nil)
	q := req.URL.Query()
	q.Add("id", `1`)
	q.Add("id", `2`)
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)
	response := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
}

func TestActionHandleMiddleware(t *testing.T) {
	r := router()
	r.Use(Middleware)
	r.GET("/dogs", func(c *gin.Context) {
		config := Configuration{
			DB:        gormDB,
			Table:     "dogs",
			Columms:   "*",
			Model:     &Dog{},
			Results:   []interface{}{},
			NeedCount: true,
		}
		c.Set("lazy-configuration", &config)
		return
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dogs", nil)
	q := req.URL.Query()
	q.Add("id", `1`)
	q.Add("id", `2`)
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)
	response := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
	logrus.Print(response)
}

func TestBeforeActionHandle(t *testing.T) {
	r := router()

	r.GET("/dogs", func(c *gin.Context) {
		var ret []interface{}
		config := Configuration{
			DB: gormDB,
			Before: &ActionConfiguration{
				Table:     "profiles",
				Columms:   "dog_id",
				Model:     &Profile{},
				ResultMap: map[string]string{"dog_id": "id"},
				Action:    DefaultBeforeAction,
			},
			Table:   "dogs",
			Columms: "*",
			Model:   &Dog{},
			Results: ret,
		}
		c.Set("lazy-configuration", &config)
		if _, err := Handle(c); err != nil {
			c.Set("error_msg", err.Error())
			return
		}
		c.Set("ret", map[string]interface{}{"data": config.Results})
		return
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dogs", nil)
	q := req.URL.Query()
	q.Add("before_dog_id", `1`)
	q.Add("before_dog_id", `2`)
	req.URL.RawQuery = q.Encode()

	r.ServeHTTP(w, req)
	ret := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &ret)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
	logrus.Print(ret)
}

// func TestAfterActionHandle(t *testing.T) {
// 	r := router()

// 	r.GET("/dogs", func(c *gin.Context) {
// 		var ret []interface{}
// 		config := Configuration{
// 			DB: gormDB,
// 			After: &ActionConfiguration{
// 				Table:     "profiles",
// 				Columms:   "dog_id",
// 				Model:     &Profile{},
// 				ResultMap: map[string]string{"dog_id": "id"},
// 				Action:    DefaultBeforeAction,
// 			},
// 			Table:   "dogs",
// 			Columms: "*",
// 			Model:   &Dog{},
// 			Results: ret,
// 		}
// 		c.Set("lazy-configuration", &config)
// 		if _, err := Handle(c); err != nil {
// 			c.Set("error_msg", err.Error())
// 			return
// 		}
// 		c.Set("ret", map[string]interface{}{"data": config.Results})
// 		return
// 	})

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/dogs", nil)
// 	q := req.URL.Query()
// 	q.Add("before_dog_id", `1`)
// 	q.Add("before_dog_id", `2`)
// 	req.URL.RawQuery = q.Encode()

// 	r.ServeHTTP(w, req)
// 	ret := Response{}
// 	err := json.Unmarshal(w.Body.Bytes(), &ret)
// 	assert.Equal(t, 200, w.Code)
// 	assert.NoError(t, err)
// 	logrus.Print(ret)
// }
