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
