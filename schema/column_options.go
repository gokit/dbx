package schema

type ColumnOption func(column *ColumnSchema)

// rename the column name
func Rename(rename string) ColumnOption {
	return func(column *ColumnSchema) {
		column.rename = rename
	}
}

// Place the column "after" another column (MySQL)
func After(after string) ColumnOption {
	return func(column *ColumnSchema) {
		column.after = after
	}
}

// Set INTEGER columns as auto-increment (primary key)
func AutoIncrement() ColumnOption {
	return func(column *ColumnSchema) {
		column.autoIncrement = true
	}
}

// Set INTEGER columns as not auto-increment (primary key)
func NotAutoIncrement() ColumnOption {
	return func(column *ColumnSchema) {
		column.autoIncrement = false
	}
}

// set column allowed values when column type is enum、set
func AllowedValues(values ...interface{}) ColumnOption {
	return func(column *ColumnSchema) {
		column.allowedValues = values
	}
}

// Specify a character set for the column (MySQL)
func Charset(charset string) ColumnOption {
	return func(column *ColumnSchema) {
		column.charset = charset
	}
}

// Add a comment to the column (MySQL)
func Comment(comment string) ColumnOption {
	return func(column *ColumnSchema) {
		column.comment = comment
	}
}

// Specify a collation for the column (MySQL/PostgreSQL/SQL Server)
func Collation(collation string) ColumnOption {
	return func(column *ColumnSchema) {
		column.collate = collation
	}
}

// Specify a collation for the column (MySQL/PostgreSQL/SQL Server)
func Default(value interface{}) ColumnOption {
	return func(column *ColumnSchema) {
		column.defaultValue = value
	}
}

// Specify a collation for the column (MySQL/PostgreSQL/SQL Server)
func First(first string) ColumnOption {
	return func(column *ColumnSchema) {
		column.first = first
	}
}

// Create a SQL compliant identity column (PostgreSQL)
func GeneratedAs(expression string, always ...bool) ColumnOption {
	return func(column *ColumnSchema) {
		column.generateAs = expression
		if len(always) > 0 {
			column.always = always[0]
		}
	}
}

// set column nullable
func Nullable(nullable bool) ColumnOption {
	return func(column *ColumnSchema) {
		column.nullable = nullable
	}
}

// Add a primary index
func PrimaryKey() ColumnOption {
	return func(column *ColumnSchema) {
		column.primaryKey = true
	}
}

// Create a stored generated column (MySQL)
func StoredAs(expression string) ColumnOption {
	return func(column *ColumnSchema) {
		column.storedAs = expression
	}
}

// Specify a type for the column
func Type(dataType DataType) ColumnOption {
	return func(column *ColumnSchema) {
		column.dataType = dataType
	}
}

// Set the INTEGER column as UNSIGNED (MySQL)
func Unsigned() ColumnOption {
	return func(column *ColumnSchema) {
		column.unsigned = true
	}
}

// Set the TIMESTAMP column to use CURRENT_TIMESTAMP as default value
func UseCurrent() ColumnOption {
	return func(column *ColumnSchema) {
		column.useCurrent = true
	}
}

// Create a virtual generated column (MySQL)
func VirtualAs(expression string) ColumnOption {
	return func(column *ColumnSchema) {
		column.virtualAs = expression
	}
}

// set column length when column type is string、int
func Length(length int) ColumnOption {
	return func(column *ColumnSchema) {
		column.length = length
	}
}

// set column precision when column type is decimal
func Precision(precision int) ColumnOption {
	return func(column *ColumnSchema) {
		column.precision = precision
	}
}

// set column scale when column type is decimal
func Scale(scale int) ColumnOption {
	return func(column *ColumnSchema) {
		column.scale = scale
	}
}

// set column precision and scale when column type is decimal
func Size(precision int, scale ...int) ColumnOption {
	return func(column *ColumnSchema) {
		column.precision = precision
		if len(scale) > 0 {
			column.scale = scale[0]
		}
	}
}

// set column not null
func NotNull() ColumnOption {
	return func(column *ColumnSchema) {
		column.nullable = false
	}
}

// allow column null
func AllowNull() ColumnOption {
	return func(column *ColumnSchema) {
		column.nullable = true
	}
}

// set column projection
func Projection(projection string) ColumnOption {
	return func(column *ColumnSchema) {
		column.projection = projection
	}
}
