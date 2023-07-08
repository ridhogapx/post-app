package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMail(t *testing.T) {
	newMail, err := q.CreateMail(context.Background(), CreateMailParams{
		Title:    "Farewell",
		Body:     "Goodbye, mate!",
		Receiver: 3,
		Sender:   2,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, newMail)
}

func TestListMails(t *testing.T) {
	list, err := q.ListMails(context.Background(), 2)

	assert.Nil(t, err)
	assert.NotEmpty(t, list)
}
