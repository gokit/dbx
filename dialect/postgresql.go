package dialect

import (
	"fmt"
	"github.com/gokit/dbx/schema"
	"github.com/gokit/dbx/utils"
	"strings"
	"time"
)

type postgreSQL struct{}

func (d postgreSQL) DriverName() string {
	return "postgres"
}

func (d postgreSQL) QuoteIdent(s string) string {
	return utils.QuoteIdent(s, `"`)
}

func (d postgreSQL) EncodeString(s string) string {
	// http://www.postgresql.org/docs/9.2/static/sql-syntax-lexical.html
	return `'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

func (d postgreSQL) EncodeBool(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

func (d postgreSQL) EncodeTime(t time.Time) string {
	return MySQL.EncodeTime(t)
}

func (d postgreSQL) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`E'\\x%x'`, b)
}

func (d postgreSQL) Placeholder(n int) string {
	return fmt.Sprintf("$%d", n+1)
}

func (d postgreSQL) Schema(query schema.Query) schema.Dialect {
	return schema.PostgreSQL(query)
}
