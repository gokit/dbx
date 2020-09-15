package schema

// Column interface
type Column interface {
	Name() string                 // name of this column (without quotes).
	Rename() string               // change the name of this column (without quotes).
	DataType() DataType           // the abstract DB type of this column. Possible DB types vary according to the type of DBMS.
	DefaultValue() interface{}    // Specify a "default" value for the column, the value type must be a string or fmt.Stringer interface
	Nullable() bool               // Allow NULL values to be inserted into the column
	AllowedValues() []interface{} // set column allowed values when column type is enum or set
	Length() int                  // display size of the column.
	Precision() int               // precision of the column data, if it is numeric.
	Scale() int                   // scale of the column data, if it is numeric
	PrimaryKey() bool             // is a primary index
	AutoIncrement() bool          // Set INTEGER columns as auto-increment (primary key)
	Unsigned() bool               // whether this column is unsigned. This is only meaningful when [[type]] is `smallint`, `integer` or `bigint`.
	GeneratedAs() string          // Create a SQL compliant identity column (PostgreSQL)
	Alawys() bool                 // Used as a modifier for generatedAs() (PostgreSQL)
	Charset() string              // Specify a character set for the column (MySQL)
	Collate() string              // Specify a collate for the column (MySQL)
	After() string                // Place the column "after" another column (MySQL)
	First() string                // Place the column "first" in the table (MySQL)
	Append() string               // append
	Added() bool                  // Change the columns
	Changed() bool                // Change the columns
	StoredAs() string             // the stored generated column (MySQL)
	Srid() int                    // the srid of the column
	UseCurrent() bool             // the TIMESTAMP column to use CURRENT_TIMESTAMP as default value
	VirtualAs() string            // the virtual generated column (MySQL)
	Projection() string           //
	Comment() string              // the comment of the column
}
