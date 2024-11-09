package search

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	s := "bool{filter[{match{query:test}},{term{value:3}}],should[{match{query:test_should}},{term{query{name:testquery}}}]},ab:b"
	q := NewQuery()
	err := q.Parse(s)
	fmt.Println(err)
	fmt.Println(q)
	fmt.Println(q.Params["bool"].(Querier).GetParams().(map[string]interface{})["should"].(Querier).GetParams().([]interface{})[1].(Querier).GetParams())
}
