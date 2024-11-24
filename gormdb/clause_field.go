package gormdb

import (
	"github.com/luongduc1246/ultility/reqparams"

	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type fieldSelect struct {
	Columns []clause.Column
	Joins   map[string][]clause.Join
}

func newFieldSelect() *fieldSelect {
	return &fieldSelect{
		Columns: make([]clause.Column, 0),
		Joins:   make(map[string][]clause.Join),
	}
}

func (s *fieldSelect) parse(scm *schema.Schema, field *reqparams.Fields) {
	fieldsByDB := scm.FieldsByDBName
	for _, name := range field.Columns {
		if _, ok := fieldsByDB[name]; ok {
			s.Columns = append(s.Columns, clause.Column{
				Table: scm.Table,
				Name:  name,
			})
		}
	}
	if field.Relatives != nil {
		for key, f := range field.Relatives {
			relations := scm.Relationships.Relations
			if rs, ok := relations[key]; ok {
				nameCache := scm.Table + rs.JoinTable.Table
				if join, exist := cacheJoin[nameCache]; exist {
					s.Joins[nameCache] = join
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
					s.Joins[nameCache] = join
				}
				s.parse(rs.FieldSchema, f)
			}
		}
	}
}
