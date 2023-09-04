package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	//ping db to check conn
	err = db.Ping()
	assert.NoError(t, err)
}
