#!/bin/bash
if [ `uname -s` == "Darwin" ];then
  curl -L -s https://github.com/simsalabim/stockprice/raw/master/bin/stockprice -o /usr/local/bin/stockprice
elif [ `uname -s` == "Linux"];then
  curl -L -s https://github.com/simsalabim/stockprice/raw/master/bin/stockprice-lin -o /usr/local/bin/stockprice
fi
chmod a+x /usr/local/bin/stockprice
echo "\`stockprice\` is now installed and available on your system."
