# Stock Price Reporter

A CLI tool that looks up stock prices.
Now you can check stock price right in your console or in VIM without losing focus on what matters :computer:

```shell
$ stockprice grpn
4.02
```

## VIM Integration
As amazing as it sounds, you may not check the stock price of your interest right in your code editor.

```vim
" ~/.vimrc
function! Stockprice(symbol)
  echo system('stockprice ' . a:symbol)
endfunction

map <leader>fb :call Stockprice('fb')<cr>
```

## Alternative VIM Integration
If you're only interested in VIM integration and not in the system binary, the following web version that
uses this very [package](https://github.com/simsalabim/rebot.works/blob/069fdf0cb33169026b4fea83483747f9f67af23e/main.go#L6) might work for you:
```vim
function! Stockprice(symbol)
  echo system("curl -s http://rebot.works/stockprice/" . a:symbol)
endfunction
```

## More fun
What can be more fun than system aliases? You can create some for frequently checked stocks:
```shell
# .zshrc/.bashrc
alias fb="stockprice fb"

# further usage in console:
$ fb
137.17
```

## Installation

Download an executable for your OS: [OSX](https://github.com/simsalabim/stockprice/blob/master/bin/stockprice), [Linux](https://github.com/simsalabim/stockprice/blob/master/bin/stockprice-lin), [Windows](https://github.com/simsalabim/stockprice/blob/master/bin/stockprice.exe) and place it in a directory listed in your $PATH.

## Development

For development you will need a Go language compiler installed. I used Go v1.8.

To rebuild the binaries for 3 main platforms run `./build_executables`.

Run tests with `cd finance && go test -v`
