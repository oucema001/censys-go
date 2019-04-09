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

	meta := Metadata{}
	star := Starttls{

		Ehlo:     "250-mx0b-00262c01.pphosted.com Hello CLIENT_HOSTNAME [CLIENT_IP] (may be forged), pleased to meet you\r\n250 ENHANCEDSTATUSCODES",
		Starttls: "500 5.5.1 Command unrecognized: \"STARTTLS\"",
		Banner:   "554 Blocked - see https://ipcheck.proofpoint.com/?ip=CLIENT_IP",
		Metadata: meta,
	}

	smtp := SMTP{
		Starttls: star,
	}
	num25 := Num25{
		SMTP: smtp,
	}

	ports :=
		[]int{
			80,
			25,
			443,
		}
	protocols :=
		[]string{
			"443/https_www",
			"443/https",
			"80/http_www",
			"25/smtp",
			"80/http",
		}
	tags := []string{
		"http",
		"smtp",
		"https",
	}

	expectedRes := &View{
		Num25:     num25,
		Ports:     ports,
		Protocols: protocols,
		Domain:    "gamespot.com",
		AlexaRank: 389,
		Tags:      tags,
	}

	assert.Nil(t, err)
	assert.IsType(t, expectedRes, view)
	assert.EqualValues(t, expectedRes, view)
}
