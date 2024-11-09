package elasticsearch

import (
	"context"
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/luongduc1246/ultility/reqparams"
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
