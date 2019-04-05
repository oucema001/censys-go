package censys

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	mux.HandleFunc(dataPath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet)

		w.Write(getStubs(t, "data"))
	})

	lastresult := LatestResult{
		Timestamp:  "20151013T120047",
		Name:       "20151013T1200",
		DetailsURL: "https://censys.io/api/v1/data/certificates/20151013T1200",
	}
	certificates := Certificates{
		Description:  "[...]",
		DetailsURL:   "https://censys.io/api/v1/data/certificates",
		LatestResult: lastresult,
		ID:           "certificates",
		Name:         "certificates",
	}
	domain := Domain{
		Description:  "[...]",
		DetailsURL:   "https://censys.io/api/v1/data/domain",
		LatestResult: lastresult,
		ID:           "domain",
		Name:         "domain",
	}
	ipv4 := Ipv4{
		Description:  "[...]",
		DetailsURL:   "https://censys.io/api/v1/data/ipv4",
		LatestResult: lastresult,
		ID:           "ipv4",
		Name:         "ipv4",
	}
	primary := PrimarySeries{
		Certificates: certificates,
		Domain:       domain,
		Ipv4:         ipv4,
	}

	lastresult1 := LatestResult{
		Timestamp:  "20150930T005634",
		Name:       "20150930T0056",
		DetailsURL: "https://censys.io/api/v1/data/22-ssh-banner-full_ipv4/20150930T0056",
	}
	lastresult2 := LatestResult{
		Timestamp:  "20150727T060848",
		Name:       "20150727T0608",
		DetailsURL: "https://censys.io/api/v1/data/443-https-ssl3-alexa_top1mil/20150727T0608",
	}
	two2ss := Two2SSHBannerFullIpv4{
		Subprotocol:  "banner",
		Description:  "[...]",
		Protocol:     "ssh",
		Name:         "22-ssh-banner-full_ipv4",
		DetailsURL:   "https://censys.io/api/v1/data/22-ssh-banner-full_ipv4",
		LatestResult: lastresult1,
		Destination:  "full_ipv4",
		ID:           "22-ssh-banner-full_ipv4",
		Port:         22,
	}
	four := Four43HTTPSSsl3AlexaTop1Mil{
		Subprotocol:  "ssl3",
		Description:  "[...]",
		Protocol:     "https",
		Name:         "443-https-ssl3-alexa_top1mil",
		DetailsURL:   "https://censys.io/api/v1/data/443-https-ssl3-alexa_top1mil",
		LatestResult: lastresult2,
		Destination:  "alexa_top1mil",
		ID:           "443-https-ssl3-alexa_top1mil",
		Port:         443,
	}
	rawseries := RawSeries{
		Two2SSHBannerFullIpv4:       two2ss,
		Four43HTTPSSsl3AlexaTop1Mil: four,
	}
	expectedData := &Data{
		PrimarySeries: primary,
		RawSeries:     rawseries,
	}
	data, err := client.GetData(context.Background())
	assert.Nil(t, err)
	assert.IsType(t, data, expectedData)
	assert.Equal(t, data, expectedData)

}
