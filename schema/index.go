package schema

type IndexType string

const (
	PrimaryIndex  IndexType = "PRIMARY KEY" //
	UniqueIndex   IndexType = "UNIQUE"      //
	NormalIndex   IndexType = "INDEX"       //
	SpatialIndex  IndexType = "SPATIAL"     // [MySQL:SPATIAL] [PostgreSQL:GIST]
	FullTextIndex IndexType = "FULLTEXT"    // [MySQL:FULLTEXT]
	GistIndex     IndexType = "GIST"        // [MySQL:SPATIAL] [PostgreSQL:GIST]
	GinIndex      IndexType = "GIN"         // [PostgreSQL:GIN]
)

type Index interface {
	Name() string          // index name
	Type() IndexType       // index type
	ColumnNames() []string // the column names of index
	Algorithm() string     // index algorithm
	Comment() string       // the comment of index
}
