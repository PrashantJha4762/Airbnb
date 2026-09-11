package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ReverseProxy(targetUrl, PathPrefix string) http.HandlerFunc {
	target,err:=url.Parse(targetUrl)
	if err != nil {
		fmt.Println("Some error occured",err)
		return nil
	}
	proxy:=httputil.NewSingleHostReverseProxy(target)

	originalDirector:=proxy.Director

	proxy.Director=func(r *http.Request){
		originalDirector(r)
		r.URL.Path=strings.TrimPrefix(r.URL.Path,PathPrefix)
		r.Host=target.Host
		if userId,ok:=r.Context().Value("userId").(string);ok{
			r.Header.Set("X-User-Id",userId)
		}
	}
	return proxy.ServeHTTP
}