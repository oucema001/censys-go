package censys

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpv4Scan(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	expectedIP := &SearchQuery{
		Query:   "144.57.12.1",
		Page:    1,
		Fields:  []string{},
		Flatten: true,
	}

	mux.HandleFunc(string(IPV4SEARCH), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		var query SearchQuery
		dec := json.NewDecoder(r.Body)
		dec.Decode(&query)
		ip := query.Query
		assert.NotEmpty(t, ip)
		assert.NotNil(t, net.ParseIP(expectedIP.Query))
		w.Write(getStubs(t, "searchipv4"))
	})
	search, err := client.Search(context.Background(), expectedIP, IPV4SEARCH)

	ExMetadata := metadata{
		Count:       127530942,
		Query:       "*",
		BackendTime: 263,
		Page:        1,
		Pages:       1275310,
	}
	exResults := []results{
		results{
			Country:           "United States",
			RegisteredCountry: "United States",
			Longitude:         -119.7143,
			City:              "Boardman",
			IP:                "173.205.31.126",
			Protocols: []string{
				"80/http",
				"443/https",
			},
			RegisteredCountyCode: "US",
			CountryCode:          "US",
			Province:             "Oregon",
			Continent:            "North America",
			PostalCode:           "97818",
			TimeZone:             "America/Los_Angeles",
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

func TestWebSiteScan(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	expectedIP := &SearchQuery{
		Query:   "144.57.12.1",
		Page:    1,
		Fields:  []string{},
		Flatten: true,
	}

	mux.HandleFunc(string(IPV4SEARCH), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		var query SearchQuery
		dec := json.NewDecoder(r.Body)
		dec.Decode(&query)
		ip := query.Query
		assert.NotEmpty(t, ip)
		assert.NotNil(t, net.ParseIP(expectedIP.Query))
		w.Write(getStubs(t, "searchwebsite"))
	})
	search, err := client.Search(context.Background(), expectedIP, IPV4SEARCH)

	ExMetadata := metadata{
		Count:       198062,
		Query:       "www.google.com",
		BackendTime: 172,
		Page:        1,
		Pages:       1981,
	}
	exResults := []results{
		results{
			Domain:    "google.com",
			AlexaRank: 1,
		},
		results{
			Domain:    "twitter.com",
			AlexaRank: 11,
		},
		results{
			Domain:    "google.com.hk",
			AlexaRank: 37,
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

func TestCertificateScan(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	expectedCertQuery := &SearchQuery{
		Query:   "www.google.com",
		Page:    1,
		Fields:  []string{},
		Flatten: true,
	}

	mux.HandleFunc(string(IPV4SEARCH), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		var query SearchQuery
		dec := json.NewDecoder(r.Body)
		dec.Decode(&query)
		ip := query.Query
		assert.NotEmpty(t, ip)
		_, err := url.Parse(expectedCertQuery.Query)
		assert.Nil(t, err)
		w.Write(getStubs(t, "searchcertificates"))
	})
	search, err := client.Search(context.Background(), expectedCertQuery, IPV4SEARCH)

	exMetadata := metadata{
		Count:       2849,
		Query:       "www.google.com",
		BackendTime: 792,
		Page:        1,
		Pages:       29,
	}
	exResultsCert := []results{
		results{
			FingerprintSha256: "e082e9338a4639b41f95cbd199539efacc667278bbb50e4bf6b2138c8ea25f14",
			SubjectDn:         "C=US, ST=California, L=Mountain View, O=Google LLC, CN=www.google.com",
			IssuerDn:          "C=US, O=Google Trust Services, CN=GTS CA 1O1",
		},
		results{
			FingerprintSha256: "a9a1cd8961fecc1ae407a2d4e6d3ebce011213caa955bfbc9f855e0b6d086b45",
			SubjectDn:         "C=US, ST=California, L=Mountain View, O=Google LLC, CN=www.google.com",
			IssuerDn:          "C=US, O=Google Trust Services, CN=GTS CA 1O1",
		},
	}
	expectedSearch := &Search{
		Status:   "ok",
		Metadata: exMetadata,
		Results:  exResultsCert,
	}
	assert.Nil(t, err)
	assert.IsType(t, expectedSearch, search)
	assert.Equal(t, expectedSearch, search)
}
