package main

const API = "genelpara.com"

type CurrencyData struct {
	Sell    string `json:"satis"`
	Buy     string `json:"alis"`
	Changes string `json:"degisim"`
}

type CurrenciesRateData map[string]CurrencyData
