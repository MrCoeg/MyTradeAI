package main

import (
	"data-collection/domain"
	"data-collection/pkg"
	"time"
)

func main() {
	cryptocoins := domain.CryptoCoins{}
	records := [][]string{}

	for {
		pkg.GetJSONFromUrl(&cryptocoins, "https://indodax.com/api/tickers")
		records = append(records, pkg.UnpackedTickerToString(cryptocoins.Coins))
		pkg.WriteCSV(records)
		time.Sleep(5 * time.Second)
	}

}
