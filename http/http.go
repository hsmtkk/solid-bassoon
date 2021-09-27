package http

import (
	"fmt"
	"io"
	"net/http"
)

type Getter interface {
	Get(query map[string]string) ([]byte, error)
}

type getterImpl struct {
	url    string
	client *http.Client
}

func New(url string, client *http.Client) Getter {
	return &getterImpl{url, client}
}

func (g *getterImpl) Get(query map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, g.url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest failed; %w", err)
	}

	params := req.URL.Query()
	for k, v := range query {
		params.Add(k, v)
	}
	req.URL.RawQuery = params.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP Get failed; %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll failed; %w", err)
	}
	return respBytes, nil
}
