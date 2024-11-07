package datastore

import (
	"database/sql"
	"net/url"
	"time"

	_ "github.com/lib/pq"
)

func NewDatabase(databaseURL string) (db *sql.DB, err error) {
	parseDatabaseURL, _ := url.Parse(databaseURL)
	db, err = sql.Open(parseDatabaseURL.Scheme, databaseURL)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)

	return
}
