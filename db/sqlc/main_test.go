package db

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var conn *sql.DB
var q *Queries

func init() {
	var err error
	conn, err = sql.Open("postgres", "postgres://root:root@localhost:5432/post_app?sslmode=disable")

	if err != nil {
		panic("Failed to connect Database!")
	}

	q = New(conn)
}

func TestCreateUser(t *testing.T) {
	user, err := q.CreateUser(context.Background(), "john23")

	assert.Nil(t, err)
	assert.NotEmpty(t, user)

}
