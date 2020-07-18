package dbx

func buildLogicCond(d Dialect, buf Buffer, cond ...Builder) error {
	for i, c := range cond {
		if i > 0 {
			buf.WriteString(" ")

			switch c.(type) {
			case orBuildFunc:
				buf.WriteString("OR")
			default:
				buf.WriteString("AND")
			}

			buf.WriteString(" ")
		}

		if len(cond) > 1 {
			buf.WriteString("(")
		}
		err := c.Build(d, buf)
		if err != nil {
			return err
		}
		if len(cond) > 1 {
			buf.WriteString(")")
		}
	}

	return nil
}

func logicCond(cond ...Builder) Builder {
	return BuildFunc(func(d Dialect, buf Buffer) error {
		return buildLogicCond(d, buf, cond...)
	})
}
