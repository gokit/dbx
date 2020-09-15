package schema

import (
	"strings"
)

// MySQL schema dialect
func MySQL(query Query) Dialect {
	return mysql{
		query: query,
	}
}

// PostgreSQL schema dialect
func PostgreSQL(query Query) Dialect {
	return postgreSQL{
		query: query,
	}
}

// SQLite3 schema dialect
func SQLite3(query Query) Dialect {
	return sqlite3{
		query: query,
	}
}

// query and load
type Query func(sql string, dest interface{}, args ...interface{}) (int, error)

// schema dialect
type Dialect interface {
	// Create the column definition for a char type.
	TypeChar(column Column) (string, error)
	// Create the column definition for a string type.
	TypeString(column Column) (string, error)
	// Create the column definition for a text type.
	TypeText(column Column) (string, error)
	// Create the column definition for a medium text type.
	TypeMediumText(column Column) (string, error)
	// Create the column definition for a long text type.
	TypeLongText(column Column) (string, error)
	// Create the column definition for a big integer type.
	TypeBigInteger(column Column) (string, error)
	// Create the column definition for an integer type.
	TypeInteger(column Column) (string, error)
	// Create the column definition for a medium integer type.
	TypeMediumInteger(column Column) (string, error)
	// Create the column definition for a tiny integer type.
	TypeTinyInteger(column Column) (string, error)

	// Create the column definition for a tiny blob type.
	TypeTinyBlob(column Column) (string, error)
	// Create the column definition for an blob type.
	TypeBlob(column Column) (string, error)
	// Create the column definition for a medium blob type.
	TypeMediumBlob(column Column) (string, error)
	// Create the column definition for a long blob type.
	TypeLongBlob(column Column) (string, error)

	// Create the column definition for a small integer type.
	TypeSmallInteger(column Column) (string, error)
	// Create the column definition for a float type.
	TypeFloat(column Column) (string, error)
	// Create the column definition for a double type.
	TypeDouble(column Column) (string, error)
	// Create the column definition for a decimal type.
	TypeDecimal(column Column) (string, error)
	// Create the column definition for a boolean type.
	TypeBoolean(column Column) (string, error)
	// Create the column definition for an enumeration type.
	TypeEnum(column Column) (string, error)
	// Create the column definition for a set enumeration type.
	TypeSet(column Column) (string, error)
	// Create the column definition for a json type.
	TypeJson(column Column) (string, error)
	// Create the column definition for a jsonb type.
	TypeJsonb(column Column) (string, error)
	// Create the column definition for a date type.
	TypeDate(column Column) (string, error)
	// Create the column definition for a date-time type.
	TypeDateTime(column Column) (string, error)
	// Create the column definition for a date-time (with time zone) type.
	TypeDateTimeTz(column Column) (string, error)
	// Create the column definition for a time type.
	TypeTime(column Column) (string, error)
	// Create the column definition for a time (with time zone) type.
	TypeTimeTz(column Column) (string, error)
	// Create the column definition for a timestamp type.
	TypeTimestamp(column Column) (string, error)
	// Create the column definition for a timestamp (with time zone) type.
	TypeTimestampTz(column Column) (string, error)
	// Create the column definition for a year type.
	TypeYear(column Column) (string, error)
	// Create the column definition for a binary type.
	TypeBinary(column Column) (string, error)
	// Create the column definition for a uuid type.
	TypeUuid(column Column) (string, error)
	// Create the column definition for an IP address type.
	TypeIpAddress(column Column) (string, error)
	// Create the column definition for a MAC address type.
	TypeMacAddress(column Column) (string, error)
	// Create the column definition for a spatial Geometry type.
	TypeGeometry(column Column) (string, error)
	// Create the column definition for a spatial Point type.
	TypePoint(column Column) (string, error)
	// Create the column definition for a spatial LineString type.
	TypeLineString(column Column) (string, error)
	// Create the column definition for a spatial Polygon type.
	TypePolygon(column Column) (string, error)
	// Create the column definition for a spatial GeometryCollection type.
	TypeGeometryCollection(column Column) (string, error)
	// Create the column definition for a spatial MultiPoint type.
	TypeMultiPoint(column Column) (string, error)
	// Create the column definition for a spatial MultiLineString type.
	TypeMultiLineString(column Column) (string, error)
	// Create the column definition for a spatial MultiPolygon type.
	TypeMultiPolygon(column Column) (string, error)
	// Modify the column
	ModifyColumn(column Column) string
	// Compile the query exists of the table
	CompileTableExists(tableName string, tableSchema ...string) (string, error)
	// Compile the query to determine the list of columns.
	CompileColumnListing(tableName string, tableSchema ...string) (string, error)
	// Compile a create table command.
	CompileCreate(table Table) (string, error)
	//  Compile a modify columns command.
	CompileModifyColumns(table Table) (string, error)
	//  Compile a modify column command.
	CompileModifyColumn(table Table, columnName string) (string, error)
	// Compile add columns.
	CompileAddColumns(table Table) (string, error)
	// Compile add columns.
	CompileAddColumn(table Table, columnName string) (string, error)
	// Compile a primary key command.
	CompilePrimaryKey(table Table, columnNames ...string) (string, error)
	// Compile a drop primary key command.
	CompileDropPrimaryKey(table Table) (string, error)
	// Compile an index creation command.
	CompileIndex(table Table, index Index) (string, error)
	// Compile a drop index command.
	CompileDropIndex(table Table, indexName string) (string, error)
	// Compile a drop unique index command.
	CompileDropUnique(table Table, indexName string) (string, error)
	// Compile a drop spatial index command.
	CompileDropSpatialIndex(table Table, indexName string) (string, error)
	// Compile a drop foreign index command.
	CompileDropForeign(table Table, indexName string) (string, error)
	// Compile a drop table command.
	CompileDrop(tableName string) (string, error)
	// Compile a drop table (if exists) command.
	CompileDropIfExists(tableName string) (string, error)
	// Compile a drop column command.
	CompileDropColumn(table Table, columnNames []string) (string, error)
	// Compile a rename table command.
	CompileRenameTable(table Table, toName string) (string, error)
	// Compile a rename index command.
	CompileRenameIndex(table Table, from string, to string) (string, error)
	// Compile the SQL needed to drop all tables. [PostgreSQL]
	CompileDropAllTables(tableNames ...string) (string, error)
	// Compile the SQL needed to drop all views.
	CompileDropAllViews(viewNames ...string) (string, error)
	// Compile the SQL needed to drop all types.
	CompileDropAllTypes(schemaNames ...string) (string, error)
	// Compile the SQL needed to retrieve all table names.
	CompileGetAllTables(schemaNames ...string) (string, error)
	// Compile the SQL needed to retrieve all view names.
	CompileGetAllViews(schemaNames ...string) (string, error)
	// Compile the SQL needed to retrieve all type names. [PostgreSQL]
	CompileGetAllTypes() (string, error)
	// Compile the SQL needed to rebuild the database. [SQLite]
	CompileRebuild() (string, error)
	// Compile the command to enable foreign key constraints.
	CompileEnableForeignKeyConstraints() (string, error)
	// Compile the command to disable foreign key constraints.
	CompileDisableForeignKeyConstraints() (string, error)
	// Load table columns from the database
	LoadColumns(tableName string, tableSchema ...string) ([]Column, error)
	// Load table from the database
	LoadTable(tableName string, tableSchema ...string) (Table, error)
	// Load table constraints from the database
	LoadTableConstraints(tableName string, tableSchema ...string) (TableConstraints, error)
}

// wrap table name if prefix exists
func wrapTableName(table Table, quoteIdent func(s string) string) string {
	var b strings.Builder

	if table.Schema() != "" {
		b.WriteString(quoteIdent(table.Schema()))
		b.WriteString(".")
	}

	if table.Prefix() == "" {
		b.WriteString(quoteIdent(table.Name()))
	} else {
		b.WriteString(quoteIdent(table.Prefix() + table.Name()))
	}

	return b.String()
}
