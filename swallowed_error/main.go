package main

import (
	"fmt"
	"io/ioutil"
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
	req.Header.SetMethod("POST")
	req.SetBodyString("p=q")
	req.Header.Add("Content-Length", "One million dollars!")
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
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		fmt.Fprint(w, "Body of your request: ", string(body))
	}
}
