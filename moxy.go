package moxy

import (
	"net/http"
)

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
