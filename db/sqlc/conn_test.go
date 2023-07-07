package db

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func TestConnection(t *testing.T) {
	var err error
	conn, err = sql.Open("postgres", "postgres://root:root@localhost:5432/post_app?sslmode=disable")

	if err != nil {
		panic("Failed to connect Database!")
	}

	New(conn)
}
