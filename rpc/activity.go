package rpc

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"net/http"
)

const (
	defaultActivityServerScheme = "http"
	defaultActivityServerAddr   = "47.244.34.27:7000" // 默认内网转发,本地回环
)

var activityServerHost = []string{defaultActivityServerAddr}
var transport *http.Transport

func init() {
	transport = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
		MaxIdleConns:        defaultMaxIdleConns,
		DialContext: (&net.Dialer{
			Timeout:   defaultDialTimeout,
			KeepAlive: defaultKeepAliveTimeout,
		}).DialContext,
		IdleConnTimeout:       defaultIdleConnTimeout,
		ExpectContinueTimeout: defaultExpectContinueTimeout,
		DisableKeepAlives:     false,
	}
}

// 转发
func ReverseProxy(c *gin.Context) {
	// step 1: resolve proxy address, change scheme and host in requets
	req := c.Request
	// step 1.1 trace

	req.URL.Scheme = defaultActivityServerScheme
	req.URL.Host = getActivityHost()
	req.Host = getActivityHost()

	// step 2: use http.Transport to do request to real server.
	resp, err := transport.RoundTrip(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error")
		return
	}
	// step 3: return real server response to upstream.
	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}

	defer resp.Body.Close()
	if _, err := bufio.NewReader(resp.Body).WriteTo(c.Writer); err != nil {
	}
}
func getActivityHost() string {
	l := len(activityServerHost)
	r := rand.Intn(l) // 随机一个
	return activityServerHost[r]
}
