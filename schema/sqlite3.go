package schema

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/gokit/dbx/schema/constraint"
	"github.com/gokit/dbx/utils"
	"regexp"
	"strings"
)

type sqlite3 struct {
	query Query
}

func (d sqlite3) QuoteIdent(s string) string {
	return utils.QuoteIdent(s, `"`)
}

// wrap index name
func (d sqlite3) wrapIndexName(table Table, indexName string) string {
	var b strings.Builder

	if table.Schema() != "" {
		b.WriteString(d.QuoteIdent(table.Schema()))
		b.WriteString(".")
	}

	b.WriteString(d.QuoteIdent(indexName))

	return b.String()
}

// Create the column definition for a char type.
func (d sqlite3) TypeChar(column Column) (string, error) {
	return "varchar", nil
}

// Create the column definition for a string type.
func (d sqlite3) TypeString(column Column) (string, error) {
	return "varchar", nil
}

// Create the column definition for a text type.
func (d sqlite3) TypeText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a medium text type.
func (d sqlite3) TypeMediumText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a long text type.
func (d sqlite3) TypeLongText(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for an integer type.
func (d sqlite3) TypeInteger(column Column) (string, error) {
	return "integer", nil
}

// Create the column definition for a big integer type.
func (d sqlite3) TypeBigInteger(column Column) (string, error) {
	return "integer", nil
}

// Create the column definition for a medium integer type.
func (d sqlite3) TypeMediumInteger(column Column) (string, error) {
	return "integer", nil
}

// Create the column definition for a tiny integer type.
func (d sqlite3) TypeTinyInteger(column Column) (string, error) {
	return "integer", nil
}

// Create the column definition for a small integer type.
func (d sqlite3) TypeSmallInteger(column Column) (string, error) {
	return "integer", nil
}

// Create the column definition for a tiny blob type.
func (d sqlite3) TypeTinyBlob(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for an blob type.
func (d sqlite3) TypeBlob(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for a medium blob type.
func (d sqlite3) TypeMediumBlob(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for a long blob type.
func (d sqlite3) TypeLongBlob(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for a float type.
func (d sqlite3) TypeFloat(column Column) (string, error) {
	return "float", nil
}

// Create the column definition for a double type.
func (d sqlite3) TypeDouble(column Column) (string, error) {
	return "float", nil
}

// Create the column definition for a decimal type.
func (d sqlite3) TypeDecimal(column Column) (string, error) {
	return "numeric", nil
}

// Create the column definition for a boolean type.
func (d sqlite3) TypeBoolean(column Column) (string, error) {
	return "tinyint(1)", nil
}

// Create the column definition for an enumeration type.
func (d sqlite3) TypeEnum(column Column) (string, error) {
	return fmt.Sprintf(
		"varchar check (%s in (%s))",
		d.QuoteIdent(column.Name()),
		utils.QuoteInterfaces(column.AllowedValues()...),
	), nil
}

// Create the column definition for an enumeration type.
func (d sqlite3) TypeSet(column Column) (string, error) {
	return "", errors.New("This database driver not support type 'set'.")
}

// Create the column definition for a json type.
func (d sqlite3) TypeJson(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a jsonb type.
func (d sqlite3) TypeJsonb(column Column) (string, error) {
	return "text", nil
}

// Create the column definition for a date type.
func (d sqlite3) TypeDate(column Column) (string, error) {
	return "date", nil
}

// Create the column definition for a date-time type.
func (d sqlite3) TypeDateTime(column Column) (string, error) {
	return d.TypeTimestamp(column)
}

// Create the column definition for a date-time (with time zone) type.
func (d sqlite3) TypeDateTimeTz(column Column) (string, error) {
	return d.TypeDateTime(column)
}

// Create the column definition for a time type.
func (d sqlite3) TypeTime(column Column) (string, error) {
	return "time", nil
}

// Create the column definition for a time (with time zone) type.
func (d sqlite3) TypeTimeTz(column Column) (string, error) {
	return d.TypeTime(column)
}

// Create the column definition for a timestamp type.
func (d sqlite3) TypeTimestamp(column Column) (string, error) {
	return "datetime", nil
}

// Create the column definition for a timestamp (with time zone) type.
func (d sqlite3) TypeTimestampTz(column Column) (string, error) {
	return d.TypeTimestamp(column)
}

// Create the column definition for a year type.
func (d sqlite3) TypeYear(column Column) (string, error) {
	return d.TypeInteger(column)
}

// Create the column definition for a binary type.
func (d sqlite3) TypeBinary(column Column) (string, error) {
	return "blob", nil
}

// Create the column definition for a uuid type.
func (d sqlite3) TypeUuid(column Column) (string, error) {
	return "varchar", nil
}

// Create the column definition for an IP address type.
func (d sqlite3) TypeIpAddress(column Column) (string, error) {
	return "varchar", nil
}

// Create the column definition for a MAC address type.
func (d sqlite3) TypeMacAddress(column Column) (string, error) {
	return "varchar", nil
}

// Create the column definition for a spatial Geometry type.
func (d sqlite3) TypeGeometry(column Column) (string, error) {
	return "geometry", nil
}

// Create the column definition for a spatial Point type.
func (d sqlite3) TypePoint(column Column) (string, error) {
	return "point", nil
}

// Create the column definition for a spatial LineString type.
func (d sqlite3) TypeLineString(column Column) (string, error) {
	return "linestring", nil
}

// Create the column definition for a spatial Polygon type.
func (d sqlite3) TypePolygon(column Column) (string, error) {
	return "polygon", nil
}

// Create the column definition for a spatial GeometryCollection type.
func (d sqlite3) TypeGeometryCollection(column Column) (string, error) {
	return "geometrycollection", nil
}

// Create the column definition for a spatial MultiPoint type.
func (d sqlite3) TypeMultiPoint(column Column) (string, error) {
	return "multipoint", nil
}

// Create the column definition for a spatial MultiLineString type.
func (d sqlite3) TypeMultiLineString(column Column) (string, error) {
	return "multilinestring", nil
}

// Create the column definition for a spatial MultiPolygon type.
func (d sqlite3) TypeMultiPolygon(column Column) (string, error) {
	return "multipolygon", nil
}

func (d sqlite3) ModifyColumn(column Column) string {
	// 'Nullable', 'Default', 'Increment'

	var b strings.Builder

	if column.Nullable() {
		b.WriteString(" NULL")
	} else {
		b.WriteString(" NOT NULL")
	}

	if column.DefaultValue() != nil {
		value := utils.DefaultValue(column.DefaultValue())

		if value != "" {
			b.WriteString(" DEFAULT ")
			b.WriteString(utils.DefaultValue(column.DefaultValue()))
		}
	}

	// AutoIncrement must be primary key
	if column.AutoIncrement() {
		switch column.DataType() {
		case TypeInt, TypeBigInt, TypeMediumInt, TypeTinyInt, TypeSmallInt:
			b.WriteString(" PRIMARY KEY AUTOINCREMENT")
		}
	}

	return b.String()
}

// Compile the query to determine the list of tables.
func (d sqlite3) CompileTableExists(tableName string, tableSchema ...string) (string, error) {
	var b strings.Builder

	b.WriteString("select * from sqlite_master where type = 'table' and name = ")
	b.WriteString(utils.QuoteStrings(tableName))

	return b.String(), nil
}

// Compile the query to determine the list of columns.
func (d sqlite3) CompileColumnListing(tableName string, tableSchema ...string) (string, error) {
	var b strings.Builder

	b.WriteString("PRAGMA TABLE_INFO (")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(")")

	return b.String(), nil
}

//  Compile a create table command.
func (d sqlite3) CompileCreate(table Table) (string, error) {
	var b strings.Builder

	if table.IsTemporary() {
		b.WriteString("CREATE TEMPORARY TABLE ")
	} else {
		b.WriteString("CREATE TABLE ")
	}
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString("(\n")

	// primary key names
	// sqlite 联合主键不支持自增
	var pkNames []string

	columns := table.Columns()

	hasAutoIncrement := false

	for _, column := range columns {
		if column.PrimaryKey() {
			pkNames = append(pkNames, d.QuoteIdent(column.Name()))
		}

		if column.AutoIncrement() {
			hasAutoIncrement = true
		}
	}

	// auto increment must be primary key for sqlite
	appendPrimaryKey := len(pkNames) > 0 && (len(pkNames) == 1 && hasAutoIncrement) == false

	for i, column := range columns {

		b.WriteString("\t")
		b.WriteString(d.QuoteIdent(column.Name()))
		b.WriteString(" ")

		columnType, err := ColumnType(d, column)

		if err != nil {
			return "", err
		}

		b.WriteString(columnType)
		b.WriteString(d.ModifyColumn(column))

		if i < len(columns)-1 || appendPrimaryKey {
			b.WriteString(",")
		}

		b.WriteString("\n")
	}

	if appendPrimaryKey {
		b.WriteString("\tPRIMARY KEY (")
		b.WriteString(strings.Join(pkNames, ", "))
		b.WriteString(")\n")
	}

	b.WriteString(")")

	b.WriteString(";")

	return b.String(), nil
}

//  Compile a create table command.
func (d sqlite3) CompileModifyColumns(table Table) (string, error) {
	panic("implement me")
}

//  Compile a modify column command.
func (d sqlite3) CompileModifyColumn(table Table, columnName string) (string, error) {
	panic("implement me")
}

// Compile add columns.
func (d sqlite3) CompileAddColumns(table Table) (string, error) {
	panic("implement me")
}

// Compile add a column.
func (d sqlite3) CompileAddColumn(table Table, columnName string) (string, error) {
	panic("implement me")
}

// Compile a primary key command.
func (d sqlite3) CompilePrimaryKey(table Table, columnNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" ADD PRIMARY KEY (")
	b.WriteString(utils.QuoteStrings(columnNames...))
	b.WriteString(");")

	return b.String(), nil
}

// Compile a drop primary key command.
func (d sqlite3) CompileDropPrimaryKey(table Table) (string, error) {
	var b strings.Builder

	b.WriteString("ALTER TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" DROP PRIMARY KEY;")

	return b.String(), nil
}

// Compile an index creation command.
func (d sqlite3) CompileIndex(table Table, index Index) (string, error) {
	var b strings.Builder

	b.WriteString("CREATE")

	switch index.Type() {
	case NormalIndex:
		b.WriteString(" UNIQUE ")
	case UniqueIndex:
		b.WriteString(" INDEX ")
	case SpatialIndex, GistIndex:
		return "", errors.New("The database driver in use does not support spatial indexes.")
	case FullTextIndex:
		return "", errors.New("The database driver in use does not support fulltext indexes.")
	default:
		return "", errors.New("The database driver in use does not support " + strings.ToLower(string(index.Type())) + " indexes.")
	}

	b.WriteString(d.QuoteIdent(index.Name()))
	b.WriteString(" ON ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))

	b.WriteString(d.QuoteIdent(index.Name()))

	b.WriteString("(")
	b.WriteString(utils.QuoteIdents(index.ColumnNames(), d.QuoteIdent))
	b.WriteString(");")

	return b.String(), nil
}

// Compile a drop index command.
func (d sqlite3) CompileDropIndex(table Table, indexName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP INDEX ")
	b.WriteString(d.wrapIndexName(table, indexName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop unique index command.
func (d sqlite3) CompileDropUnique(table Table, indexName string) (string, error) {
	return d.CompileDropIndex(table, indexName)
}

// Compile a drop spatial index command.
func (d sqlite3) CompileDropSpatialIndex(table Table, indexName string) (string, error) {
	return "", errors.New(fmt.Sprintf("dbx: The database driver in use does not support spatial indexes"))
}

// Compile a drop foreign index command.
func (d sqlite3) CompileDropForeign(table Table, indexName string) (string, error) {
	panic("implement me")
}

// Compile a drop table command.
func (d sqlite3) CompileDrop(tableName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP TABLE ")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop table (if exists) command.
func (d sqlite3) CompileDropIfExists(tableName string) (string, error) {
	var b strings.Builder
	b.WriteString("DROP TABLE IF EXISTS")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(";")
	return b.String(), nil
}

// Compile a drop column command.
func (d sqlite3) CompileDropColumn(table Table, columnNames []string) (string, error) {
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
func (d sqlite3) CompileRenameTable(table Table, toName string) (string, error) {
	var b strings.Builder

	b.WriteString("RENAME TABLE ")
	b.WriteString(wrapTableName(table, d.QuoteIdent))
	b.WriteString(" TO ")
	b.WriteString(d.QuoteIdent(toName))
	b.WriteString(";")

	return b.String(), nil
}

// Compile a rename index command.
func (d sqlite3) CompileRenameIndex(table Table, from string, to string) (string, error) {
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
func (d sqlite3) CompileDropAllTables(tableNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("delete from sqlite_master where type in ('table', 'index', 'trigger')")

	if len(tableNames) > 0 {
		b.WriteString(" and name in (")

		for i, name := range tableNames {
			tableNames[i] = d.QuoteIdent(name)
		}

		b.WriteString(strings.Join(tableNames, ","))
		b.WriteString(")")
	}

	b.WriteString(";")

	return b.String(), nil
}

// Compile the SQL needed to drop all views.
func (d sqlite3) CompileDropAllViews(viewNames ...string) (string, error) {
	var b strings.Builder

	b.WriteString("delete from sqlite_master where type in ('view')")

	for i, name := range viewNames {
		viewNames[i] = d.QuoteIdent(name)
	}

	if len(viewNames) > 0 {
		b.WriteString(" and name in (")

		for i, name := range viewNames {
			viewNames[i] = d.QuoteIdent(name)
		}

		b.WriteString(strings.Join(viewNames, ","))
		b.WriteString(")")
	}

	b.WriteString(";")

	return b.String(), nil
}

// Compile the SQL needed to drop all types. [PostgreSQL]
func (d sqlite3) CompileDropAllTypes(typeNames ...string) (string, error) {
	return "", errors.New("sqlite3 not support 'CompileDropAllTypes'")
}

// Compile the SQL needed to retrieve all table names.
func (d sqlite3) CompileGetAllTables(schemaNames ...string) (string, error) {
	return `SHOW FULL TABLES WHERE table_type = 'BASE TABLE';`, nil
}

// Compile the SQL needed to retrieve all view names.
func (d sqlite3) CompileGetAllViews(schemaNames ...string) (string, error) {
	return `SHOW FULL TABLES WHERE table_type = 'VIEW';`, nil
}

// Compile the SQL needed to rebuild the database. [SQLite]
func (d sqlite3) CompileRebuild() (string, error) {
	return "vacuum", nil
}

// Compile the SQL needed to retrieve all type names. [PostgreSQL]
func (d sqlite3) CompileGetAllTypes() (string, error) {
	return "", errors.New("sqlite3 not support 'CompileGetAllTypes'")
}

// Compile the command to enable foreign key constraints.
func (d sqlite3) CompileEnableForeignKeyConstraints() (string, error) {
	return `PRAGMA foreign_keys = ON;`, nil
}

// Compile the command to disable foreign key constraints.
func (d sqlite3) CompileDisableForeignKeyConstraints() (string, error) {
	return `PRAGMA foreign_keys = OFF;`, nil
}

type sqliteColumnInfo struct {
	Cid        int            `db:"cid"`
	Name       string         `db:"name"`
	DataType   string         `db:"type"`
	NotNull    bool           `db:"notnull"`
	Default    sql.NullString `db:"dflt_value"`
	PrimaryKey int            `db:"pk"`
}

// Load table columns
func (d sqlite3) LoadColumns(tableName string, tableSchema ...string) ([]Column, error) {

	var createTableSql string

	count, err := d.query("select sql from sqlite_master where type = ? and tbl_name = ?", &createTableSql, "table", tableName)

	if err != nil {
		return nil, errors.Wrapf(err, "dbx: error on find create table sql of the table '%s'", tableName)
	}

	if count == 0 {
		return nil, errors.New(fmt.Sprintf("dbx: not found table '%s'", tableName))
	}

	lines := strings.Split(createTableSql, "\n")
	lines = lines[1 : len(lines)-1]

	var columnMap = make(map[string]string)

	pattern := regexp.MustCompile("\"(.+)\"")

	for _, line := range lines {
		line = strings.Trim(line, " ,\t")

		if strings.HasPrefix(line, "\"") {
			name := pattern.FindString(line)
			name = name[1 : len(name)-1]
			columnMap[name] = line
		}
	}

	var b strings.Builder

	b.WriteString("PRAGMA TABLE_INFO (")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(")")

	var columnInfos []sqliteColumnInfo

	_, err = d.query(b.String(), &columnInfos)

	if err != nil {
		return nil, err
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

		column.nullable = info.NotNull == false

		// 丑陋的代码
		if info.Default.Valid {
			if info.Default.String == "CURRENT_TIMESTAMP" {
				column.defaultValue = utils.Express("CURRENT_TIMESTAMP")
				column.useCurrent = true
			} else {
				if strings.HasPrefix(info.Default.String, "'") && strings.HasSuffix(info.Default.String, "'") {
					info.Default.String = info.Default.String[1 : len(info.Default.String)-1]
				}
				column.defaultValue = info.Default.String
			}
		}

		if info.PrimaryKey > 0 {
			column.primaryKey = true
			// AUTOINCREMENT
			if expr, ok := columnMap[info.Name]; ok {
				if strings.Contains(expr, "PRIMARY KEY AUTOINCREMENT") {
					column.autoIncrement = true
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
func (d sqlite3) LoadTable(tableName string, tableSchema ...string) (Table, error) {

	var names []string

	_, err := d.query("select sql from sqlite_master where type = ? and tbl_name = ?", &names, "table", tableName)

	if err != nil {
		return nil, err
	}

	columns, err := d.LoadColumns(tableName)

	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("error on query columns of the table %s", utils.QuoteStrings(tableName)))
	}

	table := NewTable(tableName, columns...)

	return table, nil
}

// Returns table columns info
func (d sqlite3) loadTableColumnsInfo(tableName string) (map[int]sqliteColumnInfo, error) {
	var b strings.Builder

	b.WriteString("PRAGMA TABLE_INFO (")
	b.WriteString(d.QuoteIdent(tableName))
	b.WriteString(")")

	var columnInfos []sqliteColumnInfo

	_, err := d.query(b.String(), &columnInfos)

	if err != nil {
		return nil, err
	}

	var m = make(map[int]sqliteColumnInfo)

	for _, info := range columnInfos {
		m[info.Cid] = info
	}

	return m, nil
}

type sqliteContraintInfo struct {
	Seq     int    `db:"seq"`
	Name    string `db:"name"`
	Unique  bool   `db:"unique"`
	Origin  string `db:"origin"`
	Partial int    `db:"partial"`
}

type sqliteIndexInfo struct {
	Seqno int    `db:seqno`
	Cid   int    `db:cid`
	Name  string `db:name`
}

type sqliteIndexInfos []sqliteIndexInfo

func (s sqliteIndexInfos) ColumnNames() []string {
	var columnNames []string
	for _, info := range s {
		columnNames = append(columnNames, info.Name)
	}
	return columnNames
}

// Load table constraints from the database
func (d sqlite3) LoadTableConstraints(tableName string, tableSchema ...string) (TableConstraints, error) {
	var contraintInfos []sqliteContraintInfo

	var b strings.Builder

	b.WriteString("PRAGMA INDEX_LIST (")
	b.WriteString(utils.QuoteStrings(tableName))
	b.WriteString(")")

	_, err := d.query(b.String(), &contraintInfos)

	var constraints TableConstraints

	if err != nil {
		return constraints, err
	}

	var tableColumns = make(map[int]sqliteColumnInfo)

	if len(contraintInfos) > 0 && contraintInfos[0].Origin != "" {

		/*
		 * SQLite may not have an "origin" column in INDEX_LIST
		 * See https://www.sqlite.org/src/info/2743846cdba572f6
		 */
		tableColumns, err = d.loadTableColumnsInfo(tableName)

		if err != nil {
			return constraints, err
		}
	}

	for _, info := range contraintInfos {
		var columns sqliteIndexInfos

		b.Reset()
		b.WriteString("PRAGMA INDEX_INFO (")
		b.WriteString(utils.QuoteStrings(info.Name))
		b.WriteString(")")

		_, err := d.query(b.String(), &columns)

		if err != nil {
			return constraints, err
		}

		if len(tableColumns) > 0 {
			// SQLite may not have an "origin" column in INDEX_LIST
			info.Origin = "c"

			if len(columns) > 0 && tableColumns[columns[0].Cid].PrimaryKey > 0 {
				info.Origin = "pk"
			} else if info.Unique && d.isSystemIdentifier(info.Name) {
				info.Origin = "u"
			}
		}

		if info.Origin == "u" {
			constraints.AddUniques(constraint.Unique{
				Name:        info.Name,
				ColumnNames: columns.ColumnNames(),
			})
		} else if info.Origin == "pk" {
			constraints.SetPrimaryKey(constraint.PrimaryKey{
				Name:        info.Name,
				ColumnNames: columns.ColumnNames(),
			})
		}

		if constraints.HasPrimaryKey() == false {
			/*
			 * Additional check for PK in case of INTEGER PRIMARY KEY with ROWID
			 * See https://www.sqlite.org/lang_createtable.html#primkeyconst
			 */
			if len(tableColumns) == 0 {
				tableColumns, err = d.loadTableColumnsInfo(tableName)

				if err != nil {
					return constraints, err
				}
			}

			for _, tableColumn := range tableColumns {
				if tableColumn.PrimaryKey > 0 {
					constraints.SetPrimaryKey(constraint.PrimaryKey{
						ColumnNames: []string{tableColumn.Name},
					})
					break
				}
			}
		}
	}

	return constraints, err
}

func (d sqlite3) isSystemIdentifier(identifier string) bool {
	return strings.HasPrefix(identifier, "sqlite_")
}
