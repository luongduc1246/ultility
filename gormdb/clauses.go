package gormdb

import (
	"github.com/luongduc1246/ultility/reqparams"

	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

var cacheJoin map[string][]clause.Join

func init() {
	cacheJoin = make(map[string][]clause.Join)
}

type ClauseSearch struct {
	Joins   map[string][]clause.Join
	Where   *clause.Where
	OrderBy *clause.OrderBy
	Limit   *clause.Limit
}

func NewClauseSearch() *ClauseSearch {
	return &ClauseSearch{
		Joins: make(map[string][]clause.Join),
	}
}

func (cs *ClauseSearch) Parse(scm *schema.Schema, search *reqparams.Search) {
	if search.Filter != nil {
		fw := newFilterWhere()
		filter := reqparams.NewFilter()
		filter.ParseFromQuerier(search.Filter)
		fw.parse(scm, filter)
		if len(fw.Exps) > 0 {
			cs.Where = &clause.Where{
				Exprs: fw.Exps,
			}
		}
		if len(fw.Joins) > 0 {
			for k, v := range fw.Joins {
				cs.Joins[k] = v
			}
		}
	}
	if search.Sort != nil {
		sb := newSortBy()
		sort := reqparams.NewSort()
		sort.ParseQuerierToSort(search.Sort)
		sb.parse(scm, sort)
		for k, v := range sb.Relatives {
			if _, ok := cs.Joins[k]; ok {
				sb.Columns = append(sb.Columns, v...)
			}
		}
		if len(sb.Columns) > 0 {
			cs.OrderBy = &clause.OrderBy{
				Columns: sb.Columns,
			}
		}
	}
	page := NewPageLimit()
	page.Parse(search.Page, search.Limit)
	cs.Limit = &clause.Limit{
		Limit:  &page.Limit,
		Offset: page.Offset,
	}
}

func (cs *ClauseSearch) Build() []clause.Expression {
	cexps := []clause.Expression{}
	if len(cs.Joins) > 0 {
		joins := []clause.Join{}
		for _, v := range cs.Joins {
			joins = append(joins, v...)
		}
		cexps = append(cexps, clause.From{
			Joins: joins,
		})
	}
	if cs.Where != nil {
		cexps = append(cexps, cs.Where)
	}
	if cs.OrderBy != nil {
		cexps = append(cexps, cs.OrderBy)
	}
	cexps = append(cexps, cs.Limit)
	return cexps
}
