package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func route() *gin.Engine {
	r := gin.Default()
	r.Use(Trace)
	r.Use(LazyResponse)
	r.GET("/client_err", func(c *gin.Context) {
		c.Set(
			"ret",
			map[string]interface{}{
				"data":      nil,
				"error_no":  0,
				"error_msg": "a client error",
			},
		)
		return
	})
	r.GET("/server_err", func(c *gin.Context) {
		c.Set(
			"ret",
			map[string]interface{}{
				"data":      nil,
				"error_no":  0,
				"error_msg": "a server error",
			},
		)
		return
	})
	r.GET("/int", func(c *gin.Context) {
		c.Set(
			"ret",
			map[string]interface{}{
				"data": int(10086),
				// "error_no":  0,
				// "error_msg": "",
			},
		)
		return
	})
	r.GET("/struct", func(c *gin.Context) {
		c.Set(
			"ret",
			map[string]interface{}{
				"data": struct {
					Name string
				}{Name: "monkey"},
				// "error_no":  0,
				// "error_msg": "",
			},
		)
		return
	})

	return r
}

func TestLazyResponse(t *testing.T) {
	type Response struct {
		Data     interface{} `json:"data"`
		ErrorMsg string      `json:"error_msg"`
		ErrorNo  int         `json:"error_no"`
		// RequestID string      `json:"request_id"`
	}

	router := beforeRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/client_err", nil)
	router.ServeHTTP(w, req)
	want := Response{nil, `a client error`, int(0)}
	ret := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &ret)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
	assert.Equal(t, want, ret)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/server_err", nil)
	router.ServeHTTP(w, req)
	want = Response{nil, `a server error`, int(0)}
	ret = Response{}
	err = json.Unmarshal(w.Body.Bytes(), &ret)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
	assert.Equal(t, want, ret)

	// w = httptest.NewRecorder()
	// req, _ = http.NewRequest("GET", "/int", nil)
	// router.ServeHTTP(w, req)
	// want = Response{int(10086), ``, int(0)}
	// ret = Response{}
	// err = json.Unmarshal(w.Body.Bytes(), &ret)
	// assert.Equal(t, 200, w.Code)
	// assert.NoError(t, err)
	// assert.Equal(t, want, ret)

	// w = httptest.NewRecorder()
	// req, _ = http.NewRequest("GET", "/struct", nil)
	// router.ServeHTTP(w, req)
	// want = Response{
	// 	struct {
	// 		Name string
	// 	}{Name: "monkey"},
	// 	``,
	// 	int(0),
	// 	`xxx-xxx`}
	// ret = Response{}
	// err = json.Unmarshal(w.Body.Bytes(), &ret)
	// assert.Equal(t, 200, w.Code)
	// assert.NoError(t, err)
	// assert.Equal(t, want, ret)

}
