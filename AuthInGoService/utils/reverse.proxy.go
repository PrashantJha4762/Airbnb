package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ReverseProxy(targetURL, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetURL)
	if err != nil || target.Scheme == "" || target.Host == "" {
		return func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "proxy target is unavailable", http.StatusBadGateway)
		}
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix)
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}
		originalDirector(r)
		r.Host = target.Host
		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-User-Id", userId)
		}
	}
	return proxy.ServeHTTP
}

// What is a reverse proxy?
// A reverse proxy sits between the client and another server. The client sends
// a request to this application, and this function sends that request to the
// correct backend service. When the backend sends a response, the proxy gives
// that response back to the client.
//
// Example:
// A client calls:      GET /users/profile
// This application forwards the request to the user service, then returns the
// user service's response to the client.
//
// How this function works:
// 1. targetUrl is the address of the backend service, for example
//    "http://localhost:8081". url.Parse changes that text into a URL value Go
//    can use. If the address is invalid, an error is printed and nil is returned.
//
// 2. NewSingleHostReverseProxy creates the proxy. Go's built-in proxy handles
//    sending the request to the backend and returning the backend response.
//
// 3. originalDirector keeps Go's default request setup. We call it first so
//    the request is correctly changed to use the target backend URL.
//
// 4. PathPrefix is removed from the path before forwarding the request. For
//    example, when PathPrefix is "/users", a request for "/users/profile" is
//    sent to the backend as "/profile". This is useful when the gateway route
//    has a prefix that the backend service does not need.
//
// 5. The Host header is changed to the backend's host. This tells the backend
//    that the request is meant for it, instead of the gateway.
//
// 6. If earlier authentication code saved a user ID in the request context,
//    this function adds it as the X-User-Id header. The backend can read this
//    header to know which logged-in user made the request.
//
// Finally, proxy.ServeHTTP is returned as an http.HandlerFunc, so it can be
// registered in a Go router to handle requests for a particular route.
