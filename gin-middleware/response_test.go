package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	logrus.SetReportCaller(true)
}

func route() *gin.Engine {
	r := gin.Default()
	r.Use(LazyResponse)
	r.GET("/client_err", func(c *gin.Context) {
		err := errors.New("a client error")
		c.Set("client_err", err)
		c.Set("requestID", "xxx-xxx")
		return
	})
	r.GET("/server_err", func(c *gin.Context) {
		err := errors.New("a server error")
		c.Set("server_err", err)
		c.Set("requestID", "xxx-xxx")
		return
	})
	r.GET("/int", func(c *gin.Context) {
		c.Set("data", int(10086))
		c.Set("requestID", "xxx-xxx")
		return
	})
	r.GET("/struct", func(c *gin.Context) {
		c.Set(
			"data",
			struct {
				Name string
			}{Name: "monkey"},
		)
		c.Set("requestID", "xxx-xxx")
		return
	})

	return r
}

func TestLazyResponse(t *testing.T) {
	type Response struct {
		Data      interface{} `json:"data"`
		ErrorMsg  string      `json:"error_msg"`
		ErrorNo   int         `json:"error_no"`
		RequestID string      `json:"request_id"`
	}

	router := route()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/client_err", nil)
	router.ServeHTTP(w, req)
	want := Response{nil, `a client error`, int(0), `xxx-xxx`}
	ret := Response{}
	err := json.Unmarshal(w.Body.Bytes(), &ret)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
	assert.Equal(t, want, ret)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/server_err", nil)
	router.ServeHTTP(w, req)
	want = Response{nil, `a server error`, int(0), `xxx-xxx`}
	ret = Response{}
	err = json.Unmarshal(w.Body.Bytes(), &ret)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, err)
	assert.Equal(t, want, ret)

	// w = httptest.NewRecorder()
	// req, _ = http.NewRequest("GET", "/int", nil)
	// router.ServeHTTP(w, req)
	// want = Response{int(10086), ``, int(0), `xxx-xxx`}
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
