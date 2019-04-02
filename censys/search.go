package censys

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type scanType string

const (
	scanPath              = "/search/"
	CERTIFICATES scanType = scanPath + "certificates"
	IPV4         scanType = scanPath + "ipv4"
	WEBSITES     scanType = scanPath + "websites"
)

type Search struct {
	Status   string    `json:"status"`
	Metadata Metadata  `json:"metadata"`
	Results  []Results `json:"results"`
}

type Metadata struct {
	Count       int    `json:"count"`
	Query       string `json:"query"`
	BackendTime int    `json:"backend_time"`
	Page        int    `json:"page"`
	Pages       int    `json:"pages"`
}

type Results struct {
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
}

type SearchQuery struct {
	Query   string   `json:"query"`
	Page    int      `json:"page"`
	Fields  []string `json:"fields"`
	Flatten bool     `json:"flatten"`
}

func (c *Client) Search(ctx context.Context, query string, scantype scanType) (*Search, error) {
	var search Search
	s := make([]string, 0, 0)
	queryJSON := &SearchQuery{
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
