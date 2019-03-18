package models

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

// Connecting with db
func NewDB(dataSourceName string) (*sql.DB, error) {
	// open connection
	db, err := sql.Open("mysql", dataSourceName)

	// catch exceptions
	if err != nil {
		return nil, err
	}

	// if db doesn't connect
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// return result or null
	return db, nil
}
