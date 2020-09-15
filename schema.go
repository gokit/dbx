package dbx

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/gokit/dbx/schema"
)

func (sess *Session) queryAndLoad(sql string, dest interface{}, args ...interface{}) (int, error) {
	rows, err := sess.Query(sql, args...)

	if err != nil {
		return 0, err
	}

	return Load(rows, dest)
}

// Query exists of the table
func (sess *Session) TableExists(tableName string, tableSchema ...string) (bool, error) {
	sql, err := sess.Schema(sess.queryAndLoad).CompileTableExists(tableName, tableSchema...)

	if err != nil {
		return false, err
	}

	var count int

	_, err = sess.queryAndLoad(sql, &count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Create table
func (sess *Session) CreateTable(tableName string, handler schema.TableHandler) error {

	table := schema.NewTable(tableName)

	// add create table command
	table.AddCommand(func(d schema.Dialect) (string, error) {
		return d.CompileCreate(table)
	})

	handler(table)

	schema := sess.Schema(sess.queryAndLoad)

	for _, command := range table.Commands() {
		sql, err := command(schema)

		fmt.Println(sql)

		if err != nil {
			return err
		}

		_, err = sess.Exec(sql)

		if err != nil {
			return err
		}
	}

	return nil
}

// Modify the table
func (sess *Session) Table(tableName string, handler schema.TableHandler) error {
	t, err := sess.GetTable(tableName)

	if err != nil {
		return errors.Wrapf(err, "dbx: error on find table '%s'", tableName)
	}

	if t == nil {
		return errors.New(fmt.Sprintf("dbx: not found the table '%s'", tableName))
	}

	var table = t.(*schema.TableSchema)

	handler(table)

	schema := sess.Schema(sess.queryAndLoad)

	for _, command := range table.Commands() {
		sql, err := command(schema)

		if err != nil {
			return err
		}

		fmt.Println(sql)

		_, err = sess.Exec(sql)

		if err != nil {
			return err
		}
	}

	return nil
}

// Drop the table
func (sess *Session) Drop(tableName string) (sql.Result, error) {
	sql, err := sess.Schema(sess.queryAndLoad).CompileDrop(tableName)

	if err != nil {
		return nil, err
	}

	res, err := sess.Exec(sql)

	return res, err
}

// Drop the table if exists
func (sess *Session) DropIfExists(tableName string) (sql.Result, error) {
	sql, err := sess.Schema(sess.queryAndLoad).CompileDropIfExists(tableName)

	if err != nil {
		return nil, err
	}

	res, err := sess.Exec(sql)

	return res, err
}

// Get table from the database
func (sess *Session) GetTable(tableName string, tableSchema ...string) (schema.Table, error) {
	return sess.Schema(sess.queryAndLoad).LoadTable(tableName, tableSchema...)
}

// Get table columns from the database
func (sess *Session) GetColumns(tableName string, tableSchema ...string) ([]schema.Column, error) {
	return sess.Schema(sess.queryAndLoad).LoadColumns(tableName, tableSchema...)
}

// Get table constraints from the database
func (sess *Session) GetTableConstraints(tableName string, tableSchema ...string) (schema.TableConstraints, error) {
	return sess.Schema(sess.queryAndLoad).LoadTableConstraints(tableName, tableSchema...)
}
