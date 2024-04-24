package client

import (
	rest "github.com/go-resty/resty/v2"
)

type RestClient struct {
	client rest.Client
}

func NewRestClient() RestClient {
	return RestClient{client: *rest.New()}

}

func (r RestClient) GenRestConn() (rest.Client, error) {
	return r.client, nil
}
