package dbx

// AndWhere adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *UpdateStmt) AndWhere(query interface{}, value ...interface{}) *UpdateStmt {
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

// OrWhere adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *UpdateStmt) OrWhere(query interface{}, value ...interface{}) *UpdateStmt {
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
func (b *UpdateStmt) ToSql() (string, []interface{}, error) {
	return ToSql(b.Dialect, b)
}

// ToRawSql return the raw sql
func (b *UpdateStmt) ToRawSql() (string, error) {
	return ToRawSql(b.Dialect, b)
}