package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hsmtkk/solid-bassoon/api"
	myhttp "github.com/hsmtkk/solid-bassoon/http"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use: "query",
}

var companyOverview = &cobra.Command{
	Use:  "companyoverview symbol",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		apiKey := getApiKey()
		runCompanyOverview(apiKey, symbol)
	},
}

var intraday = &cobra.Command{
	Use:  "intraday symbol",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		apiKey := getApiKey()
		runIntraday(apiKey, symbol)
	},
}

func init() {
	command.AddCommand(companyOverview)
	command.AddCommand(intraday)
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getApiKey() string {
	key := os.Getenv("API_KEY")
	if key == "" {
		log.Fatal("API_KEY environment variable must be defined")
	}
	return key
}

func runCompanyOverview(apiKey, symbol string) {
	getter := myhttp.New("https://www.alphavantage.co/query", http.DefaultClient)
	app := api.New(apiKey, getter)
	resp, err := app.CompanyOverview(symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func runIntraday(apiKey, symbol string) {
	getter := myhttp.New("https://www.alphavantage.co/query", http.DefaultClient)
	app := api.New(apiKey, getter)
	resp, err := app.Intraday(symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
