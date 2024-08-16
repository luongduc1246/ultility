package gormdb

import (
	"github.com/luongduc1246/ultility/reqparams"

	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type filterWhere struct {
	Exps  []clause.Expression
	Joins map[string][]clause.Join
}

func newFilterWhere() *filterWhere {
	return &filterWhere{
		Exps:  make([]clause.Expression, 0),
		Joins: make(map[string][]clause.Join),
	}
}

func (f *filterWhere) parse(scm *schema.Schema, filter *reqparams.Filter) {
	f.Exps = parseFilterToFilterWhere(scm, filter, f.Joins)
}

func parseFilterToFilterWhere(scm *schema.Schema, filter reqparams.IFilter, joins map[string][]clause.Join) []clause.Expression {
	exps := []clause.Expression{}
	fieldsByDB := scm.FieldsByDBName
	for _, exp := range filter.GetExps() {
		switch ex := exp.(type) {
		case reqparams.Eq:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Eq{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.Neq:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Neq{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.Lt:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Lt{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.Lte:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Lte{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.Gt:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Gt{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.Gte:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Gte{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.Like:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.Like{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Value: ex.Value,
				})
			}
		case reqparams.In:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, clause.IN{
					Column: clause.Column{
						Name:  ex.Column,
						Table: scm.Table,
					},
					Values: ex.Values,
				})
			}
		/* làm việc với Json */
		case reqparams.Extract:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, datatypes.JSONQuery(ex.Column).Extract(ex.Value))
			}
		case reqparams.Contains:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, datatypes.JSONArrayQuery(ex.Column).Contains(ex.Value))
			}
		case reqparams.Haskey:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, datatypes.JSONQuery(ex.Column).HasKey(ex.Values...))
			}
		case reqparams.Likes:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, datatypes.JSONQuery(ex.Column).Likes(ex.Value, ex.Keys...))
			}
		case reqparams.Equals:
			if _, ok := fieldsByDB[ex.Column]; ok {
				exps = append(exps, datatypes.JSONQuery(ex.Column).Equals(ex.Value, ex.Keys...))
			}
		case reqparams.Or:

			expcs := parseFilterToFilterWhere(scm, ex, joins)
			if len(expcs) > 0 {
				exps = append(exps, clause.Or(
					expcs...,
				))
			}

		case reqparams.Not:
			expcs := parseFilterToFilterWhere(scm, ex, joins)
			if len(expcs) > 0 {
				exps = append(exps, clause.Not(
					expcs...,
				))
			}
		case reqparams.And:
			expcs := parseFilterToFilterWhere(scm, ex, joins)
			if len(expcs) > 0 {
				exps = append(exps, clause.And(
					expcs...,
				))
			}
		}
	}

	if filter.GetRelatives() != nil {
		for key, f := range filter.GetRelatives() {
			relations := scm.Relationships.Relations
			if rs, ok := relations[key]; ok {
				nameCache := scm.Table + rs.JoinTable.Table
				if join, exist := cacheJoin[nameCache]; exist {
					joins[nameCache] = join
				} else {
					join := []clause.Join{}
					switch rs.Type {
					case schema.Many2Many:
						for _, rf := range rs.References {
							if scm.Table == rf.PrimaryKey.Schema.Table {
								join = append(join, clause.Join{
									Table: clause.Table{Name: rf.ForeignKey.Schema.Table},
									ON: clause.Where{
										Exprs: []clause.Expression{clause.Eq{Column: clause.Column{
											Name:  rf.PrimaryKey.DBName,
											Table: rf.PrimaryKey.Schema.Table,
										}, Value: clause.Column{
											Name:  rf.ForeignKey.DBName,
											Table: rf.ForeignKey.Schema.Table,
										}}},
									},
								})
							} else {
								join = append(join, clause.Join{
									Table: clause.Table{Name: rf.PrimaryKey.Schema.Table},
									ON: clause.Where{
										Exprs: []clause.Expression{clause.Eq{Column: clause.Column{
											Name:  rf.PrimaryKey.DBName,
											Table: rf.PrimaryKey.Schema.Table,
										}, Value: clause.Column{
											Name:  rf.ForeignKey.DBName,
											Table: rf.ForeignKey.Schema.Table,
										}}},
									},
								})
							}
						}
					default:
						for _, rf := range rs.References {
							join = append(join, clause.Join{
								Table: clause.Table{Name: rf.ForeignKey.Schema.Table},
								ON: clause.Where{
									Exprs: []clause.Expression{clause.Eq{Column: clause.Column{
										Name:  rf.PrimaryKey.DBName,
										Table: rf.PrimaryKey.Schema.Table,
									}, Value: clause.Column{
										Name:  rf.ForeignKey.DBName,
										Table: rf.ForeignKey.Schema.Table,
									}}},
								},
							})
						}
					}
					cacheJoin[nameCache] = join
					joins[nameCache] = join
				}
				expcs := parseFilterToFilterWhere(rs.FieldSchema, f, joins)

				if len(expcs) > 0 {
					exps = append(exps,
						expcs...,
					)
				}
			}
		}
	}
	return exps
}
