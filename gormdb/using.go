package gormdb

import "gorm.io/gorm/clause"

type Using struct {
	Table string
}

func (d Using) Name() string {
	return "USING"
}

func (d Using) Build(builder clause.Builder) {
	builder.WriteString("USING")

	if d.Table != "" {
		builder.WriteByte(' ')
		builder.WriteString(d.Table)
	}
}

func (d Using) MergeClause(clause *clause.Clause) {
	clause.Name = ""
	clause.Expression = d
}
