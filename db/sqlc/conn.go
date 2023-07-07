package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func TestMain(m *testing.M) {
	var err error
	conn, err = sql.Open("postgres", "postgres://root:root@localhost:5432/post_app?sslmode=disable")

	if err != nil {
		panic("Failed to connect Database!")
	}

	New(conn)
	os.Exit(m.Run())
}
