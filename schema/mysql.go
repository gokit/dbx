package schema

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/gokit/dbx/schema/constraint"
	"github.com/gokit/dbx/utils"
	"strings"
)

type mysqlColumnInfo struct {
	Name       string         `db:"COLUMN_NAME"`
	Position   int            `db:"ORDINAL_POSITION"`
	DataType   string         `db:"DATA_TYPE"`
	Length     sql.NullInt64  `db:"CHARACTER_MAXIMUM_LENGTH"`
	Precision  sql.NullInt64  `db:"NUMERIC_PRECISION"`
	Scale      sql.NullInt64  `db:"NUMERIC_SCALE"`
	Nullable   string         `db:"IS_NULLABLE"`
	Default    sql.NullString `db:"COLUMN_DEFAULT"`
	Charset    sql.NullString `db:"CHARACTER_SET_NAME"`
	Collation  sql.NullString `db:"COLLATION_NAME"`
	ColumnType string         `db:"COLUMN_TYPE"`
	Key        sql.NullString `db:"COLUMN_KEY"`
	Extra      sql.NullString `db:"EXTRA"`
	Comment    sql.NullString `db:"COLUMN_COMMENT"`
}

type mysqlTableInfo struct {
	Name      string         `db:"TABLE_NAME"`
	Schema    string         `db:"TABLE_SCHEMA"`
	Engine    sql.NullString `db:"ENGINE"`
	Collation sql.NullString `db:"TABLE_COLLATION"`
	Options   sql.NullString `db:"CREATE_OPTIONS"`
	Comment   sql.NullString `db:"TABLE_COMMENT"`
}

type mysql struct {
	query Query
}

func (d mysql) QuoteIdent(s string) string {
	return utils.QuoteIdent(s, "`")
}

// Create the column definition for a char type.
func (d mysql) TypeChar(column Column) (string, error) {
	return fmt.Sprintf("char(%d)", column.Length()), nil
}

// Create the column definition for a string type.
func (d mysql) TypeString(column Column) (string, error) {
	return fmt.Sprintf("varchar(%d)", column.Length()), nil
}

// Create the column definition for a text type.
func (d mysql) TypeText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a medium text type.
func (d mysql) TypeMediumText(column Column) (string, error) {
	return "mediumtext", nil
}

// Create the column definition for a long text type.
func (d mysql) TypeLongText(column Column) (string, error) {
	return "longtext", nil
}

// Create the column definition for a big integer type.
func (d mysql) TypeBigInteger(column Column) (string, error) {
	return "bigint", nil
}

// Create the column definition for an integer type.
func (d mysql) TypeInteger(column Column) (string, error) {
	return "int", nil
}

// Create the column definition for a medium integer type.
func (d mysql) TypeMediumInteger(column Column) (string, error) {
	return "mediumint", nil
}

// Create the column definition for a tiny integer type.
func (d mysql) TypeTinyInteger(column Column) (string, error) {
	return "tinyint", nil
}

// Create the column definition for a small integer type.
func (d mysql) TypeSmallInteger(column Column) (string, error) {
	return "smallint", nil
}

// Create the column definition for a tiny blob type.
func (d mysql) TypeTinyBlob(column Column) (string, error) {
	return "tinyblob", nil
}

// Create the column definition for an blob type.
func (d mysql) TypeBlob(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for a medium blob type.
func (d mysql) TypeMediumBlob(column Column) (string, error) {
	return "mediumblob", nil
}

// Create the column definition for a long blob type.
func (d mysql) TypeLongBlob(column Column) (string, error) {
	return "longblob", nil
}

// Create the column definition for a float type.
func (d mysql) TypeFloat(column Column) (string, error) {
	if column.Precision() > 0 {
		return fmt.Sprintf("float(%d, %d)", column.Precision(), column.Scale()), nil
	}
	return "float", nil
}

// Create the column definition for a double type.
func (d mysql) TypeDouble(column Column) (string, error) {
	if column.Precision() > 0 {
		return fmt.Sprintf("double(%d, %d)", column.Precision(), column.Scale()), nil
	}
	return "double", nil
}

// Create the column definition for a decimal type.
func (d mysql) TypeDecimal(column Column) (string, error) {
	return fmt.Sprintf("decimal(%d, %d)", column.Precision(), column.Scale()), nil
}

// Create the column definition for a boolean type.
func (d mysql) TypeBoolean(column Column) (string, error) {
	return "tinyint(1)", nil
}

// Create the column definition for an enumeration type.
func (d mysql) TypeEnum(column Column) (string, error) {
	return fmt.Sprintf("enum(%v)", utils.QuoteInterfaces(column.AllowedValues()...)), nil
}

// Create the column definition for a set enumeration type.
func (d mysql) TypeSet(column Column) (string, error) {
	return fmt.Sprintf("set(%v)", utils.QuoteInterfaces(column.AllowedValues()...)), nil
}

// Create the column definition for a json type.
func (d mysql) TypeJson(column Column) (string, error) {
	return "json", nil
}

// Create the column definition for a jsonb type.
func (d mysql) TypeJsonb(column Column) (string, error) {
	return "json", nil
}

// Create the column definition for a date type.
func (d mysql) TypeDate(column Column) (string, error) {
	return "date", nil
}

// Create the column definition for a date-time type.
func (d mysql) TypeDateTime(column Column) (string, error) {
	var res strings.Builder

	if column.Precision() > 0 {
		res.WriteString(fmt.Sprintf("datetime(%d)", column.Precision()))
	} else {
		res.WriteString("datetime")
	}

	return res.String(), nil
}

// Create the column definition for a date-time (with time zone) type.
func (d mysql) TypeDateTimeTz(column Column) (string, error) {
	return d.TypeDateTime(column)
}

// Create the column definition for a time type.
func (d mysql) TypeTime(column Column) (string, error) {
	if column.Precision() > 0 {
		return fmt.Sprintf("time(%d)", column.Precision()), nil
	} else {
		return "time", nil
	}
}

// Create the column definition for a time (with time zone) type.
func (d mysql) TypeTimeTz(column Column) (string, error) {
	return d.TypeTime(column)
}

// Create the column definition for a timestamp type.
func (d mysql) TypeTimestamp(column Column) (string, error) {

	var res strings.Builder

	if column.Precision() > 0 {
		res.WriteString(fmt.Sprintf("timestamp(%d)", column.Precision()))

		if column.UseCurrent() {
			res.WriteString(fmt.Sprintf(" default CURRENT_TIMESTAMP(%d)", column.Precision()))
		}
	} else {
		res.WriteString("timestamp")

		if column.UseCurrent() {
			res.WriteString(" default CURRENT_TIMESTAMP")
		}
	}

	return res.String(), nil
}

// Create the column definition for a timestamp (with time zone) type.
func (d mysql) TypeTimestampTz(column Column) (string, error) {
	return d.TypeTimestamp(column)
}

// Create the column definition for a year type.
func (d mysql) TypeYear(column Column) (string, error) {
	return "year", nil
}

// Create the column definition for a binary type.
func (d mysql) TypeBinary(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for a uuid type.
func (d mysql) TypeUuid(column Column) (string, error) {
	return "char(36)", nil
}

// Create the column definition for an IP address type.
func (d mysql) TypeIpAddress(column Column) (string, error) {
	return "varchar(45)", nil
}

// Create the column definition for a MAC address type.
func (d mysql) TypeMacAddress(column Column) (string, error) {
	return "varchar(17)", nil
}

// Create the column definition for a spatial Geometry type.
func (d mysql) TypeGeometry(column Column) (string, error) {
	return "geometry", nil
}

// Create the column definition for a spatial Point type.
func (d mysql) TypePoint(column Column) (string, error) {
	return "point", nil
}

// Create the column definition for a spatial LineString type.
func (d mysql) TypeLineString(column Column) (string, error) {
	return "linestring", nil
}

// Create the column definition for a spatial Polygon type.
func (d mysql) TypePolygon(column Column) (string, error) {
	return "polygon", nil
}

// Create the column definition for a spatial GeometryCollection type.
func (d mysql) TypeGeometryCollection(column Column) (string, error) {
	return "geometrycollection", nil
}

// Create the column definition for a spatial MultiPoint type.
func (d mysql) TypeMultiPoint(column Column) (string, error) {
	return "multipoint", nil
}

// Create the column definition for a spatial MultiLineString type.
func (d mysql) TypeMultiLineString(column Column) (string, error) {
	return "multilinestring", nil
}

// Create the column definition for a spatial MultiPolygon type.
func (d mysql) TypeMultiPolygon(column Column) (string, error) {
	return "multipolygon", nil
}

// Create the column definition for a generated, computed column type.
func (d mysql) TypeComputed(column Column) (string, error) {
	return "", errors.New("This database driver requires a type, see the virtualAs / storedAs modifiers.")
}

// Modify the column
func (d mysql) ModifyColumn(column Column) string {
	// 'Unsigned', 'Charset', 'Collate', 'VirtualAs', 'StoredAs', 'Nullable', 'Srid', 'Default', 'Increment', 'Comment', 'After', 'First'

	var b strings.Builder

	if column.Unsigned() {
		b.WriteString(" UNSIGNED")
	}

	if column.Charset() != "" {
		b.WriteString(" CHARACTER SET ")
		b.WriteString(column.Charset())
	}

	if column.Collate() != "" {
		b.WriteString(" COLLATE ")
		b.WriteString(utils.QuoteStrings(column.Collate()))
	}

	if column.VirtualAs() != "" {
		b.WriteString(" AS (")
		b.WriteString(column.VirtualAs())
		b.WriteString(")")
	}

	if column.StoredAs() != "" {
		b.WriteString(" AS (")
		b.WriteString(column.VirtualAs())
		b.WriteString(") STORED")
	}

	if column.VirtualAs() == "" && column.StoredAs() == "" {
		if column.Nullable() {
			b.WriteString(" NULL")
		} else {
			b.WriteString(" NOT NULL")
		}
	} else if column.Nullable() == false {
		b.WriteString(" NOT NULL")
	}

	if column.Srid() > 0 {
		b.WriteString(fmt.Sprintf(" SRID %d", column.Srid()))
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

	if column.AutoIncrement() {
		switch column.DataType() {
		case TypeInt, TypeBigInt, TypeMediumInt, TypeTinyInt, TypeSmallInt:
			b.WriteString(" AUTO_INCREMENT")
		}
	}

	if column.Comment() != "" {
		b.WriteString(" COMMENT '")
		b.WriteString(utils.Addslashes(column.Comment()))
		b.WriteString("'")
	}

	if column.After() != "" {
		b.WriteString(" AFTER ")
		b.WriteString(d.QuoteIdent(column.After()))
	}

	if column.First() != "" {
		b.WriteString(" First")
	}

	return b.String()
}

// Compile the query to determine the list of tables.
func (d mysql) CompileTableExists(tableName string, tableSchema ...string) (string, error) {
	var b strings.Builder

	b.WriteString("select count(*) from information_schema.tables where table_name = ")
	b.WriteString(utils.QuoteStrings(tableName))

	if len(tableSchema) > 0 {
		b.WriteString(" and table_schema = ")
		b.WriteString(utils.QuoteStrings(tableSchema[0]))
	}

	b.WriteString(" and table_type = 'BASE TABLE'")

	return b.String(), nil
}

// Compile the query to determine the list of columns.
func (d mysql) CompileColumnListing(tableName string, tableSchema ...string) (string, error) {
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
func (d mysql) CompileCreate(table Table) (string, error) {
	var b strings.Builder

	if table.IsTemporary() {
		b.WriteString("CREATE TEMPORARY TABLE ")
	} else {
		b.WriteString("CREATE TABLE ")
	}

	b.WriteString(wrapTableName(table, d.QuoteIdent))
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

	b.WriteString(")")

	if table.Engine() != "" {
		b.WriteString(" ENGINE=")
		b.WriteString(table.Engine())
	}

	if table.Charset() != "" {
		b.WriteString(" CHARSET=")
		b.WriteString(table.Charset())
	}

	if table.Collation() != "" {
		b.WriteString(" COLLATE=")
		b.WriteString(table.Collation())
	}

	if table.Comment() != "" {
		b.WriteString(" COMMENT=")
		b.WriteString(utils.QuoteStrings(table.Comment()))
	}

	b.WriteString(";")

	return b.String(), nil
}

//  Compile a modify table command.
func (d mysql) CompileModifyColumns(table Table) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" \n")

	columns := table.ChangedColumns()

	for i, column := range columns {

		if column.Rename() == "" {
			b.WriteString("MODIFY COLUMN ")
			b.WriteString(d.QuoteIdent(column.Name()))
			b.WriteString(" ")
		} else {
			b.WriteString("CHANGE COLUMN ")
			b.WriteString(d.QuoteIdent(column.Name()))
			b.WriteString(" ")
			b.WriteString(d.QuoteIdent(column.Rename()))
			b.WriteString(" ")
		}

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(columnType)
		b.WriteString(d.ModifyColumn(column))

		if i < len(columns)-1 {
			b.WriteString(",")
			b.WriteString("\n")
		}
	}

	b.WriteString(";")

	return b.String(), nil
}

//  Compile a modify column command.
func (d mysql) CompileModifyColumn(table Table, columnName string) (string, error) {
	var b strings.Builder

	columns := table.ChangedColumns()

	for _, column := range columns {

		if column.Name() != columnName {
			continue
		}

		b.WriteString("ALTER TABLE ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(" ")

		if column.Rename() == "" {
			b.WriteString("MODIFY COLUMN ")
			b.WriteString(d.QuoteIdent(column.Name()))
			b.WriteString(" ")
		} else {
			b.WriteString("CHANGE COLUMN ")
			b.WriteString(d.QuoteIdent(column.Name()))
			b.WriteString(" ")
			b.WriteString(d.QuoteIdent(column.Rename()))
			b.WriteString(" ")
		}

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(columnType)
		b.WriteString(d.ModifyColumn(column))
		b.WriteString(";")

		break
	}

	if b.String() == "" {
		return "", errors.New(fmt.Sprintf("dbx: not found changed column '%s'", columnName))
	}

	return b.String(), nil
}

// Compile add a column.
func (d mysql) CompileAddColumns(table Table) (string, error) {
	var b strings.Builder
	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" \n")

	columns := table.AddedColumns()

	for i, column := range columns {
		b.WriteString("ADD COLUMN ")
		b.WriteString(d.QuoteIdent(column.Name()))

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(" ")
		b.WriteString(columnType)
		b.WriteString(d.ModifyColumn(column))

		if i < len(columns)-1 {
			b.WriteString(",\n")
		}
	}

	b.WriteString(";")

	return b.String(), nil
}

// Compile add a column.
func (d mysql) CompileAddColumn(table Table, columnName string) (string, error) {
	var b strings.Builder

	columns := table.AddedColumns()

	for _, column := range columns {

		if column.Name() != columnName {
			continue
		}

		b.WriteString("ALTER TABLE ")
		b.WriteString(wrapTableName(table, d.QuoteIdent))
		b.WriteString(" ")

		b.WriteString("ADD COLUMN ")
		b.WriteString(d.QuoteIdent(column.Name()))

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(" ")
		b.WriteString(columnType)
		b.WriteString(d.ModifyColumn(column))
		b.WriteString(";")

		break
	}

	if b.String() == "" {
		return "", errors.New(fmt.Sprintf("dbx: not found added column '%s'", columnName))
	}

	return b.String(), nil
}

// Compile a primary key command.
func (d mysql) CompilePrimaryKey(table Table, columnNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP PRIMARY KEY,")
	b.WriteString(" ADD PRIMARY KEY (")
	b.WriteString(utils.QuoteIdents(columnNames, d.QuoteIdent))
	b.WriteString(");")

	return b.String(), nil
}

// Compile a drop primary key command.
func (d mysql) CompileDropPrimaryKey(table Table) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP PRIMARY KEY;")

	return b.String(), nil
}

// Compile an index creation command.
func (d mysql) CompileIndex(table Table, index Index) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" ADD")

	switch index.Type() {
	case NormalIndex:
		b.WriteString(" INDEX ")
	case UniqueIndex:
		b.WriteString(" UNIQUE ")
	case SpatialIndex, GistIndex:
		b.WriteString(" SPATIAL INDEX ")
	case FullTextIndex:
		b.WriteString(" FULLTEXT INDEX ")
	}

	b.WriteString(d.QuoteIdent(index.Name()))

	if index.Algorithm() != "" {
		b.WriteString(" USING ")
		b.WriteString(index.Algorithm())
	}

	b.WriteString("(")
	b.WriteString(utils.QuoteIdents(index.ColumnNames(), d.QuoteIdent))
	b.WriteString(");")

	return b.String(), nil
}

// Compile a drop index command.
func (d mysql) CompileDropIndex(table Table, indexName string) (string, error) {
	var b strings.Builder
	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP INDEX ")
	b.WriteString(d.QuoteIdent(indexName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop unique index command.
func (d mysql) CompileDropUnique(table Table, indexName string) (string, error) {
	return d.CompileDropIndex(table, indexName)
}

// Compile a drop spatial index command.
func (d mysql) CompileDropSpatialIndex(table Table, indexName string) (string, error) {
	return d.CompileDropIndex(table, indexName)
}

// Compile a drop foreign index command.
func (d mysql) CompileDropForeign(table Table, indexName string) (string, error) {
	var b strings.Builder
	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP FOREIGN KEY ")
	b.WriteString(d.QuoteIdent(indexName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop table command.
func (d mysql) CompileDrop(tableName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP TABLE ")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop table (if exists) command.
func (d mysql) CompileDropIfExists(tableName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP TABLE IF EXISTS")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop column command.
func (d mysql) CompileDropColumn(table Table, columnNames []string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))

	if len(columnNames) > 1 {
		b.WriteString(" \n")
	} else {
		b.WriteString(" ")
	}

	for i, name := range columnNames {

		if len(columnNames) > 1 {
			b.WriteString("\t")
		}

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
func (d mysql) CompileRenameTable(table Table, toName string) (string, error) {
	var b strings.Builder

	b.WriteString("RENAME TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" TO ")
	b.WriteString(d.QuoteIdent(toName))
	b.WriteString(";")

	return b.String(), nil
}

// Compile a rename index command.
func (d mysql) CompileRenameIndex(table Table, from string, to string) (string, error) {
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
func (d mysql) CompileDropAllTables(tableNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("DROP TABLE ")

	for i, name := range tableNames {
		tableNames[i] = d.QuoteIdent(name)
	}

	b.WriteString(strings.Join(tableNames, ","))
	b.WriteString(";")

	return b.String(), nil
}

// Compile the SQL needed to drop all views.
func (d mysql) CompileDropAllViews(viewNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("DROP VIEW ")

	for i, name := range viewNames {
		viewNames[i] = d.QuoteIdent(name)
	}

	b.WriteString(strings.Join(viewNames, ","))
	b.WriteString(";")

	return b.String(), nil
}

// Compile the SQL needed to drop all types.
func (d mysql) CompileDropAllTypes(typeNames ...string) (string, error) {
	return "", errors.New("mysql not support CompileDropAllTypes")
}

// Compile the SQL needed to retrieve all table names.
func (d mysql) CompileGetAllTables(schemaNames ...string) (string, error) {
	return `SHOW FULL TABLES WHERE table_type = 'BASE TABLE';`, nil
}

// Compile the SQL needed to retrieve all view names.
func (d mysql) CompileGetAllViews(schemaNames ...string) (string, error) {
	return `SHOW FULL TABLES WHERE table_type = 'VIEW';`, nil
}

// Compile the SQL needed to retrieve all type names.
func (d mysql) CompileGetAllTypes() (string, error) {
	return "", errors.New("mysql not support 'CompileGetAllTypes'")
}

// Compile the SQL needed to rebuild the database. [SQLite]
func (d mysql) CompileRebuild() (string, error) {
	return "", errors.New("mysql not support 'CompileRebuild'")
}

// Compile the command to enable foreign key constraints.
func (d mysql) CompileEnableForeignKeyConstraints() (string, error) {
	return `SET FOREIGN_KEY_CHECKS=1;`, nil
}

// Compile the command to disable foreign key constraints.
func (d mysql) CompileDisableForeignKeyConstraints() (string, error) {
	return `SET FOREIGN_KEY_CHECKS=0;`, nil
}

// Load table columns from the database
func (d mysql) LoadColumns(tableName string, tableSchema ...string) ([]Column, error) {

	var sql = "select * from information_schema.columns where table_name = ?"

	var columnInfos []mysqlColumnInfo

	_, err := d.query(sql, &columnInfos, tableName)

	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("error on query columns of the table %s", utils.QuoteStrings(tableName)))
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

		column.nullable = info.Nullable == "YES"

		if info.Length.Valid {
			column.length = int(info.Length.Int64)
		}

		if info.Precision.Valid {
			column.precision = int(info.Precision.Int64)
		}

		if info.Scale.Valid {
			column.scale = int(info.Scale.Int64)
		}

		if info.Default.Valid {
			column.defaultValue = info.Default.String
		}

		if info.Comment.Valid {
			column.comment = info.Comment.String
		}

		if info.Charset.Valid {
			column.charset = info.Charset.String
		}

		if info.Collation.Valid {
			column.collate = info.Collation.String
		}

		if info.Key.Valid && info.Key.String == "PRI" {
			column.primaryKey = true
		}

		if info.Extra.Valid {
			if info.Extra.String == "auto_increment" {
				column.autoIncrement = true
			} else if strings.Contains(info.Extra.String, "CURRENT_TIMESTAMP") {
				column.useCurrent = true
			}
		}

		if strings.Contains(info.ColumnType, "unsigned") {
			column.unsigned = true
		}

		if column.DataType() == TypeEnum {
			columnTypeParams := info.ColumnType[strings.Index(info.ColumnType, "(")+1 : strings.LastIndex(info.ColumnType, ")")]
			params := strings.Split(columnTypeParams, ",")

			var values []interface{}

			for _, param := range params {
				values = append(values, strings.Trim(param, "'"))
			}

			column.allowedValues = values
		}

		// save old column
		column.store()

		columns = append(columns, column)
	}

	return columns, nil
}

func (d mysql) LoadTable(tableName string, tableSchema ...string) (Table, error) {
	var sql strings.Builder

	sql.WriteString("select * from information_schema.tables where table_name = ?")

	var args = []interface{}{tableName}

	if len(tableSchema) > 0 {
		sql.WriteString(" and table_schema = ?")
		args = append(args, tableSchema[0])
	}

	var info mysqlTableInfo

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

	if info.Engine.Valid {
		table.SetEngine(info.Engine.String)
	}
	if info.Collation.Valid {
		table.SetCollation(info.Collation.String)
	}
	if info.Options.Valid {
		table.SetOptions(info.Options.String)
	}
	if info.Comment.Valid {
		table.SetComment(info.Comment.String)
	}

	// load from databse
	table.added = false

	return table, nil
}

// Load table constraints from the database
func (d mysql) LoadTableConstraints(tableName string, tableSchema ...string) (TableConstraints, error) {
	var tableConstraints TableConstraints

	var sqlStr = `
SELECT
    kcu.CONSTRAINT_NAME AS name,
    kcu.COLUMN_NAME AS column_name,
    tc.CONSTRAINT_TYPE AS type,
    CASE
        WHEN :schemaName IS NULL AND kcu.REFERENCED_TABLE_SCHEMA = DATABASE() THEN NULL
        ELSE kcu.REFERENCED_TABLE_SCHEMA
    END AS foreign_table_schema,
    kcu.REFERENCED_TABLE_NAME AS foreign_table_name,
    kcu.REFERENCED_COLUMN_NAME AS foreign_column_name,
    rc.UPDATE_RULE AS on_update,
    rc.DELETE_RULE AS on_delete,
    kcu.ORDINAL_POSITION AS position
FROM
    information_schema.KEY_COLUMN_USAGE AS kcu,
    information_schema.REFERENTIAL_CONSTRAINTS AS rc,
    information_schema.TABLE_CONSTRAINTS AS tc
WHERE
    kcu.TABLE_SCHEMA = COALESCE(:schemaName, DATABASE()) AND kcu.CONSTRAINT_SCHEMA = kcu.TABLE_SCHEMA AND kcu.TABLE_NAME = :tableName
    AND rc.CONSTRAINT_SCHEMA = kcu.TABLE_SCHEMA AND rc.TABLE_NAME = :tableName AND rc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME
    AND tc.TABLE_SCHEMA = kcu.TABLE_SCHEMA AND tc.TABLE_NAME = :tableName AND tc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME AND tc.CONSTRAINT_TYPE = 'FOREIGN KEY'
UNION
SELECT
    kcu.CONSTRAINT_NAME AS name,
    kcu.COLUMN_NAME AS column_name,
    tc.CONSTRAINT_TYPE AS type,
    NULL AS foreign_table_schema,
    NULL AS foreign_table_name,
    NULL AS foreign_column_name,
    NULL AS on_update,
    NULL AS on_delete,
    kcu.ORDINAL_POSITION AS position
FROM
    information_schema.KEY_COLUMN_USAGE AS kcu,
    information_schema.TABLE_CONSTRAINTS AS tc
WHERE
    kcu.TABLE_SCHEMA = COALESCE(:schemaName, DATABASE()) AND kcu.TABLE_NAME = :tableName
    AND tc.TABLE_SCHEMA = kcu.TABLE_SCHEMA AND tc.TABLE_NAME = :tableName AND tc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME AND tc.CONSTRAINT_TYPE IN ('PRIMARY KEY', 'UNIQUE')
ORDER BY position ASC`

	if len(tableSchema) > 0 {
		sqlStr = strings.ReplaceAll(sqlStr, ":schemaName", tableSchema[0])
	} else {
		sqlStr = strings.ReplaceAll(sqlStr, ":schemaName", "NULL")
	}

	sqlStr = strings.ReplaceAll(sqlStr, ":tableName", "?")

	args := []interface{}{
		tableName,
		tableName,
		tableName,
		tableName,
		tableName,
	}

	var constraints TableConstraintInfos

	_, err := d.query(sqlStr, &constraints, args...)

	if err != nil {
		return tableConstraints, err
	}

	for typ, nameGroup := range constraints.Group() {
		for name, items := range nameGroup {
			switch typ {
			case "PRIMARY KEY":
				tableConstraints.SetPrimaryKey(constraint.PrimaryKey{
					ColumnNames: items.ColumnNames(),
				})
			case "FOREIGN KEY":
				tableConstraints.AddForeignKeys(constraint.ForeignKey{
					Name:               name,
					ColumnNames:        items.ColumnNames(),
					ForeignSchemaName:  items[0].ForeignTableSchema.String,
					ForeignTableName:   items[0].ForeignTableName.String,
					ForeignColumnNames: items.ForeignColumnNames(),
					OnDelete:           items[0].OnDelete.String,
					OnUpdate:           items[0].OnUpdate.String,
				})
			case "UNIQUE":
				tableConstraints.AddUniques(constraint.Unique{
					Name:        name,
					ColumnNames: items.ColumnNames(),
				})
			}
		}
	}

	return tableConstraints, err
}
