package domain

import "fmt"

type Chart struct {
	Name        string
	Candles     []Candle
	TimeSeconds int
	TimeFrame   int
	VolMean     float64
	Pattern     []func(candle []Candle) string
}

func (ch *Chart) CreateNewCandle(frameRate int, callback []func(candle []Candle) string) (int, string) {
	var status string
	var value int

	ch.TimeSeconds += frameRate
	if ch.TimeSeconds > ch.TimeFrame {
		for i := 0; i < len(ch.Pattern); i++ {
			message := ch.Pattern[i](ch.Candles)
			if message != "" {
				fmt.Println(ch.Name, "\t", message)
			}
		}
		status = ch.Candles[len(ch.Candles)-1].Status
		value = ch.Candles[len(ch.Candles)-1].Last
		ch.Candles[len(ch.Candles)-1].Close = ch.Candles[len(ch.Candles)-1].Last
		ch.Candles = append(ch.Candles, Candle{})
		ch.TimeSeconds = 0
	} else {
		value = ch.Candles[len(ch.Candles)-1].Last
		status = ch.Candles[len(ch.Candles)-1].Status
	}

	return value, status
}

func CreateNewChart(timeFrame int, coinName string) Chart {
	ch := Chart{
		TimeFrame: timeFrame,
		Candles: []Candle{
			Candle{},
		},
		Name: coinName,
	}
	return ch
}
