package main

import (
	"flag"
	"fmt"
	"os"
	"stockprice/finance"
)

var format = flag.String("f", "term", "formatting: term or vim")

func init() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("Usage: stockprice [symbol], e.g. stockprice grpn")
		os.Exit(1)
	}
}

func main() {
	price, delta, percentage, err := finance.FindStockPrice(flag.Args()[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result string

	if *format == "vim" {
		result = formatForVim(price, delta, percentage)
	} else if *format == "term" {
		result = formatForTerminal(price, delta, percentage)
	}

	fmt.Println(result)
}

func formatForTerminal(price float64, delta float64, percentage string) string {
	var deltaFormatted string
	var percentageFormatted = percentage

	if delta > 0 {
		deltaFormatted = fmt.Sprintf("\x1b[32m+%.2f\x1b[0m", delta)
		percentageFormatted = fmt.Sprintf("\x1b[32m%s\x1b[0m", percentage)
	} else if delta < 0 {
		deltaFormatted = fmt.Sprintf("\x1b[31m%.2f\x1b[0m", delta)
		percentageFormatted = fmt.Sprintf("\x1b[31m%s\x1b[0m", percentage)
	} else {
		deltaFormatted = fmt.Sprintf("%.2f", delta)
	}

	return fmt.Sprintf("%.2f %s %s", price, deltaFormatted, percentageFormatted)
}

func formatForVim(price float64, delta float64, percentage string) string {
	var result string
	if delta > 0 {
		result = fmt.Sprintf(`echohl Normal
echo "%.2f"
echohl MoreMsg
echon " +%.2f %s"
echohl Normal`, price, delta, percentage)

	} else if delta < 0 {
		result = fmt.Sprintf(`echohl Normal
echo "%.2f"
echohl WarningMsg
echon " %.2f %s"
echohl Normal`, price, delta, percentage)
	} else {
		result = fmt.Sprintf(`echohl Normal
echo "%.2f"
echon " %.2f %s"
echohl Normal`, price, delta, percentage)
	}
	return result
}
