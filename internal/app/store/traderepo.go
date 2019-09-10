package store

import (
	"local/go-crypto-trader/internal/app/model"
)

// TradeRepo ...
type TradeRepo struct {
	store *Store
}

// SaveTrades ...
func (r *TradeRepo) SaveTrades(t *model.TradesResponse) error {

	return nil
}
