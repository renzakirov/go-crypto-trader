package store

import (
	"fmt"
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

// SaveTradeElement ...
func (r *TradeRepo) SaveTradeElement(t *model.TradeElement) error {
	result, err := r.store.db.NamedExec(`
		INSERT into trade(trade_id, operation_type_id, price, quantity, amount, date)
		VALUES (:trade_id, :operation_type_id, :price, :quantity, :amount, to_timestamp(:date) AT TIME ZONE 'UTC')
		RETURNING id
	`,
		t,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

// FetchOperationTypes ...
func (r *TradeRepo) FetchOperationTypes() ([]model.OperationType, error) {
	tt := []model.OperationType{}
	err := r.store.db.Select(&tt, "SELECT * FROM operation_type")
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", tt)
	return tt, nil
}

// FetchOperationTypeByID ...
func (r *TradeRepo) FetchOperationTypeByID(m *model.OperationType) (*model.OperationType, error) {
	rows, err := r.store.db.NamedQuery("SELECT * FROM operation_type where id = :id", m)
	if err != nil {
		return nil, err
	}
	t := model.OperationType{}
	for rows.Next() {
		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("%#v", t)
	return &t, nil
}
