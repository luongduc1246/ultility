package fulltextsearch

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	s := "bool{test:abc,filter[{match{query:test}},{term{value:3}}]}"
	q := NewQuery()
	err := q.Parse(s)
	fmt.Println(err)
	showAll(q)
}
func TestParseSlice(t *testing.T) {
	s := "bool{test:abc,filter[[a,b,c],[d,e,f]]}"
	q := NewQuery()
	err := q.Parse(s)
	fmt.Println(err)
	showAll(q)
}

func showAll(q Querier) {
	params := q.GetParams()
	fmt.Println(reflect.TypeOf(q))
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			queri, ok := value.(Querier)
			if ok {
				showAll(queri)
			}
			fmt.Println(key, value)
		}
	case []interface{}:
		for _, value := range t {
			queri, ok := value.(Querier)
			if ok {
				showAll(queri)
			}
			fmt.Println("mang", value)
		}
	}
}
