package finance

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

func FindStockPrice(symbol string) (float64, float64, string, error) {
	return findStockPriceByUrl(stockPriceUrl(symbol))
}

func findStockPriceByUrl(stockPriceUrl string) (float64, float64, string, error) {
	doc, err := goquery.NewDocument(stockPriceUrl)
	if err != nil {
		return -1, -1, "", errors.New("Your search produces no matches.")
	}

	selection := doc.Find("#market-data-div .pr")

	if len(selection.Nodes) == 0 {
		return -2, -2, "", errors.New("Your search produces no matches.")
	}

	changes := doc.Find(".id-price-change .ch.bld")
	delta, _ := strconv.ParseFloat(strings.Replace(changes.Find("span").First().Text(), ",", "", -1), 64)
	percentageStr := changes.Find("span").Last().Text()

	stockPriceStr := strings.TrimSpace(strings.Replace(selection.Text(), ",", "", -1))
	stockPrice, err := strconv.ParseFloat(stockPriceStr, 64)
	if err != nil {
		return -3, -3, "", err
	}

	return stockPrice, delta, percentageStr, nil
}

func stockPriceUrl(symbol string) string {
	return "https://www.google.com/finance?q=" + url.QueryEscape(symbol)
}
