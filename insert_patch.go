package dbx

// ToSql return the sql and args
func (b *InsertStmt) ToSql() (string, []interface{}, error) {
	return ToSql(b.Dialect, b)
}

// ToRawSql return the raw sql
func (b *InsertStmt) ToRawSql() (string, error) {
	return ToRawSql(b.Dialect, b)
}
