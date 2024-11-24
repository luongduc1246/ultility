package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/luongduc1246/ultility/reqparams"
)

func ParseSortQueryToSort(sort reqparams.Querier) []types.SortCombinations {
	sliceQuery := []types.SortCombinations{}
	options, ok := sort.GetParams().([]interface{})
	if !ok {
		return nil
	}
	for _, q := range options {
		switch q.(type) {
		case *reqparams.Query:
			v, ok := q.(reqparams.Querier)
			if !ok {
				break
			}
			sort := ParseSortCombinations(v)
			sliceQuery = append(sliceQuery, sort)
		case string:
			sliceQuery = append(sliceQuery, q)
		}
	}
	return sliceQuery
}
