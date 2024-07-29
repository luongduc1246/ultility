package reqparams

import (
	"fmt"
	"testing"
)

func TestParseQueryToSortOld(t *testing.T) {
	s := `asc[name],desc[id],roles[asc[name]desc[id]],abc[asc[babe],desc[name]]`
	fmt.Println(ParseQueryToSortOld(nil, s))
}
func TestParseQueryToSortMap(t *testing.T) {
	s := `asc[name],desc[id],roles[asc[afda],desc[addf]]`
	b := `roles[asc[adfds]]`
	sort := NewSort()
	sort.Parse(s)
	sort.Parse(b)
	fmt.Println(sort)
	fmt.Println(sort.Relatives["Roles"])
}
