package censys

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	scanPath = "/search/certificates"
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
	IP        string   `json:"ip"`
	Protocols []string `json:"protocols"`
}

type SearchQuery struct {
	Query   string   `json:"query"`
	Page    int      `json:"page"`
	Fields  []string `json:"fields"`
	Flatten bool     `json:"flatten"`
}

func (c *Client) Search(ctx context.Context, query string) (*Search, error) {
	var search Search
	//body := neturl.Values{}
	//body.Add("query", query)
	queryJson := &SearchQuery{
		Query:   query,
		Page:    1,
		Flatten: true,
	}
	b, err := json.Marshal(queryJson)
	if err != nil {
		return nil, err
	}
	//body.Add(b)
	//fmt.Println(strings.NewReader(body.Encode()))

	req, err := c.NewRequest(http.MethodPost, scanPath, nil, strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	if err := c.Do(ctx, req, &search); err != nil {
		return nil, err
	}
	fmt.Printf("%v+", &req.Body)
	return &search, err
}
