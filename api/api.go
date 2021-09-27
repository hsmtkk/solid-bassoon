package api

import (
	"github.com/hsmtkk/solid-bassoon/api/companyoverview"
	"github.com/hsmtkk/solid-bassoon/api/intraday"
	"github.com/hsmtkk/solid-bassoon/http"
)

type AlphaVantageClient interface {
	CompanyOverview(symbol string) (companyoverview.Response, error)
	Intraday(symbol string) (intraday.Response, error)
}

type clientImpl struct {
	apiKey string
	getter http.Getter
}

func New(apiKey string, getter http.Getter) AlphaVantageClient {
	return &clientImpl{apiKey, getter}
}

func (c *clientImpl) CompanyOverview(symbol string) (companyoverview.Response, error) {
	return companyoverview.New(c.apiKey, c.getter).Query(symbol)
}

func (c *clientImpl) Intraday(symbol string) (intraday.Response, error) {
	return intraday.New(c.apiKey, c.getter).Query(symbol)
}
