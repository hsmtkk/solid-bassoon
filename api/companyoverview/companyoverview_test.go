package companyoverview_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hsmtkk/solid-bassoon/api/companyoverview"
	myhttp "github.com/hsmtkk/solid-bassoon/http"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL)
		resp, err := os.ReadFile("./example.json")
		assert.Nil(t, err)
		w.Write(resp)
	}))
	defer ts.Close()

	client := ts.Client()
	querier := companyoverview.New("apiKey", myhttp.New(ts.URL, client))
	want := companyoverview.Response{Symbol: "IBM", Name: "International Business Machines Corporation", EPS: "5.92"}
	got, err := querier.Query("IBM")
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
