package domain

type Ticker struct {
	High        string `json:"high"`
	Low         string `json:"low"`
	Last        string `json:"last"`
	Buy         string `json:"buy"`
	Sell        string `json:"sell"`
	Server_time int    `json:"server_time"`
}
