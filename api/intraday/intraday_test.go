package intraday_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsmtkk/solid-bassoon/api/intraday"
	myhttp "github.com/hsmtkk/solid-bassoon/http"
)

func TestQuery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := os.ReadFile("./example.json")
		assert.Nil(t, err)
		w.Write(resp)
	}))
	defer ts.Close()

	client := ts.Client()
	querier := intraday.New("apiKey", myhttp.New(ts.URL, client))
	got, err := querier.Query("IBM")
	assert.Nil(t, err)
	assert.Greater(t, len(got.MetaData), 0)
	assert.Greater(t, len(got.TimeSeries), 0)
}
