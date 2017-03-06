package finance

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

func FindStockPrice(symbol string) (float64, error) {
	return findStockPriceByUrl(stockPriceUrl(symbol))
}

func findStockPriceByUrl(stockPriceUrl string) (float64, error) {
	doc, err := goquery.NewDocument(stockPriceUrl)
	if err != nil {
		return -1, errors.New("Your search produces no matches.")
	}

	selection := doc.Find("#market-data-div .pr")

	if len(selection.Nodes) == 0 {
		return -2, errors.New("Your search produces no matches.")
	}

	stock_price_str := strings.TrimSpace(selection.Text())
	stock_price, err := strconv.ParseFloat(stock_price_str, 64)
	if err != nil {
		return -3, err
	}

	return stock_price, nil
}

func stockPriceUrl(symbol string) string {
	return "https://www.google.com/finance?q=" + url.QueryEscape(symbol)
}
