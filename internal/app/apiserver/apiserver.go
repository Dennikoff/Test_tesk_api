package apiserver

import "database/sql"

func Start(config Config) error {
	db, err := NewDB(config.DriverName, config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
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
