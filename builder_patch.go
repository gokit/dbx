package dbx

// orBuildFunc implements Builder.
type orBuildFunc func(Dialect, Buffer) error

// Build calls itself to build SQL.
func (b orBuildFunc) Build(d Dialect, buf Buffer) error {
	return b(d, buf)
}
