package censys

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReport(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()

	expectedIP := "80.http.get.headers.server: nginx"
	mux.HandleFunc(string(IPV4REPORT), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var query ReportQuery
		dec := json.NewDecoder(r.Body)
		dec.Decode(&query)
		assert.Equal(t, query.Query, expectedIP)
		w.Write(getStubs(t, "report"))
	})
	query := ReportQuery{
		Query:   "80.http.get.headers.server: nginx",
		Field:   "location.country",
		Buckets: 10,
	}
	rep, err := client.GetReport(context.Background(), IPV4REPORT, query)
	assert.Nil(t, err)
	report := &Report{

		Status: "ok",
		Results: []reportResults{
			{
				Key:      "80/http",
				DocCount: 60360211,
			},
			{
				Key:      "443/https",
				DocCount: 55055587,
			},
			{
				Key:      "7547/cwmp",
				DocCount: 38971729,
			},
			{
				Key:      "21/ftp",
				DocCount: 13434960,
			},
			{
				Key:      "25/smtp",
				DocCount: 8967614,
			},
		},
		Metadata: reportMetadata{
			Count:            130517957,
			BackendTime:      1599,
			NonnullCount:     206619766,
			OtherResultCount: 29829665,
			Buckets:          5,
			ErrorBound:       0,
			Query:            "*",
		},
	}
	assert.IsType(t, rep, report)
	assert.Equal(t, rep, report)
}
