package finance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestFindStockPriceByUrl(t *testing.T) {
	ms := mockServer("../stubs/sample_stock_price.html")
	defer ms.Close()

	testStockPrice, testDelta, testPercentage, _ := findStockPriceByUrl(ms.URL)

	expectedPrice := 1004.0203092013
	if testStockPrice == expectedPrice {
		t.Logf("\tShould find a correct stock price \"%f\" %v", expectedPrice, checkMark)
	} else {
		t.Errorf("\tShould find a correct stock price \"%f\", but found \"%f\" %v",
			expectedPrice, testStockPrice, ballotX)
	}

	expectedDelta := 1000.03
	if testDelta == expectedDelta {
		t.Logf("\tShould find a correct price change \"%f\" %v", expectedDelta, checkMark)
	} else {
		t.Errorf("\tShould find a correct price change \"%f\", but found \"%f\" %v",
			expectedDelta, testDelta, ballotX)
	}

	expectedPercentage := "+0.74%"
	if testPercentage == expectedPercentage {
		t.Logf("\tShould find a correct percentage change \"%f\" %v", expectedPercentage, checkMark)
	} else {
		t.Errorf("\tShould find a correct percentage change \"%f\", but found \"%f\" %v",
			expectedPercentage, testPercentage, ballotX)
	}

	ms.Close()
}

func mockServer(stubbedHtmlFile string) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		feed, _ := ioutil.ReadFile(stubbedHtmlFile)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, string(feed))
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
