package db

import (
	"database/sql"

	_ "github.com/lib/pq" // This blank import ensures the PostgreSQL driver is registered.
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgressql://root:password@localhost:5433/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) getDB() *sql.DB {
	return d.db
}
