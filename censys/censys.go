package censys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://censys.io/api/v1"
)

type Client struct {
	ApiID     string
	ApiSecret string
	BaseURL   string
	Debug     string
	Client    *http.Client
}

func NewClient(client *http.Client, apiID string, apiSecret string) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		ApiID:     apiID,
		ApiSecret: apiSecret,
		BaseURL:   baseURL,
		Client:    client,
	}
}

func (c *Client) NewRequest(method string, path string, params interface{}, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	return c.newRequest(method, u, params, body)
}

func (c *Client) newRequest(method string, u *url.URL, params interface{}, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.ApiID, c.ApiSecret)
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	}
	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, destination interface{}) error {
	req = req.WithContext(ctx)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return err
	}
	if destination == nil {
		return nil
	}
	//	fmt.Printf("%v+", &resp.Body)
	return c.ParseResponse(destination, resp.Body)
}

func (c *Client) ParseResponse(destination interface{}, body io.Reader) error {
	var err error
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	s := buf.String() //
	fmt.Printf("%v+", string(s))
	if w, ok := destination.(io.Writer); ok {
		_, err = io.Copy(w, body)
	} else {
		dec := json.NewDecoder(body)
		dec.Decode(&destination)
	}
	return err
}
