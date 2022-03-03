package pkg

import (
	"strconv"
	"trading-ai/domain"
)

func unpackedData(c domain.Coin) (unpacked [][]string, packed []domain.Ticker) {
	t, len := c.CreateArrayOfTicker()
	var data = make([][]string, len)
	for i := 0; i < len; i++ {
		data[i] = make([]string, 5)
		data[i][0] = t[i].High
		data[i][1] = t[i].Low
		data[i][2] = t[i].Last
		data[i][3] = t[i].Sell
		data[i][4] = strconv.FormatInt(int64(t[i].Server_time), 10)
	}

	return data, t
}

func UnpackedTickerToInt(c domain.Coin, tickerData int) []int {
	unpacked, _ := unpackedData(c)
	var data = make([]int, len(unpacked))
	for i, val := range unpacked {
		data[i], _ = strconv.Atoi(val[tickerData])
	}
	return data
}

func UnpackedTickerToString(c domain.Coin, tickerData int) []string {
	unpacked, _ := unpackedData(c)
	var data = make([]string, len(unpacked))
	for i, val := range unpacked {
		data[i] = val[tickerData]
	}
	return data
}
