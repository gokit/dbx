package dbx

// Select creates a SelectStmt.
func (s *SelectStmt) AddSelect(column ...string) *SelectStmt {
	s.Column = append(s.Column, prepareSelect(column)...)
	return s
}

// Where adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) AndWhere(query interface{}, value ...interface{}) *SelectStmt {
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

// Where adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) OrWhere(query interface{}, value ...interface{}) *SelectStmt {
	switch query := query.(type) {
	case string:
		b.WhereCond = append(b.WhereCond, OrBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, Expr(query, value...))
		}))
	case Builder:
		b.WhereCond = append(b.WhereCond, OrBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, query)
		}))
	}
	return b
}

// Where adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) AndHaving(query interface{}, value ...interface{}) *SelectStmt {
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

// Where adds a where condition.
// query can be Builder or string. value is used only if query type is string.
func (b *SelectStmt) OrHaving(query interface{}, value ...interface{}) *SelectStmt {
	switch query := query.(type) {
	case string:
		b.HavingCond = append(b.HavingCond, OrBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, Expr(query, value...))
		}))
	case Builder:
		b.HavingCond = append(b.HavingCond, OrBuildFunc(func(d Dialect, buf Buffer) error {
			return buildLogicCond(d, buf, query)
		}))
	}
	return b
}