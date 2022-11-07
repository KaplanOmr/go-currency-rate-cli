package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	currencyFlag := flag.String("cur", "all", "a string")
	flag.Parse()

	currenciesData := getCurrencies()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Currency", "Buy", "Shell", "Changes")
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(columnFmt)

	for currency, info := range currenciesData {
		if *currencyFlag == "all" || *currencyFlag == "" {
			tbl.AddRow(currency, info.Buy, info.Sell, getChangesWithColor(info.Changes))
		} else if *currencyFlag == strings.ToLower(currency) {
			tbl.AddRow(currency, info.Buy, info.Sell, getChangesWithColor(info.Changes))
			break
		} else {
			fmt.Println("INCORRECT_CURRENCY")
			return
		}
	}

	tbl.Print()

	fmt.Printf("\nDate: %s", time.Now().Format("2006.01.02 15:04:05"))
	fmt.Println("\nResource: https://" + API)
}

func getCurrencies() CurrenciesRateData {
	resp, err := http.Get("https://api." + API + "/embed/doviz.json")
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

func getChangesWithColor(rate string) string {
	parsedRate, err := strconv.ParseFloat(rate, 3)

	if err != nil {
		panic(err)
	}

	if parsedRate < 0 {
		return fmt.Sprintf("\u001b[41;1m %s \u001b[0m", rate)
	} else if parsedRate > 0 {
		return fmt.Sprintf("\u001b[42;1m %s \u001b[0m", rate)
	}

	return rate
}
