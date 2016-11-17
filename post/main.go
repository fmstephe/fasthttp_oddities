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
	req.Header.String() // Request has now cached that it is a GET request
	req.Header.SetMethod("POST")
	req.SetBodyString("p=q")
	fmt.Fprintf(os.Stdout, "IsGet() =  %t\n", req.Header.IsGet())
	fmt.Fprintf(os.Stdout, "Method() = %s\n\n", req.Header.Method())
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, "Request Body: ", string(body))
}
