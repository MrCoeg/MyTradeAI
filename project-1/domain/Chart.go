package internal

type Chart struct {
	name        string
	candles     []Candle
	timeSeconds int
	Pattern     []func(candle []Candle) string
}
