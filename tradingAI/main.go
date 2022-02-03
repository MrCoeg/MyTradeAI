package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Candle struct {
	high  int
	low   int
	open  int
	close int
	last  int
}

type Chart struct {
	candles     []Candle
	timeSeconds int
	Pattern     []func(candle []Candle) string
}

type Cryptocoins struct {
	CoinNames []string
	Coins     Coin `json:"tickers"`
}
type Coin struct {
	Btc_idr   Ticker `json:"btc_idr"`
	Abyss_idr Ticker `json:"abyss_idr"`
}
type Ticker struct {
	High        string `json:"high"`
	Low         string `json:"low"`
	Last        string `json:"last"`
	Buy         string `json:"buy"`
	Sell        string `json:"sell"`
	Server_time int    `json:"server_time"`
}

func main() {

	// Init
	CoinNames := []string{
		"btc_idr",
		"Abyss_idr",
	}
	charts := make([]Chart, 2)
	for i := 0; i < len(charts); i++ {
		charts[i] = Chart{
			Pattern: make([]func(candle []Candle) string, 6),
		}
		charts[i].Pattern = []func(candle []Candle) string{
			func(candle []Candle) string {
				length := len(candle)
				var newCandle []Candle
				if length < 3 {
					return "candle is not enough to analyse"
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					high := [2]int{newCandle[0].last, newCandle[1].last}
					// low := [2]int{newCandle[0].low, newCandle[1].low}

					if high[0] > high[1] {
						return "CLOUD COVER - STOP BUY - SELL NOW"
					}
				} else {
					newCandle = candle
					high := [2]int{newCandle[0].last, newCandle[1].last}
					// low := [2]int{newCandle[0].low, newCandle[1].low}

					if high[0] > high[1] {
						return "CLOUD COVER - STOP BUY - SELL NOW"
					}
				}

				return "seluw"
			},
			func(candle []Candle) string {
				return ""
			},
			func(candle []Candle) string {
				return ""
			},
			func(candle []Candle) string {
				return ""
			},
			func(candle []Candle) string {
				return ""
			},
			func(candle []Candle) string {
				return ""
			},
		}
		charts[i].candles = append(charts[i].candles, Candle{})
	}
	interval := 10
	frameRate := 5

	// Request Data
	for {
		res, err := http.Get("https://indodax.com/api/tickers")

		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			panic(errors.New("Not Connected"))
		}

		body, err := io.ReadAll(res.Body)
		cryptocoins := Cryptocoins{
			CoinNames: make([]string, 2),
		}

		err = json.Unmarshal(body, &cryptocoins)
		if err != nil {
			panic(err)
		}

		// Update and Analyse
		for i := 0; i < len(charts); i++ {
			UpdateChart(&charts[i], &cryptocoins, frameRate, interval, CoinNames[i])
		}
		fmt.Println()
		time.Sleep(time.Duration(frameRate) * time.Second)
	}

}

func UpdateChart(charts *Chart, cryptocoins *Cryptocoins, frameRate int, interval int, coinName string) {
	switch coinName {
	case "btc_idr":
		charts.candles[len(charts.candles)-1].Update(cryptocoins.Coins.Btc_idr)
		fmt.Println(len(charts.candles)-1, charts.candles[len(charts.candles)-1].last)
		charts.CreateNewCandle(frameRate, interval, charts.Pattern)
	case "Abyss_idr":
		charts.candles[len(charts.candles)-1].Update(cryptocoins.Coins.Abyss_idr)
		fmt.Println(len(charts.candles)-1, charts.candles[len(charts.candles)-1].last)
		charts.CreateNewCandle(frameRate, interval, charts.Pattern)
	}
}

func (ch *Chart) CreateNewCandle(frameRate, interval int, callback []func(candle []Candle) string) {

	ch.timeSeconds += frameRate
	if ch.timeSeconds >= interval {
		ch.candles[len(ch.candles)-1].close = ch.candles[len(ch.candles)-1].last
		ch.candles = append(ch.candles, Candle{})

		for i := 0; i < len(ch.Pattern); i++ {
			message := callback[i](ch.candles)
			if message != "" {
				fmt.Println(message)
			}

		}

		ch.timeSeconds = 0
	}
}

func (ca *Candle) Update(t Ticker) {
	ca.last, _ = strconv.Atoi(t.Last)
	if ca.open != 0 {
		if ca.last > ca.high {
			ca.high = ca.last
		} else if ca.last < ca.low {
			ca.low = ca.last
		}
	} else {
		ca.low = ca.last
		ca.high = ca.last
		ca.open = ca.last
	}
}
