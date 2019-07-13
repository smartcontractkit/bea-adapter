package services

import (
	"github.com/bitly/go-simplejson"
	"net/http"
	"net/url"
)

type Client struct {
	uri *url.URL
}

func NewClient(uri ...string) (*Client, error) {
	if uri == nil {
		uri = []string{"https://apps.bea.gov/api/data/"}
	}

	u, err := url.Parse(uri[0])
	if err != nil {
		return nil, err
	}

	return &Client{u}, nil
}

func (client *Client) AddParam(key string, value string) {
	query := client.uri.Query()
	query.Set(key, value)
	client.uri.RawQuery = query.Encode()
}

func (client *Client) Execute() (*http.Response, error) {
	return http.Get(client.uri.String())
}

func (client *Client) GetData() ([]interface{}, error) {
	resp, err := client.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	return data.GetPath("BEAAPI", "Results", "Data").Array()
}
