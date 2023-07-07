package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUser(t *testing.T) {
	user, err := q.ListUsers(context.Background(), ListUsersParams{
		Limit:  1,
		Offset: 2,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}
