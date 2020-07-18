package dbx

// AddSelect specifies columns for select.
func (s *SelectStmt) AddSelect(column ...string) *SelectStmt {
	s.Column = append(s.Column, prepareSelect(column)...)
	return s
}

// AndWhere adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) AndWhere(query interface{}, value ...interface{}) *SelectStmt {
	switch query := query.(type) {
	case string:
		b.WhereCond = append(b.WhereCond, Expr(query, value...))
	case Builder:
		b.WhereCond = append(b.WhereCond, query)
	}
	return b
}

// OrWhere adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) OrWhere(query interface{}, value ...interface{}) *SelectStmt {
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

// AndHaving adds a having condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) AndHaving(query interface{}, value ...interface{}) *SelectStmt {
	switch query := query.(type) {
	case string:
		b.WhereCond = append(b.WhereCond, Expr(query, value...))
	case Builder:
		b.WhereCond = append(b.WhereCond, query)
	}
	return b
}

// OrHaving adds a having condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) OrHaving(query interface{}, value ...interface{}) *SelectStmt {
	switch query := query.(type) {
	case string:
		b.HavingCond = append(b.HavingCond, orBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, Expr(query, value...))
		}))
	case Builder:
		b.HavingCond = append(b.HavingCond, orBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, query)
		}))
	}
	return b
}

// GroupBy specifies columns for grouping.
func (b *SelectStmt) AddGroupBy(col ...string) *SelectStmt {
	for _, group := range col {
		b.Group = append(b.Group, Expr(group))
	}
	return b
}

// OrderBy specifies columns for ordering.
func (b *SelectStmt) AddOrderBy(cols ...string) *SelectStmt {
	for _, col := range cols {
		b.Order = append(b.Order, Expr(col))
	}
	return b
}

// ToSql return the sql and args
func (b *SelectStmt) ToSql() (string, []interface{}, error) {
	return ToSql(b.Dialect, b)
}

// ToRawSql return the raw sql
func (b *SelectStmt) ToRawSql() (string, error) {
	return ToRawSql(b.Dialect, b)
}
