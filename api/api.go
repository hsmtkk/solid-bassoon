package api

import (
	"net/http"

	"github.com/hsmtkk/solid-bassoon/api/companyoverview"
)

type AlphaVantageClient interface {
	CompanyOverview(symbol string) (companyoverview.Response, error)
}

type clientImpl struct {
	apiKey string
	client *http.Client
}

func New(apiKey string, client *http.Client) AlphaVantageClient {
	return &clientImpl{apiKey, client}
}

func (c *clientImpl) CompanyOverview(symbol string) (companyoverview.Response, error) {
	return companyoverview.New(c.apiKey, c.client).Query(symbol)
}
