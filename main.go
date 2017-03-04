package main

import (
	"fmt"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: stockprice [symbol], e.g. stockprice grpn")
		os.Exit(1)
	}
}

func main() {
	fmt.Println(findStockPrice(os.Args[1]))
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
