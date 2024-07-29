package gormdb

import (
	"github.com/luongduc1246/ultility/reqparams"

	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type sortBy struct {
	Columns   []clause.OrderByColumn
	Relatives map[string][]clause.OrderByColumn
}

func newSortBy() *sortBy {
	return &sortBy{
		Columns:   make([]clause.OrderByColumn, 0),
		Relatives: make(map[string][]clause.OrderByColumn),
	}
}

func (s *sortBy) parse(scm *schema.Schema, sort *reqparams.Sort) {
	s.Columns = parseSortToSortBy(scm, sort, s.Relatives)
}

func parseSortToSortBy(scm *schema.Schema, sort *reqparams.Sort, mapRelas map[string][]clause.OrderByColumn) []clause.OrderByColumn {
	columns := []clause.OrderByColumn{}
	fieldsByDB := scm.FieldsByDBName
	for _, order := range sort.Orders {
		if _, ok := fieldsByDB[order.Column]; ok {
			columns = append(columns, clause.OrderByColumn{
				Column: clause.Column{
					Table: scm.Table,
					Name:  order.Column,
				},
				Desc: order.Desc,
			})
		}
	}
	if sort.Relatives != nil {
		for key, rl := range sort.Relatives {
			relations := scm.Relationships.Relations
			if rs, ok := relations[key]; ok {
				colRela := parseSortToSortBy(rs.FieldSchema, rl, mapRelas)
				if len(colRela) > 0 {
					name := rs.JoinTable.Table
					mapRelas[name] = colRela
				}
			}
		}
	}
	return columns
}
