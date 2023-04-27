package models

type USD struct {
	USDBRL struct {
		Bid  string `json:"bid"`
		Name string `json:"name"`
		Date string `json:"create_date"`
	} `json:"USDBRL"`
}

type BitCoin struct {
	BTCBRL struct {
		Bid  string `json:"bid"`
		Name string `json:"name"`
		Date string `json:"create_date"`
	} `json:"BTCBRL"`
}
