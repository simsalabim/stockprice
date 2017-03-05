package main

import (
	"fmt"
	"github.com/simsalabim/stockprice/finance"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: stockprice [symbol], e.g. stockprice grpn")
		os.Exit(1)
	}
}

func main() {
	fmt.Println(finance.FindStockPrice(os.Args[1]))
}
