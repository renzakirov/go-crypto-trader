package store

import (
	_ "github.com/jackc/pgx/stdlib" // pgx
	"github.com/jmoiron/sqlx"
)

// Store ...
type Store struct {
	config *Config
	db     *sqlx.DB

	TradeRepo *TradeRepo
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sqlx.Connect("pgx", "postgres://root:docker@localhost:5432/gocryptotrader?sslmode=disable") // s.config.DatabaseURL)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Trade() *TradeRepo {
	if s.TradeRepo != nil {
		return s.TradeRepo
	}
	s.TradeRepo = &TradeRepo{
		store: s,
	}
	return s.TradeRepo
}
