package censys

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	idAPI     = "TEST_ID"
	secretAPI = "TEST_SECRET"
	stubsDir  = "stubs"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setUpTestServe() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(nil, idAPI, secretAPI)
	client.BaseURL = server.URL
}

func tearTestServer() {
	server.Close()
}

func getStubs(t *testing.T, stubName string) []byte {
	stubpath := fmt.Sprintf("%s/%s.json", stubsDir, stubName)
	content, err := ioutil.ReadFile(stubpath)
	if err != nil {
		t.Errorf("could not get stub %v", err)
	}
	return content
}

func testNewClient(t *testing.T) {
	cli := NewClient(nil, idAPI, secretAPI)
	assert.Equal(t, idAPI, cli.APIID)
	assert.Equal(t, secretAPI, cli.APISecret)
}
