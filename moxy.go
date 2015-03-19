package moxy

import (
	"net/http"
)

// NewReverseProxy returns a new ReverseProxy that load-balances the proxy requests between multiple hosts
// It also allows to define a chain of filter functions to process the outgoing response(s)
func NewReverseProxy(hosts []string, filters []FilterFunc) *ReverseProxy {
	director := func(request *http.Request) {
		host, _ := pick(hosts)
		request.URL.Scheme = "http"
		request.URL.Host = host
	}
	return &ReverseProxy{
		Transport: NewTransport(),
		Director:  director,
		Filters:   filters,
	}
}
