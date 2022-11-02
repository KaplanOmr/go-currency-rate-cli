package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	currenciesData := getCurrencies()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

	tbl := table.New("Currency", "Buy", "Shell", "Changes").WithHeaderFormatter(headerFmt)

	for currency, info := range currenciesData {
		tbl.AddRow(currency, info.Buy, info.Sell, info.Changes)
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
