package schema

import (
	"database/sql"
	"github.com/gokit/dbx/schema/constraint"
)

// table constraint info from database
type TableConstraintInfo struct {
	Name               string         `db:"name"`
	ColumnName         string         `db:"column_name"`
	Type               string         `db:"type"`
	ForeignTableSchema sql.NullString `db:"foreign_table_schema"`
	ForeignTableName   sql.NullString `db:"foreign_table_name"`
	ForeignColumnName  sql.NullString `db:"foreign_column_name"`
	OnUpdate           sql.NullString `db:"on_update"`
	OnDelete           sql.NullString `db:"on_delete"`
	CheckExpr          sql.NullString `db:"check_expr"`
}

type TableConstraintInfos []TableConstraintInfo

// Group by type and name, return map[type]map[name]TableConstraintInfos
func (t TableConstraintInfos) Group() map[string]map[string]TableConstraintInfos {
	var group = map[string]map[string]TableConstraintInfos{}

	for _, item := range t {

		if _, ok := group[item.Type]; ok == false {
			group[item.Type] = map[string]TableConstraintInfos{}
		}

		group[item.Type][item.Name] = append(group[item.Type][item.Name], item)
	}

	return group
}

// Get the column names of the table constraint infos
func (t TableConstraintInfos) ColumnNames() []string {
	var columnNames []string
	var set map[string]struct{} = map[string]struct{}{}
	for _, item := range t {
		if _, ok := set[item.Name]; ok {
			continue
		}
		columnNames = append(columnNames, item.ColumnName)
	}
	return columnNames
}

// Get the foreign column names of the table constraint infos
func (t TableConstraintInfos) ForeignColumnNames() []string {
	var columnNames []string
	var set = map[string]struct{}{}
	set[""] = struct{}{}

	for _, item := range t {
		if _, ok := set[item.ForeignTableSchema.String]; ok {
			continue
		}
		columnNames = append(columnNames, item.ForeignTableSchema.String)
	}

	return columnNames
}

type TableConstraints struct {
	PrimaryKey     constraint.PrimaryKey
	ForeignKeys    []constraint.ForeignKey
	Checks         []constraint.Check
	Uniques        []constraint.Unique
}

func (t *TableConstraints) SetPrimaryKey(key constraint.PrimaryKey) {
	t.PrimaryKey = key
}

func (t *TableConstraints) AddForeignKeys(keys ...constraint.ForeignKey) {
	t.ForeignKeys = append(t.ForeignKeys, keys...)
}

func (t *TableConstraints) AddChecks(checks ...constraint.Check) {
	t.Checks = append(t.Checks, checks...)
}

func (t *TableConstraints) AddUniques(keys ...constraint.Unique) {
	t.Uniques = append(t.Uniques, keys...)
}

func (t *TableConstraints) HasPrimaryKey() bool {
	return len(t.PrimaryKey.ColumnNames) > 0
}

func (t *TableConstraints) HasForeignKeys() bool {
	return len(t.ForeignKeys) > 0
}

func (t *TableConstraints) HasChecks() bool {
	return len(t.Checks) > 0
}

func (t *TableConstraints) HashUniques() bool {
	return len(t.Uniques) > 0
}
