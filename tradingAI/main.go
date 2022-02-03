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

type test struct {
}

type Test interface {
}

func main() {
	chart := Chart{
		Pattern: make([]func(candle []Candle) string, 6),
	}

	var chartPattern = []func(candle []Candle) string{
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

	chart.Pattern = chartPattern

	chart.candles = append(chart.candles, Candle{})
	interval := 10
	frameRate := 5
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

		// for i := 0; i < len(cryptocoins.CoinNames); i++ {
		// 	coinName := cryptocoins.CoinNames[i]

		// }

		chart.candles[len(chart.candles)-1].Update(cryptocoins.Coins.Btc_idr)
		fmt.Println(len(chart.candles)-1, chart.candles[len(chart.candles)-1].last)
		chart.CreateNewCandle(frameRate, interval, chartPattern)
		time.Sleep(time.Duration(frameRate) * time.Second)
	}

}

func (ch *Chart) CreateNewCandle(frameRate, interval int, callback []func(candle []Candle) string) {
	// var cp = ChartPattern{}
	ch.timeSeconds += frameRate
	if ch.timeSeconds >= interval {
		ch.candles[len(ch.candles)-1].close = ch.candles[len(ch.candles)-1].last
		ch.candles = append(ch.candles, Candle{})

		for i := 0; i < len(ch.Pattern); i++ {
			message := callback[i](ch.candles)
			fmt.Println(message)
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
