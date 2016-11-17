package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/adjust/backend-vendor/github.com/valyala/fasthttp"
)

func main() {
	server := httptest.NewServer(&handler{})
	URL := server.URL
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URL)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{
		ReadTimeout: time.Millisecond,
	}
	if err := client.Do(req, resp); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", err)
	} else {
		fmt.Fprintf(os.Stdout, "No-Error: %s", string(resp.Body()))
	}
}

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Millisecond)
	fmt.Fprint(w, "hello fasthttp client")
}
