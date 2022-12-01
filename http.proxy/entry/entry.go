package entry

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"http.proxy/cache"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Initialize() *gin.Engine {

	cache.GetInstance().Init()

	r := gin.Default()

	r.Any("/*proxyPath", ProxyHandler)

	return r
}

func ProxyHandler(c *gin.Context) {

	proxymap := cache.GetInstance().GetCache()
	var target string
	path := c.Param("proxyPath")

	header := c.Request.Header["X-Lyrid-Xfh"] // only check things coming from our proxies
	if len(header) == 0 {
		target = proxymap[c.Request.Host]
	} else {
		target = proxymap[header[0]]
	}

	if target == "" {
		c.String(http.StatusNotFound, "proxy not found: "+c.Request.Host)

		for s, i := range c.Request.Header {
			for _, s2 := range i {
				fmt.Println("header: ", s, " ", s2)
			}
		}
	} else {
		remote, err := url.Parse(target)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		//Define the director func
		//This is a good place to log, for example
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = remote.Path + path
			req.URL.RawQuery = c.Request.URL.RawQuery
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
