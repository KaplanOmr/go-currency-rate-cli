package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	currenciesData := getCurrencies()

	for currency, info := range currenciesData {
		line := fmt.Sprintf("Currency: %10s Shell: %10s₺ Buy: %10s₺", currency, info.Sell, info.Buy)
		fmt.Println(line)
	}
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
