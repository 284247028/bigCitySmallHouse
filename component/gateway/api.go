package main

import (
	"bigCitySmallHouse/component/base/base_action"
	"bigCitySmallHouse/component/etcd"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func apiProxy(ctx *gin.Context) {
	serviceName := ctx.Param("serviceName")
	proxyPath := ctx.Param("proxyPath")
	host, err := etcd.GetEtcd().GetServiceHost(serviceName)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	uri := "http://" + host
	tUrl, err := url.Parse(uri)
	if err != nil {
		base_action.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(tUrl)
	reverseProxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = tUrl.Host
		req.URL.Scheme = tUrl.Scheme
		req.URL.Host = tUrl.Host
		req.URL.Path = proxyPath
	}
	reverseProxy.ServeHTTP(ctx.Writer, ctx.Request)

}
