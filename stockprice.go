package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func findStockPrice(symbol string) float64 {
	return findStockPriceByUrl(stockPriceUrl(symbol))
}

func findStockPriceByUrl(stockPriceUrl string) float64 {
	doc, err := goquery.NewDocument(stockPriceUrl)
	if err != nil {
		fmt.Println(err)
	}

	selection := doc.Find("#market-data-div .pr")

	if len(selection.Nodes) == 0 {
		fmt.Println("Your search produces no matches.")
		os.Exit(-1)
	}

	stock_price_str := strings.TrimSpace(selection.Text())
	stock_price, err := strconv.ParseFloat(stock_price_str, 64)
	checkError(err)

	return stock_price
}

func stockPriceUrl(symbol string) string {
	return "https://www.google.com/finance?q=" + url.QueryEscape(symbol)
}
