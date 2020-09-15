package schema

// The table interface
type Table interface {
	Name() string             // The name of the table
	Schema() string           // The schema of the table
	Prefix() string           // The prefix of the table.
	Engine() string           // The storage engine that should be used for the table.
	Charset() string          // The default character set that should be used for the table.
	Collation() string        // The collation that should be used for the table.
	IsTemporary() bool        // Whether to make the table temporary.
	Comment() string          // the comment of the table
	Columns() []Column        // The columns that should be added to the table.
	PrimaryKey() []string     // The primary keys of the table
	AddedColumns() []Column   // the added columns
	ChangedColumns() []Column // the changed columns
	Commands() []TableCommand // the table commands
}
