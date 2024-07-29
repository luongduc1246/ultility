package reqparams

import (
	"fmt"
	"testing"
)

func TestParseQueryFilterOld(t *testing.T) {
	s := `eq[adfa]=%25name%25,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]`
	f, err := ParseQueryToFilterOld(nil, s)
	fmt.Println(f, err)
	// s = `eq[adfa]=%25name%25`
	// fmt.Println(ParseQueryToFilter(nil, s))
}

func BenchmarkParseQueryFilterOld(b *testing.B) {
	s := `eq[phone]=test,like[name]=babaa,in[id]=1;2;3,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]`
	for i := 0; i < b.N; i++ {
		ParseQueryToFilterOld(nil, s)
	}
}

func TestParseQueryFilterMap(t *testing.T) {
	s := `eq[phone]=adsf,extract[name]=ccc,likes[name]=abc;abc:test`
	f := NewFilter()
	f.Parse(s)
	fmt.Println(f)

	// s = `eq[adfa]=%25name%25`
	// fmt.Println(ParseQueryToFilter(nil, s))
}
func BenchmarkParseQueryFilterMap(b *testing.B) {
	s := `eq[phone]=test,like[name]=babaa,in[id]=1;2;3,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]`
	for i := 0; i < b.N; i++ {
		NewFilter().Parse(s)
	}
}
