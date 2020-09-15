package schema

import (
	"strings"
)

type IndexSchema struct {
	name        string    // the name of the index
	indexType   IndexType // the type of the index
	columnNames []string  // the columns of the index
	algorithm   string    // the algorithm of the index
	comment     string    // the comment of the index
}

func (i IndexSchema) Name() string {

	if i.name != "" {
		return i.name
	}

	var b strings.Builder
	b.WriteString("idx")

	for _, name := range i.columnNames {
		b.WriteString("_")
		b.WriteString(name)
	}

	return b.String()
}

func (i IndexSchema) Type() IndexType {
	return i.indexType
}

func (i IndexSchema) ColumnNames() []string {
	return i.columnNames
}

func (i IndexSchema) Algorithm() string {
	return i.algorithm
}

func (i IndexSchema) Comment() string {
	return i.comment
}

func newIndex(indexType IndexType, columnNames []string, options ...IndexOptions) *IndexSchema {
	index := &IndexSchema{
		indexType:   indexType,
		columnNames: columnNames,
	}

	for _, option := range options {
		option(index)
	}

	return index
}
