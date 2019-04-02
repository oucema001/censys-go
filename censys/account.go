package censys

import (
	"context"
	"fmt"
	"net/http"
)

const (
	accountURL = "/account"
)

type Profile struct {
	Login      string `json:"login"`
	Email      string `json:"email"`
	FirstLogin string `json:"first_login"`
	LastLogin  string `json:"last_login"`
	Quota      Quota  `json:"quota"`
}

type Quota struct {
	Used      int    `json:"used"`
	ResetsAt  string `json:"resets_at"`
	Allowance int    `json:"allowance"`
}

func (c *Client) GetProfile(ctx context.Context) (*Profile, error) {
	var profile Profile
	req, err := c.NewRequest(http.MethodGet, accountURL, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &profile); err != nil {
		return nil, err
	}
	fmt.Printf("%v+", err)

	return &profile, nil
}
