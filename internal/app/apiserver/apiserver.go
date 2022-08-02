package apiserver

import (
	"database/sql"
	"github.com/Dennikoff/UserTagApi/internal/app/store/sqlstore"
	"net/http"
)

func Start(config Config) error {
	db, err := NewDB(config.DriverName, config.DatabaseURL)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()
	store := sqlstore.New(db)
	server := newServer(store)
	return http.ListenAndServe(config.Addr, server)
}

func NewDB(driverName string, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(driverName, databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
