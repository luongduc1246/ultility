package elasticsearch

import (
	"fmt"
	"testing"

	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

func TestParseMatch(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "match[message[query=this is a test,operator=and,zero_terms_query=test],field[analyzer=test,auto_generate_synonyms_phrase_query=true]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Match)
}
func TestParseIntervalAllof(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "intervals[fields[all_of[ordered=true,slice_intervals[intervals[all_of[ordered=false]],intervals[all_of[ordered=true]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Intervals["fields"].AllOf.Intervals[0].AllOf.Ordered)
}
func TestParseCombinedFields(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "combined_fields[boost=3,fields=a;b;c]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.CombinedFields)
}
