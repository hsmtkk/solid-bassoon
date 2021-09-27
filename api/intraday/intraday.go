package intraday

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hsmtkk/solid-bassoon/http"
)

type StockValue struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}

type Response struct {
	MetaData   map[string]string
	TimeSeries map[string]StockValue
}

type Querier interface {
	Query(symbol string) (Response, error)
}

type querierImpl struct {
	apiKey string
	getter http.Getter
}

func New(apiKey string, getter http.Getter) Querier {
	return &querierImpl{apiKey, getter}
}

func (q *querierImpl) Query(symbol string) (Response, error) {
	query := map[string]string{"function": "TIME_SERIES_INTRADAY", "symbol": symbol, "interval": "5min", "apikey": q.apiKey}
	respBytes, err := q.getter.Get(query)
	if err != nil {
		return Response{}, err
	}
	return q.parseResponse(respBytes)
}

func (q *querierImpl) parseResponse(jsonBytes []byte) (Response, error) {
	type stockValue struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	}

	type response struct {
		MetaData   map[string]string     `json:"Meta Data"`
		TimeSeries map[string]stockValue `json:"Time Series (5min)"`
	}

	resp := response{}
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return Response{}, fmt.Errorf("unmarshal failed; %w", err)
	}

	result := Response{}
	result.MetaData = resp.MetaData
	result.TimeSeries = map[string]StockValue{}
	for k, v := range resp.TimeSeries {
		stockValue := StockValue{}
		stockValue.Open, _ = strconv.ParseFloat(v.Open, 64)
		stockValue.High, _ = strconv.ParseFloat(v.High, 64)
		stockValue.Low, _ = strconv.ParseFloat(v.Low, 64)
		stockValue.Close, _ = strconv.ParseFloat(v.Close, 64)
		stockValue.Volume, _ = strconv.Atoi(v.Volume)
		result.TimeSeries[k] = stockValue
	}

	return result, nil
}
