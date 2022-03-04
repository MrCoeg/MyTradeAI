package pkg

import (
	"log"
	"trading-ai/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotInitialization() (tgbotapi.UpdatesChannel, *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI("TOKEN HERE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return updates, bot
}

func AIinitialization(cryptocoins domain.CryptoCoins, timeFrame int) ([]domain.Chart, int, []domain.Ticker) {

	t, len := cryptocoins.Coins.CreateArrayOfTicker()
	ch := make([]domain.Chart, len)
	for i, val := range t {
		ch[i] = domain.CreateNewChart(timeFrame, val.Name)
	}

	return ch, len, []domain.Ticker{}
}
