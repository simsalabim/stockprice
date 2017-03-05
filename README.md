# Stock Price Reporter

A CLI tool that looks up stock prices.

Now you can check stock price right in your console without losing focus on what matters :computer:

```shell
$ stockprice
Usage: stockprice [symbol], e.g. stockprice grpn
$ stockprice grpn
4.02
```

You can create aliases for frequently used companies and add them to your .bashrc/.zshrc file:
```shell
# .zshrc
alias fb="stockprice fb"

# further usage in console:
$ fb
137.17
```

# Installation

Download an executable for your OS: [OSX](https://github.com/simsalabim/stockprice/blob/master/bin/stockprice), [Linux](https://github.com/simsalabim/stockprice/blob/master/bin/stockprice-lin), [Windows](https://github.com/simsalabim/stockprice/blob/master/bin/stockprice.exe) and place it in a directory listed in your $PATH.

# Development

For development you will need a Go language compiler installed. I used Go v1.8.

To rebuild the binaries for 3 main platforms run `./build_executables`.

Run tests with `go test -v`
