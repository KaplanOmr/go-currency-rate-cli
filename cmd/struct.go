package main

type CurrencyData struct {
	Sell string `json:"satis"`
	Buy  string `json:"alis"`
}

type CurrenciesRateData map[string]CurrencyData
