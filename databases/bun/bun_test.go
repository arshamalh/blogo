package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	dsn := "our_actual_database_dsn" // ?

	db, err := Connect(dsn)
	assert.NoError(t, err)
	assert.NotNil(t, db)
	assert.NotNil(t, db.db)
}
