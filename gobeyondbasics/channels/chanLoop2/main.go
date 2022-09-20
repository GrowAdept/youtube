// example without goroutines
package main

import (
	"fmt"
	"os"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
)

var wg = sync.WaitGroup{}

func getQuote(symbol string, ch chan *finance.Quote) {
	q, err := quote.Get(symbol)
	if err != nil {
		// Uh-oh.
		fmt.Println(err)
	}
	// Success!
	ch <- q
	wg.Done()
}

func main() {
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprint(w, "Symbol\tName\tMarket\tMarket Price\n")
	// WARNING: code does not check for no result found situations, if a company changes stock symbol, you will get runtime error
	symbols := []string{"APPL", "2222.SR", "MSFT", "GOOG", "TSLA", "BRK-A", "NVDA", "META", "TSM", "UNH", "TCEHY", "JNJ", "V", "WMT", "005930.KS", "JPM", "XOM", "PG", "600519.SS"}
	quotes := make(chan *finance.Quote, len(symbols))
	start := time.Now()
	for _, symbol := range symbols {
		wg.Add(1)
		// finance-go package actually has a list method, but using Get() for demonstration
		go getQuote(symbol, quotes)
	}
	for i := 0; i < len(symbols); i++ {
		q := <-quotes
		/*
			Use fmt.Println(q) instead, this will print each line as they are ready,
			tabwriter will wait for all lines are ready and print all at once.  If experimenting
			with the code you want to see actuall delays, not the tabwriter delay, sorry for inconvenience
		*/
		// fmt.Fprint(w, q.Symbol, "\t", q.ShortName, "\t", q.MarketID, "\t", q.RegularMarketPrice, "\n")
		fmt.Println(q)
	}
	w.Flush()
	wg.Wait()
	fmt.Print("\n")
	fmt.Println("duration:", time.Since(start))
}

// without goroutines    2781ms
// with goroutines        405ms
