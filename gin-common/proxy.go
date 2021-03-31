package gincommon

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// WildProxy ...
func WildProxy(c *gin.Context) {
	// proxyPath := strings.TrimPrefix(c.Param("proxyPath"), `/`)
	// remote, err := url.Parse(proxyPath)
	remote, err := url.Parse("https://httpbin.org/anything/1")
	if err != nil {
		panic(err)
	}
	logrus.Print("112233 ", remote.String())
	proxy := httputil.NewSingleHostReverseProxy(remote)
	//ReverseProxy
	// proxy := &httputil.ReverseProxy{}
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path

		logrus.Print(req.URL.Scheme)
		logrus.Print(req.URL.Host)
		logrus.Print(req.URL.Path)
		// req.URL.Path = `/anything/1`
		req.Header = c.Request.Header
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
