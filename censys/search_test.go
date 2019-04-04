package censys

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpv4Scan(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	expectedIP := "144.57.12.1"

	mux.HandleFunc(string(IPV4), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		var query searchQuery
		dec := json.NewDecoder(r.Body)
		dec.Decode(&query)
		ip := query.Query
		assert.NotEmpty(t, ip)
		assert.NotNil(t, net.ParseIP(expectedIP))
		w.Write(getStubs(t, "scan"))
	})
	search, err := client.Search(context.Background(), expectedIP, IPV4)

	ExMetadata := metadata{
		Count:       127530942,
		Query:       "*",
		BackendTime: 263,
		Page:        1,
		Pages:       1275310,
	}
	exResults := []results{
		results{
			IP: "173.205.31.126",
			Protocols: []string{
				"80/http",
				"443/https",
			},
		},
		results{
			IP: "213.149.206.213",
			Protocols: []string{
				"80/http",
			},
		},
		results{
			IP: "84.206.102.184",
			Protocols: []string{
				"80/http",
			},
		},
	}
	expectedSearch := &Search{
		Status:   "ok",
		Metadata: ExMetadata,
		Results:  exResults,
	}
	assert.Nil(t, err)
	assert.IsType(t, expectedSearch, search)
	assert.Equal(t, expectedSearch, search)

}
