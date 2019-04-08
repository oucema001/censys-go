package censys

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetView(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	expextedQuery := "google.com"
	mux.HandleFunc(string(IPV4VIEW), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		paths := strings.Split(r.URL.Path, "/")
		path := paths[len(paths)-1:]
		pa := path[0]
		assert.Equal(t, expextedQuery, pa)

		w.Write(getStubs(t, "view"))
	})
	view, err := client.GetView(context.Background(), IPV4VIEW, "google.com")
	sigAlg := SignatureAlgorithm{
		Oid:  "Test",
		Name: "Test",
	}
	signature := Signature{
		SelfSigned:         false,
		Valid:              false,
		Value:              "Test",
		SignatureAlgorithm: sigAlg,
	}
	issuer := Issuer{
		CommonName: []string{"Test"},
	}
	expectedRes := &Parsed{
		FingerprintSha1: "Test",
		SubjectDn:       "Test",
		Issuer:          issuer,
	}

	assert.Nil(t, err)
	assert.IsType(t, expectedRes, view)
	assert.EqualValues(t, expectedRes, view)
}
