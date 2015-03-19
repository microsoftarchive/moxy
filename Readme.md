## Moxy - A multi-host reverse proxy for golang

[![Build Status](https://travis-ci.org/wunderlist/moxy.svg?branch=master)](https://travis-ci.org/wunderlist/moxy)

The stdlib `ReverseProxy` from `net/http/httputil` has 2 issues that this project solves

1. There is single host proxy provided by `NewSingleHostReverseProxy`, but there is nothing to create a multi-host proxy
2. While it's possible to process a request before sending it to the proxy, the standard proxy doesn't provide a way to process the proxy response before flushing it out to the clients.

moxy aims to solve both of these problems by

1. Providing `moxy.NewReverseProxy`, that supports load-balancing across multiple hosts
2. Adding suppport for `moxy.FilterFunc`, a chain of functions that run over the proxy response & optionally make changes to it, before sending it back to the clients.



#### Example (using negroni)

```go
package main

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/wunderlist/moxy"
  "net/http"
)

func AddSecurityHeaders(request *http.Request, response *http.Response) {
  response.Header.Del("X-Powered-By")
  response.Header.Set("X-Super-Secure", "Yes!!")
}

func main() {

  hosts := []string{"dynamic.host.com"}
  filters := []moxy.FilterFunc{AddSecurityHeaders}
  proxy := moxy.NewReverseProxy(hosts, filters)

  router := mux.NewRouter()
  router.HandleFunc("/resource1", proxy.ServeHTTP)
  router.HandleFunc("/resource2", proxy.ServeHTTP)

  app := negroni.New()
  app.UseHandler(router)
  app.Run(":3009")
}

```