package schema

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/gokit/dbx/schema/constraint"
	"github.com/gokit/dbx/utils"
	"strconv"
	"strings"
)

type postgresColumnInfo struct {
	Name        string         `db:"column_name"`
	DataType    string         `db:"data_type"`
	Length      sql.NullInt64  `db:"size"`
	Precision   sql.NullInt64  `db:"numeric_precision"`
	Scale       sql.NullInt64  `db:"numeric_scale"`
	Nullable    bool           `db:"is_nullable"`
	Default     sql.NullString `db:"column_default"`
	PrimaryKey  bool           `db:"is_pkey"`
	AutoInc     bool           `db:"is_autoinc"`
	Check       sql.NullString `db:"check"`
	AllowValues sql.NullString `db:"enum_values"`
	Comment     sql.NullString `db:"column_comment"`
}

type postgresTableInfo struct {
	Name    string         `db:"table_name"`
	Schema  string         `db:"table_schema"`
	Comment sql.NullString `db:"table_comment"`
}

type postgreSQL struct {
	query Query
}

func (d postgreSQL) QuoteIdent(s string) string {
	return utils.QuoteIdent(s, `"`)
}

// wrap index name
func (d postgreSQL) wrapIndexName(table Table, indexName string) string {
	var b strings.Builder

	if table.Schema() != "" {
		b.WriteString(d.QuoteIdent(table.Schema()))
		b.WriteString(".")
	}

	b.WriteString(d.QuoteIdent(indexName))

	return b.String()
}

// Create the column definition for a generatable column.
func (d postgreSQL) generatableColumn(columnType string, column Column) string {
	if !column.AutoIncrement() && column.GeneratedAs() == "" {
		return columnType
	}

	if column.Changed() {
		switch columnType {
		case "integer":
			return "int4"
		case "bigint":
			return "int8"
		case "smallint":
			return "int2"
		}
	} else {
		if column.AutoIncrement() && column.GeneratedAs() == "" {
			switch columnType {
			case "integer":
				return "serial"
			case "bigint":
				return "bigserial"
			case "smallint":
				return "smallserial"
			}
		}
	}

	var b strings.Builder

	b.WriteString(columnType)
	b.WriteString(" generated ")

	if column.Alawys() {
		b.WriteString("always")
	} else {
		b.WriteString("by default")
	}

	b.WriteString(" as identity")

	if column.GeneratedAs() != "" {
		b.WriteString("(")
		b.WriteString(column.GeneratedAs())
		b.WriteString(")")
	}

	return b.String()
}

func (d postgreSQL) formatPostGisType(columnType string, column Column) (string, error) {
	if column.Projection() != "" {
		return fmt.Sprintf("geometry(%s, %s)", columnType, column.Projection()), nil
	}
	return fmt.Sprintf("geometry(%s)", columnType), nil
}

// Create the column definition for a char type.
func (d postgreSQL) TypeChar(column Column) (string, error) {
	return fmt.Sprintf("char(%d)", column.Length()), nil
}

// Create the column definition for a string type.
func (d postgreSQL) TypeString(column Column) (string, error) {
	return fmt.Sprintf("varchar(%d)", column.Length()), nil
}

// Create the column definition for a text type.
func (d postgreSQL) TypeText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a medium text type.
func (d postgreSQL) TypeMediumText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a long text type.
func (d postgreSQL) TypeLongText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for an integer type.
func (d postgreSQL) TypeInteger(column Column) (string, error) {
	return d.generatableColumn("integer", column), nil
}

// Create the column definition for a big integer type.
func (d postgreSQL) TypeBigInteger(column Column) (string, error) {
	return d.generatableColumn("bigint", column), nil
}

// Create the column definition for a medium integer type.
func (d postgreSQL) TypeMediumInteger(column Column) (string, error) {
	return d.generatableColumn("integer", column), nil
}

// Create the column definition for a tiny integer type.
func (d postgreSQL) TypeTinyInteger(column Column) (string, error) {
	return d.generatableColumn("smallint", column), nil
}

// Create the column definition for a small integer type.
func (d postgreSQL) TypeSmallInteger(column Column) (string, error) {
	return d.generatableColumn("smallint", column), nil
}

// Create the column definition for a tiny blob type.
func (d postgreSQL) TypeTinyBlob(column Column) (string, error) {
	return "bytea", nil
}

// Create the column definition for an blob type.
func (d postgreSQL) TypeBlob(column Column) (string, error) {
	return "bytea", nil
}

// Create the column definition for a medium blob type.
func (d postgreSQL) TypeMediumBlob(column Column) (string, error) {
	return "bytea", nil
}

// Create the column definition for a long blob type.
func (d postgreSQL) TypeLongBlob(column Column) (string, error) {
	return "bytea", nil
}

// Create the column definition for a generatable column.
// Create the column definition for a float type.
func (d postgreSQL) TypeFloat(column Column) (string, error) {
	return d.TypeDouble(column)
}

// Create the column definition for a double type.
func (d postgreSQL) TypeDouble(column Column) (string, error) {
	return "double precision", nil
}

// Create the column definition for a real type.
func (d postgreSQL) TypeReal(column Column) (string, error) {
	return "real", nil
}

// Create the column definition for a decimal type.
func (d postgreSQL) TypeDecimal(column Column) (string, error) {
	return fmt.Sprintf("decimal(%d, %d)", column.Precision(), column.Scale()), nil
}

// Create the column definition for a boolean type.
func (d postgreSQL) TypeBoolean(column Column) (string, error) {
	return "boolean", nil
}

// Create the column definition for an enumeration type.
func (d postgreSQL) TypeEnum(column Column) (string, error) {
	return fmt.Sprintf("varchar(255) check (%s in (%s))", d.QuoteIdent(column.Name()), utils.QuoteInterfaces(column.AllowedValues()...)), nil
}

// Create the column definition for an enumeration type.
func (d postgreSQL) TypeSet(column Column) (string, error) {
	return "", errors.New("This database driver not support type 'set'.")
}

// Create the column definition for a json type.
func (d postgreSQL) TypeJson(column Column) (string, error) {
	return "json", nil
}

// Create the column definition for a jsonb type.
func (d postgreSQL) TypeJsonb(column Column) (string, error) {
	return "jsonb", nil
}

// Create the column definition for a date type.
func (d postgreSQL) TypeDate(column Column) (string, error) {
	return "date", nil
}

// Create the column definition for a date-time type.
func (d postgreSQL) TypeDateTime(column Column) (string, error) {
	return d.TypeTimestamp(column)
}

// Create the column definition for a date-time (with time zone) type.
func (d postgreSQL) TypeDateTimeTz(column Column) (string, error) {
	return d.TypeTimestampTz(column)
}

// Create the column definition for a time type.
func (d postgreSQL) TypeTime(column Column) (string, error) {
	var b strings.Builder
	b.WriteString("time")
	if column.Precision() != 0 {
		b.WriteString(fmt.Sprintf("(%d) without time zone", column.Precision()))
	}
	return b.String(), nil
}

// Create the column definition for a time (with time zone) type.
func (d postgreSQL) TypeTimeTz(column Column) (string, error) {
	var b strings.Builder
	b.WriteString("time")
	if column.Precision() != 0 {
		b.WriteString(fmt.Sprintf("(%d) with time zone", column.Precision()))
	}
	return b.String(), nil
}

// Create the column definition for a timestamp type.
func (d postgreSQL) TypeTimestamp(column Column) (string, error) {
	var b strings.Builder
	b.WriteString("timestamp")
	if column.Precision() != 0 {
		b.WriteString(fmt.Sprintf("(%d) without time zone", column.Precision()))
	}
	if column.UseCurrent() {
		b.WriteString(" default CURRENT_TIMESTAMP")
	}
	return b.String(), nil
}

// Create the column definition for a timestamp (with time zone) type.
func (d postgreSQL) TypeTimestampTz(column Column) (string, error) {
	var b strings.Builder
	b.WriteString("timestamp")
	if column.Precision() != 0 {
		b.WriteString(fmt.Sprintf("(%d) with time zone", column.Precision()))
	}
	if column.UseCurrent() {
		b.WriteString(" default CURRENT_TIMESTAMP")
	}
	return b.String(), nil
}

// Create the column definition for a year type.
func (d postgreSQL) TypeYear(column Column) (string, error) {
	return "return $this->typeInteger($column);", nil
}

// Create the column definition for a binary type.
func (d postgreSQL) TypeBinary(column Column) (string, error) {
	return "bytea", nil
}

// Create the column definition for a uuid type.
func (d postgreSQL) TypeUuid(column Column) (string, error) {
	return "uuid", nil
}

// Create the column definition for an IP address type.
func (d postgreSQL) TypeIpAddress(column Column) (string, error) {
	return "inet", nil
}

// Create the column definition for a MAC address type.
func (d postgreSQL) TypeMacAddress(column Column) (string, error) {
	return "macaddr", nil
}

// Create the column definition for a spatial Geometry type.
func (d postgreSQL) TypeGeometry(column Column) (string, error) {
	return d.formatPostGisType("geometry", column)
}

// Create the column definition for a spatial Point type.
func (d postgreSQL) TypePoint(column Column) (string, error) {
	return d.formatPostGisType("point", column)
}

// Create the column definition for a spatial LineString type.
func (d postgreSQL) TypeLineString(column Column) (string, error) {
	return d.formatPostGisType("linestring", column)
}

// Create the column definition for a spatial Polygon type.
func (d postgreSQL) TypePolygon(column Column) (string, error) {
	return d.formatPostGisType("polygon", column)
}

// Create the column definition for a spatial GeometryCollection type.
func (d postgreSQL) TypeGeometryCollection(column Column) (string, error) {
	return d.formatPostGisType("geometrycollection", column)
}

// Create the column definition for a spatial MultiPoint type.
func (d postgreSQL) TypeMultiPoint(column Column) (string, error) {
	return d.formatPostGisType("multipoint", column)
}

// Create the column definition for a spatial MultiLineString type.
func (d postgreSQL) TypeMultiLineString(column Column) (string, error) {
	return d.formatPostGisType("multilinestring", column)
}

// Create the column definition for a spatial MultiPolygon type.
func (d postgreSQL) TypeMultiPolygon(column Column) (string, error) {
	return d.formatPostGisType("multipolygon", column)
}

// Create the column definition for a spatial MultiPolygonZ type.
func (d postgreSQL) TypeMultiPolygonZ(column Column) (string, error) {
	return d.formatPostGisType("multipolygonz", column)
}

func (d postgreSQL) ModifyColumn(column Column) string {
	// 'Collate', 'Increment', 'Nullable', 'Default', 'VirtualAs', 'StoredAs'

	var b strings.Builder

	if column.Collate() != "" {
		b.WriteString(" COLLATE ")
		b.WriteString(utils.QuoteStrings(column.Collate()))
	}

	if column.AutoIncrement() {
		switch column.DataType() {
		case TypeInt, TypeBigInt, TypeMediumInt, TypeTinyInt, TypeSmallInt:
			// b.WriteString(" SERIAL")
		}
	}

	if column.Nullable() {
		b.WriteString(" NULL")
	} else {
		b.WriteString(" NOT NULL")
	}

	if column.UseCurrent() && (column.DataType() == TypeDateTime || column.DataType() == TypeDateTimeTz) {
		value := utils.Express("CURRENT_TIMESTAMP")
		b.WriteString(" DEFAULT ")
		b.WriteString(value.String())
	} else if column.DefaultValue() != nil {
		value := utils.DefaultValue(column.DefaultValue())
		if value != "" {
			b.WriteString(" DEFAULT ")
			b.WriteString(value)
		}
	}

	if column.VirtualAs() != "" {
		b.WriteString(" GENERATED ALWAYS AS (")
		b.WriteString(column.VirtualAs())
		b.WriteString(")")
	}

	if column.StoredAs() != "" {
		b.WriteString(" AS (")
		b.WriteString(column.VirtualAs())
		b.WriteString(") STORED")
	}

	return b.String()
}

// Compile the query to determine the list of tables.
func (d postgreSQL) CompileTableExists(tableName string, tableSchema ...string) (string, error) {
	var b strings.Builder

	b.WriteString("select * from information_schema.tables where table_name = ")
	b.WriteString(utils.QuoteStrings(tableName))

	if len(tableSchema) > 0 {
		b.WriteString(" and table_schema = ")
		b.WriteString(utils.QuoteStrings(tableSchema[0]))
	}

	b.WriteString(" and table_type = 'BASE TABLE'")

	return b.String(), nil
}

// Compile the query to determine the list of columns.
func (d postgreSQL) CompileColumnListing(tableName string, tableSchema ...string) (string, error) {
	var b strings.Builder

	b.WriteString("select column_name from information_schema.columns where table_name = ")
	b.WriteString(utils.QuoteStrings(tableName))

	if len(tableSchema) > 0 {
		b.WriteString(" and table_schema = ")
		b.WriteString(utils.QuoteStrings(tableSchema[0]))
	}

	return b.String(), nil
}

//  Compile a create table command.
func (d postgreSQL) CompileCreate(table Table) (string, error) {
	var b strings.Builder

	if table.IsTemporary() {
		b.WriteString("CREATE TEMPORARY TABLE ")
	} else {
		b.WriteString("CREATE TABLE ")
	}

	if table.Prefix() != "" {
		b.WriteString(d.QuoteIdent(table.Prefix()))
		b.WriteString(".")
	}

	b.WriteString(d.QuoteIdent(table.Name()))
	b.WriteString("(\n")

	// primary key names
	var pkNames []string

	columns := table.Columns()

	for i, column := range columns {

		if column.PrimaryKey() {
			pkNames = append(pkNames, d.QuoteIdent(column.Name()))
		}

		b.WriteString("\t")
		b.WriteString(d.QuoteIdent(column.Name()))
		b.WriteString(" ")

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(columnType)
		b.WriteString(d.ModifyColumn(column))

		if i < len(columns)-1 || len(pkNames) > 0 {
			b.WriteString(",")
		}

		b.WriteString("\n")
	}

	if len(pkNames) > 0 {
		b.WriteString("\tPRIMARY KEY (")
		b.WriteString(strings.Join(pkNames, ", "))
		b.WriteString(")\n")
	}

	b.WriteString(");")

	if table.Comment() != "" {
		b.WriteString("\nCOMMENT ON TABLE ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(" IS ")
		b.WriteString(utils.QuoteStrings(table.Comment()))
		b.WriteString(";")
	}

	for _, column := range columns {
		if column.Comment() != "" {
			b.WriteString("\nCOMMENT ON COLUMN ")
			b.WriteString(wrapTableName(table, d.QuoteIdent))
			b.WriteString(".")
			b.WriteString(d.QuoteIdent(column.Name()))
			b.WriteString(" IS ")
			b.WriteString(utils.QuoteStrings(column.Comment()))
			b.WriteString(";")
		}
	}

	return b.String(), nil
}

//  Compile a create table command.
func (d postgreSQL) CompileModifyColumns(table Table) (string, error) {
	var sql strings.Builder

	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" ")

	columns := table.ChangedColumns()

	for _, column := range columns {
		if column.Rename() == "" {
			continue
		}

		sql.WriteString(b.String())
		sql.WriteString(" RENAME COLUMN")
		sql.WriteString(column.Name())
		sql.WriteString(" TO ")
		sql.WriteString(column.Rename())
		sql.WriteString(";\n")
	}

	b.WriteString("\n")

	var columnName string
	var alterColumnSql strings.Builder

	for i, column := range columns {

		if column.Rename() != "" {
			columnName = column.Rename()
		} else {
			columnName = column.Name()
		}

		alterColumnSql.Reset()
		alterColumnSql.WriteString("\tALTER COLUMN ")
		alterColumnSql.WriteString(d.QuoteIdent(columnName))

		b.WriteString(alterColumnSql.String())

		// Column type
		b.WriteString(" TYPE ")

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(columnType)
		b.WriteString(" USING ")
		b.WriteString(d.QuoteIdent(columnName))
		b.WriteString("::")
		b.WriteString(columnType)
		b.WriteString(",\n")

		// Default value
		if column.DefaultValue() == nil {
			b.WriteString(alterColumnSql.String())
			b.WriteString(" DROP DEFAULT,")
		} else {
			value := utils.DefaultValue(column.DefaultValue())

			if value != "" {
				b.WriteString(alterColumnSql.String())
				b.WriteString(" SET DEFAULT ")
				b.WriteString(utils.DefaultValue(column.DefaultValue()))
			}
		}

		// Nullable
		b.WriteString(alterColumnSql.String())

		if column.Nullable() {
			b.WriteString(" DROP NOT NULL")
		} else {
			b.WriteString(" SET NOT NULL")
		}

		if i < len(columns)-1 {
			b.WriteString(",")
			b.WriteString("\n")
		}
	}

	b.WriteString(";\n")

	// Column comment
	for i, column := range columns {

		if column.Rename() != "" {
			columnName = column.Rename()
		} else {
			columnName = column.Name()
		}

		b.WriteString("COMMENT ON COLUMN ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(".")
		b.WriteString(d.QuoteIdent(columnName))
		b.WriteString(" IS ")
		b.WriteString(utils.QuoteStrings(column.Comment()))
		b.WriteString(";")

		if i < len(columns)-1 {
			b.WriteString("\n")
		}
	}

	sql.WriteString(b.String())

	return sql.String(), nil
}

//  Compile a modify column command.
func (d postgreSQL) CompileModifyColumn(table Table, columnName string) (string, error) {

	var column Column

	columns := table.ChangedColumns()

	for _, col := range columns {
		if col.Name() == columnName {
			column = col
			break
		}
	}

	if column == nil {
		return "", errors.New(fmt.Sprintf("dbx: not found changed column '%s'", columnName))
	}

	var sql strings.Builder

	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" ")

	if column.Rename() != "" {
		sql.WriteString(b.String())
		sql.WriteString(" RENAME COLUMN")
		sql.WriteString(column.Name())
		sql.WriteString(" TO ")
		sql.WriteString(column.Rename())
		sql.WriteString(";\n")
	}

	b.WriteString("\n")

	var alterColumnSql strings.Builder

	if column.Rename() != "" {
		columnName = column.Rename()
	} else {
		columnName = column.Name()
	}

	alterColumnSql.Reset()
	alterColumnSql.WriteString("\tALTER COLUMN ")
	alterColumnSql.WriteString(d.QuoteIdent(columnName))

	b.WriteString(alterColumnSql.String())

	// Column type
	b.WriteString(" TYPE ")

	columnType, err := ColumnType(d, column)

	if err != nil {
		return "", err
	}

	b.WriteString(columnType)
	b.WriteString(" USING ")
	b.WriteString(d.QuoteIdent(columnName))
	b.WriteString("::")
	b.WriteString(columnType)
	b.WriteString(",\n")

	// Nullable
	b.WriteString(alterColumnSql.String())

	if column.Nullable() {
		b.WriteString(" DROP NOT NULL,\n")
	} else {
		b.WriteString(" SET NOT NULL,\n")
	}

	// Default value
	if column.DefaultValue() == nil {
		b.WriteString(alterColumnSql.String())
		b.WriteString(" DROP DEFAULT")
	} else {
		value := utils.DefaultValue(column.DefaultValue())

		if value != "" {
			b.WriteString(alterColumnSql.String())
			b.WriteString(" SET DEFAULT ")
			b.WriteString(utils.DefaultValue(column.DefaultValue()))
		}
	}

	b.WriteString(";\n")

	// Column comment
	for i, column := range columns {

		if column.Rename() != "" {
			columnName = column.Rename()
		} else {
			columnName = column.Name()
		}

		b.WriteString("COMMENT ON COLUMN ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(".")
		b.WriteString(d.QuoteIdent(columnName))
		b.WriteString(" IS ")
		b.WriteString(utils.QuoteStrings(column.Comment()))
		b.WriteString(";")

		if i < len(columns)-1 {
			b.WriteString("\n")
		}
	}

	sql.WriteString(b.String())

	return sql.String(), nil
}

// Compile add columns.
func (d postgreSQL) CompileAddColumns(table Table) (string, error) {

	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" \n")

	columns := table.AddedColumns()

	for i, column := range columns {

		b.WriteString("\tADD COLUMN ")
		b.WriteString(d.QuoteIdent(column.Name()))

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(" ")
		b.WriteString(columnType)

		if column.Nullable() == false {
			b.WriteString(" NOT NULL")
		}

		// Default value
		if column.DefaultValue() != nil {
			value := utils.DefaultValue(column.DefaultValue())

			if value != "" {
				b.WriteString(" DEFAULT ")
				b.WriteString(value)
			}
		}

		if i < len(columns)-1 {
			b.WriteString(",")
			b.WriteString("\n")
		}
	}

	b.WriteString(";\n")

	// Column comment
	for i, column := range columns {

		if column.Comment() == "" {
			continue
		}

		b.WriteString("COMMENT ON COLUMN ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(".")
		b.WriteString(d.QuoteIdent(column.Name()))
		b.WriteString(" IS ")
		b.WriteString(utils.QuoteStrings(column.Comment()))
		b.WriteString(";")

		if i < len(columns)-1 {
			b.WriteString("\n")
		}
	}

	return b.String(), nil
}

// Compile add a column.
func (d postgreSQL) CompileAddColumn(table Table, columnName string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" ")

	var column Column
	columns := table.AddedColumns()

	for _, col := range columns {
		if col.Name() == columnName {
			column = col
			break
		}
	}

	if column == nil {
		return "", errors.New(fmt.Sprintf("dbx: not found added column '%s'", columnName))
	}

	b.WriteString("ADD COLUMN ")
	b.WriteString(d.QuoteIdent(column.Name()))

	columnType, err := ColumnType(d, column)

	if err != nil {
		return "", err
	}

	b.WriteString(" ")
	b.WriteString(columnType)

	if column.Nullable() == false {
		b.WriteString(" NOT NULL")
	}

	// Default value
	if column.DefaultValue() != nil {
		value := utils.DefaultValue(column.DefaultValue())

		if value != "" {
			b.WriteString(" DEFAULT ")
			b.WriteString(value)
		}
	}

	b.WriteString(";\n")

	// Column comment
	for i, column := range columns {

		if column.Comment() == "" {
			continue
		}

		b.WriteString("COMMENT ON COLUMN ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(".")
		b.WriteString(d.QuoteIdent(column.Name()))
		b.WriteString(" IS ")
		b.WriteString(utils.QuoteStrings(column.Comment()))
		b.WriteString(";")

		if i < len(columns)-1 {
			b.WriteString("\n")
		}
	}

	return b.String(), nil
}

// Compile a primary key command.
func (d postgreSQL) CompilePrimaryKey(table Table, columnNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" ADD PRIMARY KEY (")
	b.WriteString(utils.QuoteIdents(columnNames, d.QuoteIdent))
	b.WriteString(");")

	return b.String(), nil
}

// Compile a drop primary key command.
func (d postgreSQL) CompileDropPrimaryKey(table Table) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP PRIMARY KEY;")

	return b.String(), nil
}

// Compile an index creation command.
func (d postgreSQL) CompileIndex(table Table, index Index) (string, error) {
	var b strings.Builder

	if index.Type() == UniqueIndex {
		b.WriteString("ALTER TABLE ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(" ADD CONSTRAINT ")
		b.WriteString(d.QuoteIdent(index.Name()))
		b.WriteString(" UNIQUE ")

		b.WriteString("(")
		b.WriteString(utils.QuoteIdents(index.ColumnNames(), d.QuoteIdent))
		b.WriteString(");")
	} else {
		b.WriteString("CREATE INDEX ")
		b.WriteString(d.QuoteIdent(index.Name()))
		b.WriteString(" ON ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))

		algorithm := index.Algorithm()

		if index.Type() == SpatialIndex || index.Type() == GistIndex {
			algorithm = "GIST"
		}

		if algorithm != "" {
			b.WriteString(" USING ")
			b.WriteString(index.Algorithm())
		}

		b.WriteString("(")
		b.WriteString(utils.QuoteIdents(index.ColumnNames(), d.QuoteIdent))
		b.WriteString(");")
	}

	return b.String(), nil
}

// Compile a drop index command.
func (d postgreSQL) CompileDropIndex(table Table, indexName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP INDEX ")
	b.WriteString(d.wrapIndexName(table, indexName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop unique index command.
func (d postgreSQL) CompileDropUnique(table Table, indexName string) (string, error) {
	var b strings.Builder
	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP CONSTRAINT ")
	b.WriteString(d.QuoteIdent(indexName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop spatial index command.
func (d postgreSQL) CompileDropSpatialIndex(table Table, indexName string) (string, error) {
	return d.CompileDropIndex(table, indexName)
}

// Compile a drop foreign index command.
func (d postgreSQL) CompileDropForeign(table Table, indexName string) (string, error) {
	var b strings.Builder
	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP CONSTRAINT ")
	b.WriteString(d.QuoteIdent(indexName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop table command.
func (d postgreSQL) CompileDrop(tableName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP TABLE ")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop table (if exists) command.
func (d postgreSQL) CompileDropIfExists(tableName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP TABLE IF EXISTS")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop column command.
func (d postgreSQL) CompileDropColumn(table Table, columnNames []string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" \n")

	for i, name := range columnNames {
		b.WriteString("DROP COLUMN ")
		b.WriteString(d.QuoteIdent(name))

		if i < len(columnNames)-1 {
			b.WriteString(",\n")
		} else {
			b.WriteString(";")
		}
	}

	return b.String(), nil
}

// Compile a rename table command.
func (d postgreSQL) CompileRenameTable(table Table, toName string) (string, error) {
	var b strings.Builder

	b.WriteString("RENAME TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" TO ")
	b.WriteString(d.QuoteIdent(toName))
	b.WriteString(";")

	return b.String(), nil
}

// Compile a rename index command.
func (d postgreSQL) CompileRenameIndex(table Table, from string, to string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" RENAME INDEX ")
	b.WriteString(d.QuoteIdent(from))
	b.WriteString(" TO ")
	b.WriteString(d.QuoteIdent(to))
	b.WriteString(";")

	return b.String(), nil
}

// Compile the SQL needed to drop all tables.
func (d postgreSQL) CompileDropAllTables(tableNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("DROP TABLE ")

	for i, name := range tableNames {
		tableNames[i] = d.QuoteIdent(name)
	}

	b.WriteString(strings.Join(tableNames, ","))
	b.WriteString(" CASCADE;")

	return b.String(), nil
}

// Compile the SQL needed to drop all views.
func (d postgreSQL) CompileDropAllViews(viewNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("DROP VIEW ")

	for i, name := range viewNames {
		viewNames[i] = d.QuoteIdent(name)
	}

	b.WriteString(strings.Join(viewNames, ","))
	b.WriteString(" CASCADE;")

	return b.String(), nil
}

// Compile the SQL needed to drop all types. [PostgreSQL]
func (d postgreSQL) CompileDropAllTypes(typeNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("DROP TYPE ")

	for i, name := range typeNames {
		typeNames[i] = d.QuoteIdent(name)
	}

	b.WriteString(strings.Join(typeNames, ","))
	b.WriteString(" CASCADE;")

	return b.String(), nil
}

// Compile the SQL needed to retrieve all table names.
func (d postgreSQL) CompileGetAllTables(schemaNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("select tablename from pg_catalog.pg_tables where schemaname in (")

	for i, name := range schemaNames {
		schemaNames[i] = utils.QuoteStrings(name)
	}

	b.WriteString(strings.Join(schemaNames, ","))
	b.WriteString(")")

	return b.String(), nil
}

// Compile the SQL needed to retrieve all view names.
func (d postgreSQL) CompileGetAllViews(schemaNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("select viewname from pg_catalog.pg_views where schemaname in (")

	for i, name := range schemaNames {
		schemaNames[i] = utils.QuoteStrings(name)
	}

	b.WriteString(strings.Join(schemaNames, ","))
	b.WriteString(")")

	return b.String(), nil
}

// Compile the SQL needed to retrieve all type names. [PostgreSQL]
func (d postgreSQL) CompileGetAllTypes() (string, error) {
	return `select distinct pg_type.typname from pg_type inner join pg_enum on pg_enum.enumtypid = pg_type.oid`, nil
}

// Compile the SQL needed to rebuild the database. [SQLite]
func (d postgreSQL) CompileRebuild() (string, error) {
	return "", errors.New("postgreSQL not support 'CompileRebuild'")
}

// Compile the command to enable foreign key constraints.
func (d postgreSQL) CompileEnableForeignKeyConstraints() (string, error) {
	return `SET CONSTRAINTS ALL IMMEDIATE;`, nil
}

// Compile the command to disable foreign key constraints.
func (d postgreSQL) CompileDisableForeignKeyConstraints() (string, error) {
	return `SET CONSTRAINTS ALL DEFERRED;`, nil
}

func (d postgreSQL) compileColumnListing(tableName string, tableSchema ...string) (string, error) {
	var version string

	_, err := d.query("select version()", &version)

	if err != nil {
		return "", errors.WithMessage(err, "dbx: error on query database version")
	}

	versionNum, err := strconv.ParseFloat(strings.Split(version, " ")[1], 64)

	if err != nil {
		return "", errors.WithMessagef(err, "dbx: error on convert database version '%s' to integer", strings.Split(version, " ")[1])
	}

	var orIdentity = ""

	if versionNum >= 12 {
		orIdentity = "OR attidentity != ''"
	}

	var sql = `
		SELECT
			d.nspname AS table_schema,
			c.relname AS table_name,
			a.attname AS column_name,
			COALESCE(td.typname, tb.typname, t.typname) AS data_type,
			COALESCE(td.typtype, tb.typtype, t.typtype) AS type_type,
			a.attlen AS character_maximum_length,
			pg_catalog.col_description(c.oid, a.attnum) AS column_comment,
			a.atttypmod AS modifier,
			a.attnotnull = false AS is_nullable,
			CAST(pg_get_expr(ad.adbin, ad.adrelid) AS varchar) AS column_default,
			coalesce(pg_get_expr(ad.adbin, ad.adrelid) ~ 'nextval',false) ` + orIdentity + ` AS is_autoinc,
			pg_get_serial_sequence(quote_ident(d.nspname) || '.' || quote_ident(c.relname), a.attname) AS sequence_name,
			CASE WHEN COALESCE(td.typtype, tb.typtype, t.typtype) = 'e'::char
				THEN array_to_string((SELECT array_agg(enumlabel) FROM pg_enum WHERE enumtypid = COALESCE(td.oid, tb.oid, a.atttypid))::varchar[], ',')
				ELSE NULL
			END AS enum_values,
			CASE atttypid
				 WHEN 21 /*int2*/ THEN 16
				 WHEN 23 /*int4*/ THEN 32
				 WHEN 20 /*int8*/ THEN 64
				 WHEN 1700 /*numeric*/ THEN
					  CASE WHEN atttypmod = -1
					   THEN null
					   ELSE ((atttypmod - 4) >> 16) & 65535
					   END
				 WHEN 700 /*float4*/ THEN 24 /*FLT_MANT_DIG*/
				 WHEN 701 /*float8*/ THEN 53 /*DBL_MANT_DIG*/
				 ELSE null
			  END   AS numeric_precision,
			  CASE
				WHEN atttypid IN (21, 23, 20) THEN 0
				WHEN atttypid IN (1700) THEN
				CASE
					WHEN atttypmod = -1 THEN null
					ELSE (atttypmod - 4) & 65535
				END
				   ELSE null
			  END AS numeric_scale,
			CAST(
					 information_schema._pg_char_max_length(information_schema._pg_truetypid(a, t), information_schema._pg_truetypmod(a, t))
					 AS numeric
			) AS size,
			a.attnum = any (ctp.conkey) as is_pkey,
			pg_get_constraintdef(ctc.oid) as check,
			COALESCE(NULLIF(a.attndims, 0), NULLIF(t.typndims, 0), (t.typcategory='A')::int) AS dimension
		FROM
			pg_class c
			LEFT JOIN pg_attribute a ON a.attrelid = c.oid
			LEFT JOIN pg_attrdef ad ON a.attrelid = ad.adrelid AND a.attnum = ad.adnum
			LEFT JOIN pg_type t ON a.atttypid = t.oid
			LEFT JOIN pg_type tb ON (a.attndims > 0 OR t.typcategory='A') AND t.typelem > 0 AND t.typelem = tb.oid OR t.typbasetype > 0 AND t.typbasetype = tb.oid
			LEFT JOIN pg_type td ON t.typndims > 0 AND t.typbasetype > 0 AND tb.typelem = td.oid
			LEFT JOIN pg_namespace d ON d.oid = c.relnamespace
			LEFT JOIN pg_constraint ctp ON ctp.conrelid = c.oid AND ctp.contype = 'p'
			LEFT JOIN pg_constraint ctc ON ctc.conrelid = c.oid AND ctc.contype = 'c' and a.attnum = any(ctc.conkey)
		WHERE
			a.attnum > 0 AND t.typname != '' AND NOT a.attisdropped
			AND c.relname = '` + tableName + `'`

	if len(tableSchema) > 0 {
		sql += `AND d.nspname = '` + tableSchema[0] + `'`
	}

	sql += ` ORDER BY a.attnum;`

	return sql, nil
}

// Load table columns
func (d postgreSQL) LoadColumns(tableName string, tableSchema ...string) ([]Column, error) {

	sql, err := d.compileColumnListing(tableName, tableSchema...)

	if err != nil {
		return nil, err
	}

	var columnInfos []postgresColumnInfo

	_, err = d.query(sql, &columnInfos)

	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("dbx: error on query columns of the table %s", utils.QuoteStrings(tableName)))
	}

	var columns []Column

	for _, info := range columnInfos {

		dataType, err := DataTypeMapper(info.DataType)

		if err != nil {
			return nil, err
		}

		column := newColumn(info.Name, dataType)

		// not added
		column.added = false

		column.nullable = info.Nullable

		if info.Length.Valid {
			column.length = int(info.Length.Int64)
		}

		if info.Precision.Valid {
			column.precision = int(info.Precision.Int64)
		}

		if info.Scale.Valid {
			column.scale = int(info.Scale.Int64)
		}

		// 丑陋的代码
		if info.Default.Valid {
			if info.Default.String == "CURRENT_TIMESTAMP" {
				column.defaultValue = utils.Express("CURRENT_TIMESTAMP")
				column.useCurrent = true
			} else if strings.HasPrefix(info.Default.String, "'") && strings.Contains(info.Default.String, "'::") {
				column.defaultValue = info.Default.String[1:strings.LastIndex(info.Default.String, "'::")]
			} else if strings.HasPrefix(info.Default.String, "nextval('") {
				// auto increment
			} else {
				column.defaultValue = info.Default.String
			}
		}

		if info.Comment.Valid {
			column.comment = info.Comment.String
		}

		if info.PrimaryKey {
			column.primaryKey = true
		}

		if info.AutoInc {
			column.autoIncrement = true
		}

		// 丑陋的代码，但是postgresql中的枚举需要创建枚举类型，暂时没有想到好的办法...
		if info.Check.Valid {
			prefix := "CHECK (((tag)::text = ANY ((ARRAY["
			suffix := "])::text[])))"
			if strings.HasPrefix(info.Check.String, prefix) && strings.HasSuffix(info.Check.String, suffix) {
				info.Check.String = info.Check.String[len(prefix) : len(info.Check.String)-len(suffix)]

				var values []interface{}

				for _, item := range strings.Split(info.Check.String, ",") {
					item = strings.TrimSpace(item)
					item = item[1:strings.LastIndex(item, "'::")]
					values = append(values, item)
				}

				column.allowedValues = values

				if len(values) > 0 {
					column.dataType = TypeEnum
				}
			}
		}

		// save old column
		column.store()

		columns = append(columns, column)
	}

	return columns, nil
}

// Load table from the database
func (d postgreSQL) LoadTable(tableName string, tableSchema ...string) (Table, error) {
	var sql strings.Builder

	sql.WriteString("select ns.nspname as table_schema, c.relname as table_name, cast(obj_description(c.relfilenode,'pg_class') as varchar) as table_comment")
	sql.WriteString(" from pg_class as c")
	sql.WriteString(" left join pg_namespace as ns on c.relnamespace = ns.oid")
	sql.WriteString(" where c.relname = ?")

	var args = []interface{}{tableName}

	if len(tableSchema) > 0 {
		sql.WriteString(" and ns.nspname = ?")
		args = append(args, tableSchema[0])
	}

	var info postgresTableInfo

	_, err := d.query(sql.String(), &info, args...)

	if err != nil {
		return nil, err
	}

	columns, err := d.LoadColumns(info.Name, info.Schema)

	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("error on query columns of the table %s", utils.QuoteStrings(info.Name)))
	}

	table := NewTable(info.Name, columns...)

	table.schema = info.Schema

	if info.Comment.Valid {
		table.SetComment(info.Comment.String)
	}

	// load from databse
	table.added = false

	return table, nil
}

// Load table constraints from the database
func (d postgreSQL) LoadTableConstraints(tableName string, tableSchema ...string) (TableConstraints, error) {

	var tableConstraints TableConstraints

	var sql strings.Builder
	sql.WriteString(`
SELECT
    "c"."conname" AS "name",
    "a"."attname" AS "column_name",
    "c"."contype" AS "type",
    "ftcns"."nspname" AS "foreign_table_schema",
    "ftc"."relname" AS "foreign_table_name",
    "fa"."attname" AS "foreign_column_name",
    "c"."confupdtype" AS "on_update",
    "c"."confdeltype" AS "on_delete",
    pg_get_constraintdef("c"."oid") AS "check_expr"
FROM "pg_class" AS "tc"
INNER JOIN "pg_namespace" AS "tcns"
    ON "tcns"."oid" = "tc"."relnamespace"
INNER JOIN "pg_constraint" AS "c"
    ON "c"."conrelid" = "tc"."oid"
INNER JOIN "pg_attribute" AS "a"
    ON "a"."attrelid" = "c"."conrelid" AND "a"."attnum" = ANY ("c"."conkey")
LEFT JOIN "pg_class" AS "ftc"
    ON "ftc"."oid" = "c"."confrelid"
LEFT JOIN "pg_namespace" AS "ftcns"
    ON "ftcns"."oid" = "ftc"."relnamespace"
LEFT JOIN "pg_attribute" "fa"
    ON "fa"."attrelid" = "c"."confrelid" AND "fa"."attnum" = ANY ("c"."confkey")
	`)

	sql.WriteString(` WHERE "tc"."relname" = `)
	sql.WriteString(utils.QuoteStrings(tableName))

	if len(tableSchema) > 0 {
		sql.WriteString(` AND "tcns"."nspname" = `)
		sql.WriteString(utils.QuoteStrings(tableSchema[0]))
	}

	sql.WriteString(` ORDER BY "a"."attnum" ASC, "fa"."attnum" ASC`)

	var constraints TableConstraintInfos

	_, err := d.query(sql.String(), &constraints)

	if err != nil {
		return tableConstraints, err
	}

	var actionTypes = map[string]string{
		"a": "NO ACTION",
		"r": "RESTRICT",
		"c": "CASCADE",
		"n": "SET NULL",
		"d": "SET DEFAULT",
	}

	for typ, nameGroup := range constraints.Group() {
		for name, items := range nameGroup {
			switch typ {
			case "p":
				tableConstraints.SetPrimaryKey(constraint.PrimaryKey{
					Name:        name,
					ColumnNames: items.ColumnNames(),
				})
			case "f":
				var fkey = constraint.ForeignKey{
					Name:               name,
					ColumnNames:        items.ColumnNames(),
					ForeignSchemaName:  items[0].ForeignTableSchema.String,
					ForeignTableName:   items[0].ForeignTableName.String,
					ForeignColumnNames: items.ForeignColumnNames(),
				}

				if actionType, ok := actionTypes[items[0].OnDelete.String]; ok {
					fkey.OnDelete = actionType
				} else {
					fkey.OnDelete = ""
				}

				if actionType, ok := actionTypes[items[0].OnUpdate.String]; ok {
					fkey.OnUpdate = actionType
				} else {
					fkey.OnUpdate = ""
				}

				tableConstraints.AddForeignKeys(fkey)
			case "u":
				tableConstraints.AddUniques(constraint.Unique{
					Name:        name,
					ColumnNames: items.ColumnNames(),
				})
			case "c":
				tableConstraints.AddChecks(constraint.Check{
					Name:        name,
					ColumnNames: items.ColumnNames(),
					Expression:  items[0].CheckExpr.String,
				})
			}
		}
	}

	return tableConstraints, nil
}
