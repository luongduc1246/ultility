package gormdb

import "gorm.io/gorm/clause"

type With struct {
	Table clause.Table
	Expr  *clause.Expr
}

func (w With) Name() string {
	return "WITH"
}

func (w With) Build(builder clause.Builder) {
	if w.Table.Name != "" {
		builder.WriteString("WITH")
		builder.WriteByte(' ')
		builder.WriteQuoted(w.Table)
		builder.WriteByte(' ')
		builder.WriteString("AS")
		builder.WriteByte('(')
		if w.Expr != nil {
			w.Expr.Build(builder)
		}
		builder.WriteByte(')')
		builder.WriteByte(' ')
	}
}

func (w With) MergeClause(clause *clause.Clause) {
	clause.Name = ""
	clause.Expression = w
}
