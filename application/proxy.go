package application

import (
	"github.com/projectxpolaris/webcontainer/config"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var apiProxyHandler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	rs1 := config.Instance.APIProxy.Proxy
	targetUrl, _ := url.Parse(rs1)
	realPath := request.URL.Path
	if config.Instance.APIProxy.Rewrite {
		realPath = strings.TrimPrefix(request.URL.Path, config.Instance.APIProxy.Prefix)
	}
	request.URL.Path = realPath
	request.Host = targetUrl.Host
	request.URL.Scheme = targetUrl.Scheme
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(writer, request)
}
