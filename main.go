package main

import (
	"time"
	"trading-ai/domain"
	"trading-ai/pkg"
)

func main() {

	cryptocoins := domain.CryptoCoins{}
	// updates, bot := pkg.BotInitialization()
	Chart, len, tickers := pkg.AIinitialization(cryptocoins, 60)
	records := [][]string{}

	for {

		pkg.GetJSONFromUrl(&cryptocoins, "https://indodax.com/api/summaries")
		records = append(records, pkg.UnpackedTickerToString(cryptocoins.Coins, 2))
		pkg.WriteCSV(records)
		tickers, _ = cryptocoins.Coins.CreateArrayOfTicker()
		for i := 0; i < len; i++ {
			UpdateChart(&Chart[i], &cryptocoins, 60, tickers[i])
		}

		time.Sleep(30 * time.Second)
	}

	// for update := range updates {
	// 	if update.Message != nil { // If we got a message
	// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 		msg := tgbotapi.NewMessage(1279845132, "s")
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		bot.Send(msg)
	// 	}
	// }

}

func UpdateChart(charts *domain.Chart, cryptocoins *domain.CryptoCoins, frameRate int, ticker domain.Ticker) (string, int) {
	charts.Candles[len(charts.Candles)-1].Update(ticker)
	value, status := charts.CreateNewCandle(frameRate, charts.Pattern)
	return status, value
}
