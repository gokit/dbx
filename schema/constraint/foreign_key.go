package constraint

type ForeignKey struct {
	Name               string
	ColumnNames        []string
	ForeignSchemaName  string
	ForeignTableName   string
	ForeignColumnNames []string
	OnDelete           string
	OnUpdate           string
}
