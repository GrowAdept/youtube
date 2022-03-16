// example without goroutines
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/piquette/finance-go/quote"
)

func main() {
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprint(w, "Symbol\tName\tMarket\tMarket Price\n")
	symbols := []string{"APPL", "2222.SR", "MSFT", "GOOG", "TSLA", "BRK-A", "NVDA", "FB", "TSM", "UNH", "TCEHY", "JNJ", "V", "WMT", "005930.KS", "JPM", "XOM", "PG", "600519.SS"}
	start := time.Now()
	for _, symbol := range symbols {
		// finance-go package actually has a list method, but using Get() for demonstration
		q, err := quote.Get(symbol)
		if err != nil {
			// Uh-oh.
			panic(err)
		}
		// Success!
		fmt.Fprint(w, symbol, "\t", q.ShortName, "\t", q.MarketID, "\t", q.RegularMarketPrice, "\n")
	}
	w.Flush()
	fmt.Print("\n")
	fmt.Println("duration:", time.Since(start))
}
