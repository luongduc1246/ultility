package fulltextsearch

import (
	"fmt"
	"testing"
)

func TestFullTextSearch(t *testing.T) {
	s := "bool[boost=1.0,minimum_should_match=3,must[term[name=kimchi,age=3]]],boosting[negative[term[test=acd]]]"
	q := NewQuerySearch()
	err := q.Parse(s)
	fmt.Println(err)
	fmt.Println(q)
	fmt.Println(q.GetParams().(map[QueryKey]interface{})["boosting"].(Boosting).GetParams())
}

func BenchmarkBool(b *testing.B) {

	for i := 0; i < b.N; i++ {
		s := "bool[boost=1.0,minimum_should_match=3,must[term[name=kimchi,age=3]]],boosting[negative[term[test=acd]],fields=b;a]"
		q := NewQuerySearch()
		q.Parse(s)
	}
}
