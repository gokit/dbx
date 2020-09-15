package schema

import (
	"fmt"
	"github.com/friendsofgo/errors"
	"strings"
)

type TableHandler func(table *TableSchema)

type TableCommand func(d Dialect) (string, error)

// Table schema implement Table interface
type TableSchema struct {
	name      string         // the name of the table
	schema    string         // the schema of the table
	temporary bool           //
	prefix    string         // the prefix of the table
	engine    string         // the engine of the table (MySQL)
	charset   string         // the charset of the table
	collation string         // The collation that should be used for the table
	comment   string         // the comment of the table
	options   string         //
	columns   []Column       //
	indexes   []Index        //
	commands  []TableCommand // the commands of the table
	added     bool           // load from database
}

// New table
func NewTable(name string, columns ...Column) *TableSchema {
	return &TableSchema{
		name:    name,
		columns: columns,
		added:   true,
	}
}

func (t *TableSchema) Name() string {
	return t.name
}

func (t *TableSchema) Schema() string {
	return t.schema
}

func (t *TableSchema) Prefix() string {
	return t.prefix
}

func (t *TableSchema) Engine() string {
	return t.engine
}

func (t *TableSchema) Charset() string {
	return t.charset
}

func (t *TableSchema) Collation() string {
	return t.collation
}

func (t *TableSchema) IsTemporary() bool {
	return t.temporary
}

func (t *TableSchema) Comment() string {
	return t.comment
}

// Columns get all columns of the table
func (t *TableSchema) Columns() []Column {
	return t.columns
}

// Commands get all command of the table
func (t *TableSchema) Commands() []TableCommand {
	return t.commands
}

// Get the primary keys of the table
func (t *TableSchema) PrimaryKey() []string {
	var keys []string

	for _, column := range t.Columns() {
		if column.PrimaryKey() {
			keys = append(keys, column.Name())
		}
	}

	return keys
}

// add table command
func (t *TableSchema) AddCommand(commands ...TableCommand) {
	t.commands = append(t.commands, commands...)
}

// Remove a column from columns of the table.
func (t *TableSchema) RemoveColumn(name string) *TableSchema {

	for i, column := range t.columns {
		if column.Name() == name || strings.ToLower(column.Name()) == strings.ToLower(name) {
			t.columns = append(t.columns[:i], t.columns[i+1:]...)
		}
	}

	return t
}

// Get the added columns
func (t *TableSchema) AddedColumns() []Column {

	var columns []Column

	for _, column := range t.columns {
		if column.Added() {
			columns = append(columns, column)
		}
	}

	return columns
}

// Get the changed columns
func (t *TableSchema) ChangedColumns() []Column {

	var columns []Column

	for _, column := range t.columns {
		if column.Changed() {
			columns = append(columns, column)
		}
	}

	return columns
}

// set the engine for the table
func (t *TableSchema) SetEngine(engine string) *TableSchema {
	t.engine = engine
	return t
}

// set the charset for the table
func (t *TableSchema) SetCharset(charset string) *TableSchema {
	t.charset = charset
	return t
}

// set the collation for the table
func (t *TableSchema) SetCollation(collation string) *TableSchema {
	t.collation = collation
	return t
}

// set the temporary for the table
func (t *TableSchema) SetTemporary(temporary bool) *TableSchema {
	t.temporary = temporary
	return t
}

// set the comment for the table
func (t *TableSchema) SetComment(comment string) *TableSchema {
	t.comment = comment
	return t
}

// set the prefix for the table
func (t *TableSchema) SetPrefix(prefix string) *TableSchema {
	t.prefix = prefix
	return t
}

// set the options for the table
func (t *TableSchema) SetOptions(options string) *TableSchema {
	t.options = options
	return t
}

// addColumn add a column with length
func (t *TableSchema) addColumn(name string, dataType DataType, options ...ColumnOption) *ColumnSchema {
	column := newColumn(name, dataType, options...)
	t.columns = append(t.columns, column)

	if t.added == false {
		t.AddCommand(func(d Dialect) (string, error) {
			return d.CompileAddColumn(t, name)
		})
	}

	return column
}

// ID Create a new auto-incrementing integer (4-byte) column on the table.
func (t *TableSchema) ID(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeInt, options...).With(PrimaryKey(), AutoIncrement(), Unsigned())
}

// BigID Create a new auto-incrementing big integer (8-byte) column on the table.
func (t *TableSchema) BigID(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeBigInt, options...).With(PrimaryKey(), AutoIncrement(), Unsigned())
}

// Char Create a new char column on the table.
func (t *TableSchema) Char(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeChar, options...)
}

// String Create a new string column on the table.
func (t *TableSchema) String(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeString, options...)
}

// Text Create a new text column on the table.
func (t *TableSchema) Text(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeText, options...)
}

// MediumText Create a new medium text column on the table.
func (t *TableSchema) MediumText(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMediumText)
}

// LongText Create a new long text column on the table.
func (t *TableSchema) LongText(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeLongText)
}

// Int Create a new integer (4-byte) column on the table.
func (t *TableSchema) Int(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeInt, options...)
}

// TinyInt Create a new tiny integer (1-byte) column on the table.
func (t *TableSchema) TinyInt(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeTinyInt, options...)
}

//  SmallInt Create a new small integer (2-byte) column on the table.
func (t *TableSchema) SmallInt(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeSmallInt, options...)
}

//  MediumInt Create a new medium integer (3-byte) column on the table.
func (t *TableSchema) MediumInt(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMediumInt, options...)
}

// BigInt Create a new big integer (8-byte) column on the table.
func (t *TableSchema) BigInt(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeBigInt, options...)
}

// Float Create a new float column on the table.
func (t *TableSchema) Float(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeFloat, options...)
}

// Double Create a new double column on the table.
func (t *TableSchema) Double(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeDouble, options...)
}

// Boolean Create a new decimal column on the table.
func (t *TableSchema) Decimal(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeDecimal, options...)
}

// Boolean Create a new boolean column on the table.
func (t *TableSchema) Boolean(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeBoolean, options...)
}

// Enum Create a new enum column on the table.
func (t *TableSchema) Enum(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeEnum, options...)
}

// Set Create a new set column on the table.
func (t *TableSchema) Set(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeSet, options...)
}

// Json Create a new json column on the table.
func (t *TableSchema) Json(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeJson, options...)
}

// Jsonb Create a new jsonb column on the table.
func (t *TableSchema) Jsonb(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeJsonb, options...)
}

// Date Create a new date column on the table.
func (t *TableSchema) Date(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeDate)
}

// DateTime Create a new dateTime column on the table.
func (t *TableSchema) DateTime(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeDateTime, options...)
}

// DateTimeTz Create a new dateTimeTz column on the table.
func (t *TableSchema) DateTimeTz(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeDateTimeTz, options...)
}

// Time Create a new time column on the table.
func (t *TableSchema) Time(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeTime, options...)
}

// TimeTz Create a new time column (with time zone) on the table.
func (t *TableSchema) TimeTz(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeTimeTz, options...)
}

// Timestamp Create a new timestamp column on the table.
func (t *TableSchema) Timestamp(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeTimestamp, options...)
}

// TimestampTz Create a new timestampTz column (with time zone) on the table.
func (t *TableSchema) TimestampTz(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeTimestampTz, options...)
}

// Year Create a new year column on the table.
func (t *TableSchema) Year(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeYear, options...)
}

// Binary Create a new binary column on the table.
func (t *TableSchema) Binary(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeBinary, options...)
}

// UUID Create a new uuid column on the table.
func (t *TableSchema) UUID(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeUUID)
}

// IPAddress Create a new IP address column on the table.
func (t *TableSchema) IPAddress(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeIpAddress)
}

// MacAddress create a new MAC address column on the table.
func (t *TableSchema) MacAddress(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMacAddress, options...).With(Length(17))
}

// Geometry create a new geometry column on the table.
func (t *TableSchema) Geometry(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeGeometry, options...)
}

// Point create a new point column on the table.
func (t *TableSchema) Point(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypePoint, options...)
}

// LineString create a new linestring column on the table.
func (t *TableSchema) LineString(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeLineString, options...)
}

// Polygon create a new polygon column on the table.
func (t *TableSchema) Polygon(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypePolygon, options...)
}

// Geometrycollection create a new geometry collection column on the table.
func (t *TableSchema) GeometryCollection(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeGeometryCollection, options...)
}

// MultiPoint create a new multipoint column on the table.
func (t *TableSchema) MultiPoint(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMultiPoint, options...)
}

// MultiLineString create a new multilinestring column on the table.
func (t *TableSchema) MultiLineString(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMultiLineString, options...)
}

// MultiPolygon create a new multipolygon column on the table.
func (t *TableSchema) MultiPolygon(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMultiPolygon, options...)
}

// MultiPolygonZ create a new multipolygonz column on the table.
func (t *TableSchema) MultiPolygonZ(name string, options ...ColumnOption) *ColumnSchema {
	return t.addColumn(name, TypeMultiPolygonZ, options...)
}

// whether the column exists of the table
func (t *TableSchema) ColumnExists(columnName string) bool {
	for _, column := range t.columns {
		if column.Name() == columnName {
			return true
		}
	}
	return false
}

// Rename the table
func (t *TableSchema) Rename(name string) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileRenameTable(t, name)
	})
}

// Drop the table
func (t *TableSchema) Drop(name string) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileDrop(name)
	})
}

// Drop the table if exists
func (t *TableSchema) DropIfExists(name string) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileDropIfExists(name)
	})
}

// Modify the column of the table
func (t *TableSchema) Modify(columnName string, options ...ColumnOption) {
	t.AddCommand(func(d Dialect) (string, error) {
		var column *ColumnSchema

		for _, col := range t.columns {
			if col.Name() == columnName {
				column = col.(*ColumnSchema)
			}
		}

		if column == nil {
			return "", errors.New(fmt.Sprintf("not found column '%s'", columnName))
		}

		// set changed
		column.changed = true

		return d.CompileModifyColumn(t, columnName)
	})
}

// Modify the column of the table
func (t *TableSchema) DropColumns(columnNames ...string) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileDropColumn(t, columnNames)
	})
}

// Add primary keys to the table
func (t *TableSchema) AddPrimaryKey(columnNames ...string) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompilePrimaryKey(t, columnNames...)
	})
}

// Drop the primary keys of the table
func (t *TableSchema) DropPrimaryKey() {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileDropPrimaryKey(t)
	})
}

// Add a new index to the table.
func (t *TableSchema) Index(columnNames []string, options ...IndexOptions) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileIndex(t, newIndex(NormalIndex, columnNames, options...))
	})
}

// Add a new unique index to the table.
func (t *TableSchema) Unique(columnNames []string, options ...IndexOptions) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileIndex(t, newIndex(UniqueIndex, columnNames, options...))
	})
}

// Add a new spatial index to the table.
func (t *TableSchema) SpatialIndex(columnNames []string, options ...IndexOptions) {
	t.AddCommand(func(d Dialect) (string, error) {
		return d.CompileIndex(t, newIndex(SpatialIndex, columnNames, options...))
	})
}
