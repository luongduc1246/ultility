package elasticsearch

import (
	"fmt"
	"testing"

	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

func TestParseMatch(t *testing.T) {
	q := fulltextsearch.NewQuery()
	s := "match{message{query:this is a test,operator:and,zero_terms_query:test},field{analyzer:test,auto_generate_synonyms_phrase_query:true,minimum_should_match:54}}"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Match["field"].MinimumShouldMatch)
}
func TestParseIntervalAllof(t *testing.T) {
	q := fulltextsearch.NewQuery()
	s := "intervals{fields{boost:3,all_of{ordered:true,intervals[{all_of{max_gaps:4}}]}}}"
	q.Parse(s)
	showAll(q)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Intervals["fields"].AllOf.Intervals)
}
func TestParseCombinedFields(t *testing.T) {
	q := fulltextsearch.NewQuery()
	s := "combined_fields{boost:3,fields[a,b,c]}"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.CombinedFields)
}

func showAll(q fulltextsearch.Querier) {
	params := q.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			queri, ok := value.(fulltextsearch.Querier)
			if ok {
				showAll(queri)
			}
			fmt.Println(key, value)
		}
	case []interface{}:
		for _, value := range t {
			queri, ok := value.(fulltextsearch.Querier)
			if ok {
				showAll(queri)
			}
			fmt.Println("mang", value)
		}
	}
}
