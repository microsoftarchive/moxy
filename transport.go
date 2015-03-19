package moxy

import (
	"net"
	"net/http"
	"time"
)

type Transport struct {
	tr *http.Transport
}

func NewClient() *http.Client {
	transport := NewTransport()
	return &http.Client{Transport: transport}
}

func NewTransport() *Transport {
	tr := &http.Transport{
		DisableKeepAlives:     true,
		MaxIdleConnsPerHost:   100000,
		DisableCompression:    true,
		ResponseHeaderTimeout: 30 * time.Second,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
	}
	return &Transport{tr: tr}
}

func (t *Transport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	response, err = t.tr.RoundTrip(request)
	return
}
