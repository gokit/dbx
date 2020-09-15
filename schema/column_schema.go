package schema

// the column schema
type ColumnSchema struct {
	name              string        // name of this column (without quotes).
	rename            string        // change the name of this column
	dataType          DataType      // the abstract DB type of this column. Possible DB types vary according to the type of DBMS.
	defaultValue      interface{}   // Specify a "default" value for the column
	nullable          bool          // Allow NULL values to be inserted into the column
	allowedValues     []interface{} // set column allowed values when column type is enum or set
	length            int           // display size of the column.
	precision         int           // precision of the column data, if it is numeric.
	scale             int           // scale of the column data, if it is numeric
	primaryKey        bool          // is a primary index
	autoIncrement     bool          // Set INTEGER columns as auto-increment (primary key)
	unsigned          bool          // whether this column is unsigned. This is only meaningful when [[type]] is `smallint`, `integer` or `bigint`.
	generateAs        string        // Create a SQL compliant identity column (PostgreSQL)
	always            bool          // Place the column "after" another column (MySQL)
	charset           string        // Specify a character set for the column (MySQL)
	collate           string        // Specify a collate set for the column (MySQL)
	comment           string        // Add a comment to the column (MySQL)
	after             string        // Place the column "after" another column (MySQL)
	first             string        // Place the column "first" in the table (MySQL)
	append            string        // append
	added             bool          // Added the column
	changed           bool          // Change the column
	storedAs          string        // Create a stored generated column (MySQL)
	useCurrent        bool          // Set the TIMESTAMP column to use CURRENT_TIMESTAMP as default value
	virtualAs         string        // Create a virtual generated column (MySQL)
	projection        string        //
	srid              int           //
	hasNullable       bool          //
	hasLength         bool          //
	hasPrecisionScale bool          //
	hasAllowedValues  bool          //
	oldColumn         *ColumnSchema //
}

func (c *ColumnSchema) Name() string {
	return c.name
}

func (c *ColumnSchema) Rename() string {
	if c.rename != c.name {
		return c.rename
	}
	return ""
}

func (c *ColumnSchema) DataType() DataType {
	return c.dataType
}

func (c *ColumnSchema) DefaultValue() interface{} {
	return c.defaultValue
}

func (c *ColumnSchema) Nullable() bool {
	return c.nullable
}

func (c *ColumnSchema) AllowedValues() []interface{} {
	return c.allowedValues
}

func (c *ColumnSchema) Length() int {
	return c.length
}

func (c *ColumnSchema) Precision() int {
	return c.precision
}

func (c *ColumnSchema) Scale() int {
	return c.scale
}

func (c *ColumnSchema) PrimaryKey() bool {
	return c.primaryKey
}

func (c *ColumnSchema) AutoIncrement() bool {
	return c.autoIncrement
}

func (c *ColumnSchema) Unsigned() bool {
	return c.unsigned
}

func (c *ColumnSchema) GeneratedAs() string {
	return c.generateAs
}

func (c *ColumnSchema) Alawys() bool {
	return c.always
}

func (c *ColumnSchema) Charset() string {
	return c.charset
}

func (c *ColumnSchema) After() string {
	return c.after
}

func (c *ColumnSchema) First() string {
	return c.first
}

func (c *ColumnSchema) Append() string {
	return c.append
}

func (c *ColumnSchema) Added() bool {
	return c.added
}

func (c *ColumnSchema) Changed() bool {
	return c.changed
}

func (c *ColumnSchema) StoredAs() string {
	return c.storedAs
}

func (c *ColumnSchema) UseCurrent() bool {
	return c.useCurrent
}

func (c *ColumnSchema) VirtualAs() string {
	return c.virtualAs
}

func (c *ColumnSchema) Projection() string {
	return c.projection
}

func (c *ColumnSchema) Collate() string {
	return c.collate
}

func (c *ColumnSchema) Srid() int {
	return c.srid
}

func (c *ColumnSchema) Comment() string {
	return c.comment
}

// New a column
func newColumn(name string, dataType DataType, options ...ColumnOption) *ColumnSchema {

	column := &ColumnSchema{
		name:          name,
		dataType:      dataType,
		nullable:      false,
		primaryKey:    false,
		autoIncrement: false,
		unsigned:      false,
		added:         true,
		changed:       false,
	}

	column.With(options...)

	return column
}

// set options for column
func (c *ColumnSchema) With(options ...ColumnOption) *ColumnSchema {
	for _, option := range options {
		option(c)
	}
	return c
}

// store current column to old column
func (c *ColumnSchema) store() {
	c.oldColumn = c.copy()
}

func (c *ColumnSchema) copy() *ColumnSchema {
	return &ColumnSchema{
		name:              c.name,
		rename:            c.rename,
		dataType:          c.dataType,
		defaultValue:      c.defaultValue,
		nullable:          c.nullable,
		allowedValues:     c.allowedValues,
		length:            c.length,
		precision:         c.precision,
		scale:             c.scale,
		primaryKey:        c.primaryKey,
		autoIncrement:     c.autoIncrement,
		unsigned:          c.unsigned,
		generateAs:        c.generateAs,
		always:            c.always,
		charset:           c.charset,
		collate:           c.collate,
		comment:           c.comment,
		after:             c.after,
		first:             c.first,
		append:            c.append,
		added:             c.added,
		changed:           c.changed,
		storedAs:          c.storedAs,
		useCurrent:        c.useCurrent,
		virtualAs:         c.virtualAs,
		projection:        c.projection,
		srid:              c.srid,
		hasNullable:       c.hasNullable,
		hasLength:         c.hasLength,
		hasPrecisionScale: c.hasPrecisionScale,
		hasAllowedValues:  c.hasAllowedValues,
	}
}
