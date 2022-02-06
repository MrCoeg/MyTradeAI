package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"time"
	"trading-ai/domain"
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

	go func() {
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
			cryptocoins := domain.CryptoCoins{}

			err = json.Unmarshal(body, &cryptocoins)
			if err != nil {
				panic(err)
			}

			coinList[0] = cryptocoins.Coins.AbyssIdr
			coinList[1] = cryptocoins.Coins.TenIdr
			coinList[2] = cryptocoins.Coins.DaxIdr
			coinList[3] = cryptocoins.Coins.DentIdr
			coinList[4] = cryptocoins.Coins.DogeIdr
			coinList[5] = cryptocoins.Coins.GscIdr
			coinList[6] = cryptocoins.Coins.HartIdr
			coinList[7] = cryptocoins.Coins.MblIdr
			coinList[8] = cryptocoins.Coins.NxtIdr
			coinList[9] = cryptocoins.Coins.PandoIdr
			coinList[10] = cryptocoins.Coins.SlpIdr
			coinList[11] = cryptocoins.Coins.XrpIdr
			// Update and Analyse

			for i := 0; i < len(charts); i++ {
				temp1, temp2 := UpdateChart(&charts[i], &cryptocoins, frameRate, interval, coinList[i])
				coinValue[i] = temp2
				coinStatus[i] = temp1
			}

			time.Sleep(30 * time.Second)
		}
	}()

	// Request Data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join("views", "index.gohtml")
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

func UpdateChart(charts *domain.Chart, cryptocoins *domain.CryptoCoins, frameRate int, interval int, ticker domain.Ticker) (string, int) {
	charts.Candles[len(charts.Candles)-1].Update(ticker)
	value, status := charts.CreateNewCandle(frameRate, interval, charts.Pattern)
	return status, value
}
