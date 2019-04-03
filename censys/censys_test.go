package censys

import (
	"net/http"
	"net/http/httptest"
)

const (
	idAPI     = "d8b4925a-2887-41d6-bbf5-6a00ea7dc529"
	secretAPI = "q8qbfwgaEIJ3BYbLnVfHHzFIWvzd6Prb"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setUpTestServe() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient{
		idAPI, secretAPI,
	}
	client.BaseURL = server.URL
}

func tearTestServer() {
	server.Close()
}
