package dbx

type OrBuildFunc func(Dialect, Buffer) error

// Build calls itself to build SQL.
func (b OrBuildFunc) Build(d Dialect, buf Buffer) error {
	return b(d, buf)
}
