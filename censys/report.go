package censys

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type reportType string

const (
	reportPath = "/report/"
	//CERTIFICATESREPORT serach for certificates
	CERTIFICATESREPORT reportType = reportPath + "certificates"
	//IPV4REPORT REPORT for ipv4
	IPV4REPORT reportType = reportPath + "ipv4"
	//WEBSITESREPORT REPORT for websites
	WEBSITESREPORT reportType = reportPath + "websites"
)

//Report struct that represents the returned data from the report endpoint
type Report struct {
	Status   string          `json:"status"`
	Metadata reportMetadata  `json:"metadata"`
	Results  []reportResults `json:"results"`
}

type reportResults struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

type reportMetadata struct {
	Count            int    `json:"count"`
	BackendTime      int    `json:"backend_time"`
	NonnullCount     int    `json:"nonnull_count"`
	OtherResultCount int    `json:"other_result_count"`
	Buckets          int    `json:"buckets"`
	ErrorBound       int    `json:"error_bound"`
	Query            string `json:"query"`
}

//ReportQuery Struct containg the query passed to get the report from the API
//query: [required string]
//The query to be executed. For example, 80.http.get.headers.server: nginx.
//field: [required string]
//The field you are running a breakdown on in "dot notation", e.g. location.country_code.
//buckets: [optional int]
//The maximum number of values to be returned in the report. Maximum: 500. Default: 50.
type ReportQuery struct {
	Query   string `json:"query"`
	Field   string `json:"field"`
	Buckets int    `json:"buckets"`
}

//GetReport calls the report api
func (c *Client) GetReport(ctx context.Context, typeReport reportType, query ReportQuery) (*Report, error) {
	var report Report
	q, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(http.MethodPost, string(typeReport), bytes.NewReader(q))
	if err != nil {
		return nil, err
	}
	if err := c.Do(ctx, req, &report); err != nil {
		return nil, err
	}
	return &report, nil
}
