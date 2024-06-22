package crypto

import "fmt"

// PrintBar prints the details of a Bar.
func PrintBar(bar Bar) {
	if bar.Symbol == "" && bar.Open == 0 && bar.Close == 0 && bar.High == 0 && bar.Low == 0 && bar.Timestamp.IsZero() {
		return // Skip empty records
	}
	fmt.Printf("Symbol: %s | Open: %.2f | High: %.2f | Low: %.2f | Close: %.2f | Timestamp (UTC): %s\n",
		bar.Symbol, bar.Open, bar.High, bar.Low, bar.Close, bar.Timestamp.Format("2006-01-02 15:04:05"))
}
