package domain

import "strconv"

type Candle struct {
	Status string
	High   int
	Low    int
	Open   int
	Close  int
	Last   int
}

func (ca *Candle) Update(t Ticker) {
	ca.Last, _ = strconv.Atoi(t.Last)
	if ca.Open != 0 {
		if ca.Last > ca.High {
			ca.High = ca.Last
		} else if ca.Last < ca.Low {
			ca.Low = ca.Last
		}

		if ca.Last >= ca.Open {
			ca.Status = "Bullish"
		} else {
			ca.Status = "Bearish"
		}
	} else {
		ca.Low = ca.Last
		ca.High = ca.Last
		ca.Open = ca.Last
		if ca.Last >= ca.Open {
			ca.Status = "Bullish"
		} else {
			ca.Status = "Bearish"
		}
	}
}
