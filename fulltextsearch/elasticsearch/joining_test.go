package elasticsearch

import (
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/luongduc1246/ultility/reqparams"
)

func TestNested(t *testing.T) {
	q := reqparams.NewQuery()
	s := "nested{boost:42,ignore_unmapped:true,path:test,score_mode:babe}"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Nested)
}
func TestNestedInnerHit(t *testing.T) {
	q := reqparams.NewQuery()
	s := "nested{inner_hits{collapse{field:id,max_concurrent_group_searches:3},docvalue_fields[{field:testdoc}]}}"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Nested.InnerHits.DocvalueFields)
}
func TestNestedInnerHitSort(t *testing.T) {
	q := reqparams.NewQuery()
	s := "nested{inner_hits{sort[{_doc{order:testorder},_geo_distance{geo_distance_sort{fields[[1,2,4],[7,6,8]],test[abc,bca]}}}]}}"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Nested.InnerHits.Sort[0].(*types.SortOptions).GeoDistance_.GeoDistanceSort)
}
