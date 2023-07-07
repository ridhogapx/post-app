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

func TestUpdateUser(t *testing.T) {
	updated, err := q.UpdateUser(context.Background(), UpdateUserParams{
		ID:       3,
		Username: "supaman14",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, updated)
}

func TestDeleteUser(t *testing.T) {
	err := q.DeleteUser(context.Background(), 4)

	assert.Nil(t, err)
}
