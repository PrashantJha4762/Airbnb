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

// ReverseProxy creates an HTTP handler that forwards incoming requests to a
// different backend service. It is intended for use by a gateway or BFF: the
// caller registers the returned handler for a route, while this function
// transparently relays matching requests to targetUrl and returns the backend's
// response to the original client.
//
// Implementation details:
//   - targetUrl is parsed into a url.URL. This provides the destination scheme,
//     host, and optional base path required by Go's reverse-proxy implementation.
//     If it cannot be parsed, the function logs the error and returns nil, so the
//     caller should only register the handler after supplying a valid URL.
//   - httputil.NewSingleHostReverseProxy(target) creates the standard library
//     reverse proxy. It supplies the serving logic that sends the request to the
//     target service, streams its response, and handles normal proxy behavior.
//   - The proxy's original Director is preserved and called first. The Director
//     is responsible for rewriting an inbound request into an outbound request;
//     retaining the original one preserves the URL and query handling supplied
//     by NewSingleHostReverseProxy.
//   - After the default rewrite, PathPrefix is removed from r.URL.Path. For
//     example, a gateway request to /users/profile with PathPrefix /users is
//     forwarded to the target as /profile. This lets the public gateway route
//     differ from the route expected by the downstream service.
//   - r.Host is set to target.Host so the backend receives its own host header,
//     rather than the host used by the client to reach the gateway.
//   - If authentication middleware previously placed a string userId in the
//     request context, it is copied to the X-User-Id header. The downstream
//     service can then identify the authenticated user without needing to repeat
//     the gateway's authentication step.
//
// The returned proxy.ServeHTTP method satisfies http.HandlerFunc and can be
// passed directly to a router or net/http ServeMux.
