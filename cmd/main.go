package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	currencyFlag := flag.String("cur", "all", "a string")
	flag.Parse()

	currenciesData := getCurrencies()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

	tbl := table.New("Currency", "Buy", "Shell", "Changes").WithHeaderFormatter(headerFmt)

	for currency, info := range currenciesData {
		if *currencyFlag == "all" {
			tbl.AddRow(currency, info.Buy, info.Sell, info.Changes)
		} else if *currencyFlag == strings.ToLower(currency) {
			tbl.AddRow(currency, info.Buy, info.Sell, info.Changes)
			break
		}
	}

	tbl.Print()
}

func getCurrencies() CurrenciesRateData {
	resp, err := http.Get("")
	if err != nil {
		panic("CURRENCIES_DATA_FETCH_ERROR")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("CURRENCIES_DATA_BODY_READ_ERROR")
	}

	var currenciesRateData CurrenciesRateData

	if err = json.Unmarshal(body, &currenciesRateData); err != nil {
		panic("CURRENCIES_DATA_JSON_PARSE_ERROR")
	}

	return currenciesRateData
}
