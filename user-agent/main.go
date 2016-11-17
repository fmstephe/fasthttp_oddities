package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/adjust/backend-vendor/github.com/valyala/fasthttp"
)

func main() {
	server := httptest.NewServer(&handler{})
	URL := server.URL
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URL)
	req.Header.Add("User-Agent", "Test-Agent")
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", err)
	} else {
		fmt.Fprintf(os.Stdout, "No-Error: %s", string(resp.Body()))
	}
}

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Your User-Agent: ", r.Header.Get("User-Agent"))
}
