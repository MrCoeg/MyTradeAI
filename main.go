package main

import (
	"time"
	"trading-ai/domain"
	"trading-ai/pkg"
)

func main() {

	// Init
	cryptocoins := domain.CryptoCoins{}
	records := [][]string{}

	for {
		pkg.GetJSONFromUrl(&cryptocoins, "https://indodax.com/api/summaries")
		records = append(records, pkg.UnpackedTickerToString(cryptocoins.Coins))
		pkg.WriteCSV(records)
		time.Sleep(30 * time.Second)
	}

}

func UpdateChart(charts *domain.Chart, cryptocoins *domain.CryptoCoins, frameRate int, interval int, ticker domain.Ticker) (string, int) {
	charts.Candles[len(charts.Candles)-1].Update(ticker)
	value, status := charts.CreateNewCandle(frameRate, interval, charts.Pattern)
	return status, value
}
