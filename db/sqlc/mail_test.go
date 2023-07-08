package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMail(t *testing.T) {
	newMail, err := q.CreateMail(context.Background(), CreateMailParams{
		Title:    "Greet New Neighbour",
		Body:     "Hi, my name Steven. I live next to your house!",
		Receiver: 2,
		Sender:   3,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, newMail)
}

func TestListMails(t *testing.T) {
	list, err := q.ListMails(context.Background(), 2)

	assert.Nil(t, err)
	assert.NotEmpty(t, list)
}
