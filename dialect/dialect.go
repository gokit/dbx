package dialect

var (
	// MySQL dialect
	MySQL = mysql{}
	// PostgreSQL dialect
	PostgreSQL = postgreSQL{}
	// SQLite3 dialect
	SQLite3 = sqlite3{}
)

const (
	timeFormat = "2006-01-02 15:04:05.000000"
)
