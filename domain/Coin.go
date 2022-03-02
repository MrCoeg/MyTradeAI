package domain

type Coin struct {
	Btc  BtcTicker  `json:"btc_idr"`
	Aave AaveTicker `json:"aave_idr"`
	Em   EmTicker   `json:"em_idr"`
}

type CryptoCoins struct {
	Coins Coin `json:"tickers"`
}

func (c *Coin) CreateArrayOfTicker() (t []Ticker, arrLen int) {
	t = []Ticker{
		c.Btc.ChangeCatcherToTicker(),
		c.Aave.ChangeCatcherToTicker(),
	}

	return t, len(t)
}
