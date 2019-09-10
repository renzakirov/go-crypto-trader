package store

import (
	"context"

	"github.com/jackc/pgx"
)

// Store ...
type Store struct {
	config *Config
	db     *pgx.Conn
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	config, err := pgx.ParseConnectionString(s.config.DatabaseURL)
	if err != nil {
		return err
	}
	db, err := pgx.Connect(config)
	if err != nil {
		return err
	}
	err = db.Ping(context.Background())
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
