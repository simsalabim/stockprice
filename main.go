package main

import (
	"fmt"
	"os"
	"stockprice/finance"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: stockprice [symbol], e.g. stockprice grpn")
		os.Exit(1)
	}
}

func main() {
	price, delta, percentage, err := finance.FindStockPrice(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var deltaFormatted, percentageFormatted string

	if delta > 0 {
		deltaFormatted = fmt.Sprintf("\x1b[32m+%.2f\x1b[0m", delta)
		percentageFormatted = fmt.Sprintf("\x1b[32m%s\x1b[0m", percentage)
	} else if delta < 0 {
		deltaFormatted = fmt.Sprintf("\x1b[31m%.2f\x1b[0m", delta)
		percentageFormatted = fmt.Sprintf("\x1b[31m%s\x1b[0m", percentage)
	}

	fmt.Println(fmt.Sprintf("%.2f", price), deltaFormatted, percentageFormatted)
}
