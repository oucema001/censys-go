package censys

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type scanType string

const (
	scanPath = "/search/"
	//CERTIFICATES serach for certificates
	CERTIFICATES scanType = scanPath + "certificates"
	//IPV4 search for ipv4
	IPV4 scanType = scanPath + "ipv4"
	//WEBSITES search for websites
	WEBSITES scanType = scanPath + "websites"
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
	Longitude            string   `json:"location.longitude"`
	Latitude             string   `json:"location.latitude"`
	City                 string   `json:"location.city"`
	RegisteredCountyCode string   `json:"location.registered_country_code"`
	CountryCode          string   `json:"location.country_code"`
	Province             string   `json:"location.province"`
	PostalCode           string   `json:"location.postal_code"`
	TimeZone             string   `json:"location.timezone"`
	//Certificate Results
	FingerprintSha256 string `json:"parsed.fingerprint_sha256"`
	SubjectDn         string `json:"parsed.subject_dn string"`
	IssuerDn          string `json:"parsed.issuer_dn "`
	//Website Results
	Domain    string `json:"domain"`
	AlexaRank string `json:"alexa_rank"`
}

type searchQuery struct {
	Query   string   `json:"query"`
	Page    int      `json:"page"`
	Fields  []string `json:"fields"`
	Flatten bool     `json:"flatten"`
}

//Search searches a query using the API by specifying a query and a Scan Type
func (c *Client) Search(ctx context.Context, query string, scantype scanType) (*Search, error) {
	var search Search
	s := make([]string, 0, 0)
	queryJSON := &searchQuery{
		Query:   query,
		Page:    1,
		Fields:  s,
		Flatten: true,
	}
	b, err := json.Marshal(queryJSON)
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
