package gormdb

import "gorm.io/gorm/clause"

type Using struct {
	Table clause.Table
}

func (u Using) Name() string {
	return "USING"
}

func (u Using) Build(builder clause.Builder) {
	builder.WriteString("USING ")
	builder.WriteQuoted(u.Table)
}

func (u Using) MergeClause(clause *clause.Clause) {
	clause.Name = ""
	clause.Expression = u
}
