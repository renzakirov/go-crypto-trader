package model

// TickerResponse ...
type TickerResponse map[string]TickerElement

// TickerElement ...
type TickerElement struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int    `json:"updated"`
}

// OrderBookResponse ...
type OrderBookResponse struct {
	AskQuantity string  `json:"ask_quantity"`
	AskAmount   string  `json:"ask_amount"`
	AskTop      string  `json:"ask_top"`
	BidQuantity string  `json:"bid_quantity"`
	BidAmount   string  `json:"bid_amount"`
	BidTop      string  `json:"bid_top"`
	Ask         [][]int `json:"ask"`
	Bid         [][]int `json:"bid"`
}

// TradesResponse ...
type TradesResponse struct {
	BTCUSD []TradeElements
}

// TradeElements ...
type TradeElements struct {
	TradeID  int    `json:"trade_id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
	Date     int    `json:"date"`
}
