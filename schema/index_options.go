package schema

type IndexOptions func(index *IndexSchema)

func IndexName(name string) IndexOptions {
	return func(index *IndexSchema) {
		index.name = name
	}
}

func IndexComment(comment string) IndexOptions {
	return func(index *IndexSchema) {
		index.comment = comment
	}
}

func IndexAlgorithm(algorithm string) IndexOptions {
	return func(index *IndexSchema) {
		index.algorithm = algorithm
	}
}
