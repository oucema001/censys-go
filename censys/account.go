package censys

import (
	"context"
	"net/http"
)

const (
	accountURL = "/account"
)

// Profile represents the struct containing the account used for the API calls
type Profile struct {
	Login      string `json:"login"`       // user login
	Email      string `json:"email"`       //user email
	FirstLogin string `json:"first_login"` //The user's first login date
	LastLogin  string `json:"last_login"`  //The user's last login date
	Quota      Quota  `json:"quota"`       // The API quota already used by this user
}

// Quota represents The api quota limit
type Quota struct {
	Used      int    `json:"used"`      //The number of api calls already used since last reset
	ResetsAt  string `json:"resets_at"` //Data at which the api quota will reset
	Allowance int    `json:"allowance"` //Number of API calls allowed for the user at the current plan
}

// GetProfile makes an API call to retrieve account details through the accountURL endpoint and returns a Profile struct
func (c *Client) GetProfile(ctx context.Context) (*Profile, error) {
	var profile Profile
	req, err := c.NewRequest(http.MethodGet, accountURL, nil)
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}
