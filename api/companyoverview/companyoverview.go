package companyoverview

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/solid-bassoon/http"
)

type Querier interface {
	Query(symbol string) (Response, error)
}

type Response struct {
	Symbol string
	Name   string
	EPS    string
}

type querierImpl struct {
	apiKey string
	getter http.Getter
}

func New(apiKey string, client http.Getter) Querier {
	return &querierImpl{apiKey, client}
}

func (q *querierImpl) Query(symbol string) (Response, error) {
	query := map[string]string{"function": "OVERVIEW", symbol: symbol, "apiKey": q.apiKey}
	jsonBytes, err := q.getter.Get(query)
	if err != nil {
		return Response{}, err
	}
	return q.parseResponse(jsonBytes)
}

func (q *querierImpl) parseResponse(jsonBytes []byte) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return Response{}, fmt.Errorf("unmarshal failed; %w", err)
	}
	return resp, nil
}
