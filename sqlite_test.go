package ent_sqlite

import (
	"context"
	"github.com/dstgo/ent-sqlite/testdata/ent"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlite(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	err = client.Schema.Create(context.Background())
	assert.Nil(t, err)
}
