package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuoteIdent(t *testing.T) {
	assert.Equal(t, `'db'.'table'.'column'`, QuoteIdent("db.table.column", "'"))
}

func TestQuoteStrings(t *testing.T) {
	assert.Equal(t, `'1','2','3'`, QuoteStrings("1", 2, "3"))
	assert.Equal(t, `'\'1','2','3'`, QuoteStrings("'1", 2, "3"))
}
