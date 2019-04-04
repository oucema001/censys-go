package censys

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	setUpTestServe()
	defer tearTestServer()
	mux.HandleFunc(accountURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write(getStubs(t, accountURL))
	})

	quota := Quota{
		Used:      30,
		ResetsAt:  "2019-04-15 15:51:15",
		Allowance: 250,
	}
	expectedProfile := &Profile{
		Login:      "john",
		FirstLogin: "2016-09-15 15:52:14",
		LastLogin:  "2019-04-04 13:06:23.972131",
		Email:      "john@example.com",
		Quota:      quota,
	}
	profile, err := client.GetProfile(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.IsType(t, expectedProfile, profile)
	assert.EqualValues(t, expectedProfile, profile)
}
