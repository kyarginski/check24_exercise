package testhelpers

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func TestNewTestDatabase(t *testing.T) {
	got, err := NewTestDatabase(t)
	assert.NoError(t, err)
	err = got.DB().Ping()
	assert.NoError(t, err)
}
