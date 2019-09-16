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
type TradesResponse map[string]TradeElement

// TradeElement ...
type TradeElement struct {
	ID       int64  `json:"id" db:"id"`
	TradeID  int64  `json:"trade_id" db:"trade_id"`
	TypeID   int64  `json:"type_id" db:"operation_type_id"`
	Type     string `json:"type" db:"-"`
	Price    string `json:"price" db:"price"`
	Quantity string `json:"quantity" db:"quantity"`
	Amount   string `json:"amount" db:"amount"`
	Date     int    `json:"date" db:"date"`
}

// OperationType ...
type OperationType struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"operation_type_name"`
	Descr string `json:"description" db:"description"`
}
