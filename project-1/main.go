package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	internal "trading-ai/domain"
)

func main() {

	// Init
	coinList := make([]domain.Ticker, 12)
	charts := make([]domain.Chart, 12)
	interval := 60
	frameRate := 30
	coinName := []string{
		"Abyss",
		"Ten",
		"Dax",
		"Dent",
		"Doge",
		"Gsc",
		"Hart",
		"Mbl",
		"Nxt",
		"Pando",
		"Slp",
		"Xrp",
	}
	coinStatus := make([]string, 12)
	coinValue := make([]int, 12)

	for i := 0; i < len(charts); i++ {
		charts[i] = domain.Chart{
			Pattern: make([]func(candle []domain.Candle) string, 6),
		}
		charts[i].Pattern = []func(candle []domain.Candle) string{
			func(candle []domain.Candle) string {
				length := len(candle)
				var newCandle []domain.Candle
				if length < 3 {
					return ""
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					high := [2]int{newCandle[0].Close, newCandle[1].Open}
					// low := [2]int{newCandle[0].low, newCandle[1].low}
					midPoint := (high[0] - newCandle[0].Open) / 2

					if high[0] < high[1] && newCandle[1].Last < midPoint {
						return "CLOUD COVER"
					}
				} else {
					newCandle = candle
					high := [2]int{newCandle[0].Close, newCandle[1].Open}
					// low := [2]int{newCandle[0].low, newCandle[1].low}
					midPoint := (high[0] - newCandle[0].Open) / 2

					if high[0] < high[1] && newCandle[1].Last < midPoint {
						return "CLOUD COVER"
					}
				}

				return ""
			},
			func(candle []domain.Candle) string {
				length := len(candle)
				var newCandle []domain.Candle
				if length < 3 {
					return ""
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					high := [2]int{newCandle[0].Open, newCandle[1].Last}
					low := [2]int{newCandle[0].Close, newCandle[1].Open}
					midPoint := (high[0] - low[0]) / 2

					if low[0] > low[1] && high[1] > midPoint {
						return "PIERCING"
					}
				} else {
					newCandle = candle
					high := [2]int{newCandle[0].Open, newCandle[1].Last}
					low := [2]int{newCandle[0].Close, newCandle[1].Open}
					midPoint := (high[0] - low[0]) / 2

					if low[0] > low[1] && high[1] > midPoint {
						return "PIERCING"
					}
				}

				return ""
			},
			func(candle []domain.Candle) string {
				length := len(candle)
				var newCandle []domain.Candle
				if length < 3 {
					return ""
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					high := [2]int{newCandle[0].Open, newCandle[1].Last}
					midPoint := (high[0] - newCandle[0].Close) / 2

					if high[0]+midPoint < high[1] {
						return "Bullish Engulfing"
					}
				} else {
					newCandle = candle
					high := [2]int{newCandle[0].Open, newCandle[1].Last}
					midPoint := (high[0] - newCandle[0].Close) / 2

					if high[0]+midPoint < high[1] {
						return "Bullish Engulfing"
					}
				}

				return ""
			},
			func(candle []domain.Candle) string {
				length := len(candle)
				var newCandle []domain.Candle
				if length < 3 {
					return ""
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					low := [2]int{newCandle[0].Open, newCandle[1].Last}
					midPoint := (newCandle[0].High - low[0]) / 2

					if low[0]-midPoint > low[1] {
						return "Bearish Engulfing"
					}
				} else {
					newCandle = candle
					low := [2]int{newCandle[0].Open, newCandle[1].Last}
					midPoint := (newCandle[0].High - low[0]) / 2

					if low[0]-midPoint > low[1] {
						return "Bearish Engulfing"
					}
				}
				return ""
			},
			func(candle []domain.Candle) string {
				length := len(candle)
				var newCandle []domain.Candle
				if length < 3 {
					return ""
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					supportLevel := [2]int{newCandle[0].Close, newCandle[1].Open}
					high := [2]int{newCandle[0].Open, newCandle[1].Last}

					if supportLevel[0] == supportLevel[1] && high[0] < high[1] {
						return "Bearish Reversal"
					}
				} else {
					newCandle = candle
					supportLevel := [2]int{newCandle[0].Close, newCandle[1].Open}
					high := [2]int{newCandle[0].Open, newCandle[1].Last}

					if supportLevel[0] == supportLevel[1] && high[0] < high[1] {
						return "Bearish Reversal"
					}
				}

				return ""
			},
			func(candle []domain.Candle) string {
				length := len(candle)
				var newCandle []domain.Candle
				if length < 3 {
					return ""
				} else if length > 3 {

					for i := length - 3; i < length-1; i++ {
						newCandle = append(newCandle, candle[i])
					}

					supportLevel := [2]int{newCandle[0].Close, newCandle[1].Open}
					low := [2]int{newCandle[0].Open, newCandle[1].Last}

					if supportLevel[0] == supportLevel[1] && low[0] > low[1] {
						return "Bullish Reversal"
					}
				} else {
					newCandle = candle
					supportLevel := [2]int{newCandle[0].Close, newCandle[1].Open}
					high := [2]int{newCandle[0].Open, newCandle[1].Last}

					if supportLevel[0] == supportLevel[1] && high[0] < high[1] {
						return "Bullish Reversal"
					}
				}

				return ""
			},
		}
		charts[i].Candles = append(charts[i].Candles, domain.Candle{})
		charts[i].Name = coinName[i]
	}

	// Request Data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res, err := http.Get("https://indodax.com/api/tickers")

		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			panic(errors.New("Not Connected"))
		}

		body, err := io.ReadAll(res.Body)
		cryptocoins := domain.Cryptocoins{}

		err = json.Unmarshal(body, &cryptocoins)
		if err != nil {
			panic(err)
		}

		coinList[0] = cryptocoins.Coins.Abyss_idr
		coinList[1] = cryptocoins.Coins.Ten_idr
		coinList[2] = cryptocoins.Coins.Dax_idr
		coinList[3] = cryptocoins.Coins.Dent_idr
		coinList[4] = cryptocoins.Coins.Doge_idr
		coinList[5] = cryptocoins.Coins.Gsc_idr
		coinList[6] = cryptocoins.Coins.Hart_idr
		coinList[7] = cryptocoins.Coins.Mbl_idr
		coinList[8] = cryptocoins.Coins.Nxt_idr
		coinList[9] = cryptocoins.Coins.Pando_idr
		coinList[10] = cryptocoins.Coins.Slp_idr
		coinList[11] = cryptocoins.Coins.Xrp_idr
		// Update and Analyse

		for i := 0; i < len(charts); i++ {
			temp1, temp2 := UpdateChart(&charts[i], &cryptocoins, frameRate, interval, coinList[i])
			coinValue[i] = temp2
			coinStatus[i] = temp1
		}

		filepath := path.Join("views", "index.html")
		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string][]interface{}{

			"title":  {"Trading API"},
			"name":   {coinName},
			"value":  {coinValue},
			"status": {coinStatus},
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":8080", nil)
}

func UpdateChart(charts *domain.Chart, cryptocoins *domain.Cryptocoins, frameRate int, interval int, ticker domain.Ticker) (string, int) {
	charts.candles[len(charts.candles)-1].Update(ticker)
	value, status := charts.CreateNewCandle(frameRate, interval, charts.Pattern)
	return status, value
}

func (ch *domain.Chart) CreateNewCandle(frameRate, interval int, callback []func(candle []domain.Candle) string) (int, string) {

	var status string
	var value int

	ch.TimeSeconds += frameRate
	if ch.TimeSeconds > interval {
		for i := 0; i < len(ch.Pattern); i++ {
			message := ch.Pattern[i](ch.Candles)
			if message != "" {
				fmt.Println(ch.Name, "\t", message)
			}
		}
		status = ch.Candles[len(ch.Candles)-1].Status
		value = ch.Candles[len(ch.Candles)-1].Last
		fmt.Println(ch.name, "\t", ch.candles[len(ch.candles)-1].last, ch.candles[len(ch.candles)-1].status)
		ch.Candles[len(ch.Candles)-1].close = ch.candles[len(ch.candles)-1].last
		ch.Candles = append(ch.Candles, internal.Candle{})
		ch.timeSeconds = 0
	} else {
		value = ch.candles[len(ch.candles)-1].last
		status = ch.candles[len(ch.candles)-1].status
	}

	return value, status
}

func (ca *domain.Candle) Update(t domain.Ticker) {
	ca.last, _ = strconv.Atoi(t.Last)
	if ca.open != 0 {
		if ca.last > ca.high {
			ca.high = ca.last
		} else if ca.last < ca.low {
			ca.low = ca.last
		}

		if ca.last >= ca.open {
			ca.status = "Bullish"
		} else {
			ca.status = "Bearish"
		}
	} else {
		ca.low = ca.last
		ca.high = ca.last
		ca.open = ca.last
		if ca.last >= ca.open {
			ca.status = "Bullish"
		} else {
			ca.status = "Bearish"
		}
	}
}
