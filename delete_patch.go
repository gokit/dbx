package dbx

// AndWhere adds a and where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *DeleteStmt) AndWhere(query interface{}, value ...interface{}) *DeleteStmt {
	switch query := query.(type) {
	case string:
		b.WhereCond = append(b.WhereCond, BuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, Expr(query, value...))
		}))
	case Builder:
		b.WhereCond = append(b.WhereCond, BuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, query)
		}))
	}
	return b
}

// OrWhere adds a or where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *DeleteStmt) OrWhere(query interface{}, value ...interface{}) *DeleteStmt {
	switch query := query.(type) {
	case string:
		b.WhereCond = append(b.WhereCond, orBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, Expr(query, value...))
		}))
	case Builder:
		b.WhereCond = append(b.WhereCond, orBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, query)
		}))
	}
	return b
}

// ToSql return the sql and args
func (b *DeleteStmt) ToSql() (string, []interface{}, error) {
	return ToSql(b.Dialect, b)
}

// ToRawSql return the raw sql
func (b *DeleteStmt) ToRawSql() (string, error) {
	return ToRawSql(b.Dialect, b)
}