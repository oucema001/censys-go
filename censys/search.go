package censys

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type scanType string

const (
	searchPath = "/search/"
	//CERTIFICATESSEARCH serach for certificates
	CERTIFICATESSEARCH scanType = searchPath + "certificates"
	//IPV4SEARCH search for ipv4
	IPV4SEARCH scanType = searchPath + "ipv4"
	//WEBSITESSEARCH search for websites
	WEBSITESSEARCH scanType = searchPath + "websites"
)

//Search struct that contains the result returned by Search
type Search struct {
	Status   string    `json:"status"`
	Metadata metadata  `json:"metadata"`
	Results  []results `json:"results"`
}

type metadata struct {
	Count       int    `json:"count"`
	Query       string `json:"query"`
	BackendTime int    `json:"backend_time"`
	Page        int    `json:"page"`
	Pages       int    `json:"pages"`
}

type results struct {
	//IPV4 results
	IP                   string   `json:"ip"`
	Protocols            []string `json:"protocols"`
	Country              string   `json:"location.country"`
	RegisteredCountry    string   `json:"location.registered_country"`
	Longitude            float64  `json:"location.longitude"`
	Latitude             string   `json:"location.latitude"`
	City                 string   `json:"location.city"`
	RegisteredCountyCode string   `json:"location.registered_country_code"`
	CountryCode          string   `json:"location.country_code"`
	Province             string   `json:"location.province"`
	PostalCode           string   `json:"location.postal_code"`
	TimeZone             string   `json:"location.timezone"`
	Continent            string   `json:"location.continent"`
	//Certificate Results
	FingerprintSha256 string `json:"parsed.fingerprint_sha256"`
	SubjectDn         string `json:"parsed.subject_dn"`
	IssuerDn          string `json:"parsed.issuer_dn"`
	//Website Results
	Domain    string `json:"domain"`
	AlexaRank int    `json:"alexa_rank"`
}

//Search Query constructs a query for the search api
type SearchQuery struct {
	Query   string   `json:"query"`
	Page    int      `json:"page"`
	Fields  []string `json:"fields"`
	Flatten bool     `json:"flatten"`
}

//Search searches a query using the API by specifying a query and a Scan Type
func (c *Client) Search(ctx context.Context, query *SearchQuery, scantype scanType) (*Search, error) {
	var search Search
	b, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(http.MethodPost, string(scantype), nil, strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	if err := c.Do(ctx, req, &search); err != nil {
		return nil, err
	}
	return &search, err
}
