package reqparams

import (
	"fmt"
	"testing"
)

func TestParseQueryToSortMap(t *testing.T) {
	s := "test,{a:test},{options{title{order:desc}}},{test[abc,{ab:bc},{options{field{order:asc}}}]}"
	query := NewSlice()
	err := query.Parse(s)
	fmt.Println(query)
	fmt.Println(err)
	sort := NewSort()
	sort.ParseQuerierToSort(query)
	fmt.Println(sort)
	fmt.Println(sort.Relatives["Test"])

}
