package elasticsearch

import (
	"context"
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/luongduc1246/ultility/reqparams"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func TestConnect(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	info, _ := es.client.Info().Do(context.Background())
	fmt.Println(info)
}
func TestCreateIndex(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	err := es.CreateIndex(context.Background(), "test")
	fmt.Println(err)
}
func TestInsert(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := Document{
		Title:   "Title",
		Content: "content for this document.",
		Author:  "Jane Doe",
	}
	err := es.Insert(context.Background(), "test", "1", doc)
	fmt.Println(err)
}
func TestUpdate(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := Document{
		Title: "Title Haha",
	}
	err := es.Update(context.Background(), "test", "1", doc)
	fmt.Println(err)
}
func TestGet(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := &Document{}
	err := es.Get(context.Background(), "test", "1", doc)
	fmt.Println(err, doc)
}
func TestSearch(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := Document{}
	search := reqparams.NewSearch()
	search.Filter = &reqparams.Filter{
		Exps: []reqparams.Exp{
			reqparams.Eq{
				Column: "content",
				Value:  "content",
			},
		},
	}
	err := es.Search(context.Background(), "test", *search, &doc)
	fmt.Println(err, doc)
}

func TestParseMatch(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "match[message[query=this is a test,operator=and],field[analyzer=test,auto_generate_synonyms_phrase_query=true]],intervals[fields[all_of[ordered=true,filter[after[all_of[ordered=true]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query)
}
func TestParseIntervals(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "intervals[fields[all_of[ordered=true,filter[after[all_of[ordered=true]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Intervals["fields"].AllOf)
}
func TestParseBoolQuery(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "bool[boost=3,filter[match[field[query=test]],intervals[fields[all_of[ordered=true]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Bool.Filter)
}
func TestParseBoostingQuery(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "boosting[boost=3,negative[match[field[query=3],mess[query=test]],intervals[fields[all_of[ordered=true]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Boosting.Negative.Intervals)
}
func TestParseConstantScoreQuery(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "constant_score[boost=3,filter[match[field[query=test]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.ConstantScore.Filter.Match)
}
func TestParseDismaxQuery(t *testing.T) {
	q := fulltextsearch.NewQuerySearch()
	s := "constant_score[boost=3,filter[match[field[query=test]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.ConstantScore.Filter.Match)
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := fulltextsearch.NewQuerySearch()
		s := "bool[boost=3]"
		q.Parse(s)
		ParseQueryToSearch(q)
	}
}
