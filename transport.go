package moxy

import (
	"net"
	"net/http"
	"time"
)

// Transport wraps http.Transport to be able to instrument client calls
// This is required for keeping metrics on outgoing http requests made by the reverse-proxy
type Transport struct {
	tr *http.Transport
}

// NewClient creates a http.Client for making outgoing requests
// It uses an instance of our custom Transport
func NewClient() *http.Client {
	transport := NewTransport()
	return &http.Client{Transport: transport}
}

// NewTransport creates an instance of Transport with some opinionated defaults
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

// RoundTrip implements the RoundTripper interface.
// This is required to make the transport work
func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	return t.tr.RoundTrip(request)
}
