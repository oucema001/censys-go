package censys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://censys.io/api/v1"
)

//Client represents Censys http Client
type Client struct {
	APIID     string
	APISecret string
	BaseURL   string
	Debug     string
	Client    *http.Client
}

//NewClient Creates new censys Client
func NewClient(client *http.Client, apiID string, apiSecret string) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		APIID:     apiID,
		APISecret: apiSecret,
		BaseURL:   baseURL,
		Client:    client,
	}
}

//NewRequest Creates a new request to censys API
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
	req.SetBasicAuth(c.APIID, c.APISecret)
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	}
	return req, nil
}

//Do executes the API request
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
	return c.parseResponse(destination, resp.Body)
}

func (c *Client) parseResponse(destination interface{}, body io.Reader) error {
	var err error
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	s := buf.String() //
	fmt.Printf("%v+", string(s))
	if w, ok := destination.(io.Writer); ok {
		_, err = io.Copy(w, body)
	} else {
		dec := json.NewDecoder(strings.NewReader(s))
		dec.Decode(&destination)
	}
	return err
}
