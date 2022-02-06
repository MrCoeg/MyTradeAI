package domain

type Coin struct {
	AbyssIdr Ticker `json:"abyss_idr"`
	TenIdr   Ticker `json:"ten_idr"`
	DaxIdr   Ticker `json:"dax_idr"`
	DentIdr  Ticker `json:"dent_idr"`
	DogeIdr  Ticker `json:"doge_idr"`
	GscIdr   Ticker `json:"gsc_idr"`
	HartIdr  Ticker `json:"hart_idr"`
	MblIdr   Ticker `json:"mbl_idr"`
	NxtIdr   Ticker `json:"nxt_idr"`
	PandoIdr Ticker `json:"pando_idr"`
	SlpIdr   Ticker `json:"slp_idr"`
	XrpIdr   Ticker `json:"xrp_idr"`
}

type CryptoCoins struct {
	Coins Coin `json:"tickers"`
}

func (c *Coin) CreateArrayOfTicker() (t []Ticker, arrLen int) {
	t = []Ticker{
		c.AbyssIdr,
		c.TenIdr,
		c.DaxIdr,
		c.DentIdr,
		c.DogeIdr,
		c.GscIdr,
		c.HartIdr,
		c.MblIdr,
		c.NxtIdr,
		c.NxtIdr,
		c.PandoIdr,
		c.SlpIdr,
		c.XrpIdr,
	}

	return t, len(t)
}
