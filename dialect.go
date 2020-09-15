package dbx

import (
	"github.com/gokit/dbx/schema"
	"time"
)

// Dialect abstracts database driver differences in encoding
// types, and placeholders.
type Dialect interface {
	DriverName() string
	QuoteIdent(id string) string

	EncodeString(s string) string
	EncodeBool(b bool) string
	EncodeTime(t time.Time) string
	EncodeBytes(b []byte) string

	Placeholder(n int) string

	Schema(query schema.Query) schema.Dialect
}
