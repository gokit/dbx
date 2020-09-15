package dbx

// ToSql return the raw sql and args
func ToSql(d Dialect, builder Builder) (string, []interface{}, error) {
	buf := NewBuffer()
	err := builder.Build(d, buf)

	return buf.String(), buf.Value(), err
}

// ToRawSql return the raw sql
func ToRawSql(d Dialect, builder Builder) (string, error) {
	i := interpolator{
		Buffer:       NewBuffer(),
		Dialect:      d,
		IgnoreBinary: true,
	}
	err := i.encodePlaceholder(builder, true)

	return i.String(), err
}
