package elasticsearch

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/luongduc1246/ultility/reqparams"
)

type Document struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func TestConnect(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	info, err := es.client.Info().Do(context.Background())
	fmt.Println(err)
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
		Title:     "Title",
		Content:   "content for this document.",
		Author:    "Jane Doe",
		CreatedAt: time.Now(),
	}
	err := es.Insert(context.Background(), "test", "1", doc)
	fmt.Println(err)
}
func TestInsert2(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := Document{
		Title:     "Toi La Ai",
		Content:   "Xin Hay cho toi biet",
		Author:    "Duc Cot",
		CreatedAt: time.Now(),
	}
	err := es.Insert(context.Background(), "test", "2", doc)
	fmt.Println(err)
}
func TestInsert3(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := Document{
		Title:     "Cách để làm giàu",
		Content:   "Cách để làm giàu một cách nhanh chóng",
		Author:    "Lương Ngọc Đức",
		CreatedAt: time.Now(),
	}
	err := es.Insert(context.Background(), "test", "3", doc)
	fmt.Println(err)
}
func TestInsert4(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	doc := Document{
		Title:     "Cach trong cay",
		Content:   "Cach cay phat trien nhanh chong",
		Author:    "Luong Ngoc Duc",
		CreatedAt: time.Now(),
	}
	err := es.Insert(context.Background(), "test", "4", doc)
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
func TestSearchMatchAll(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	docs := []Document{}
	search := reqparams.NewSearch()
	q := reqparams.NewQuery()
	s := "match_all{}"
	q.Parse(s)
	search.Query = q
	qSort := reqparams.NewSlice()
	qSort.Parse("{options{created_at{order:desc}}}")
	search.Sort = qSort

	err := es.Search(context.Background(), "test", *search, &docs)
	fmt.Println(err, docs)
}
func TestSearchMatch(t *testing.T) {
	es := NewElasticSearch(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	docs := []Document{}
	search := reqparams.NewSearch()
	q := reqparams.NewQuery()
	s := "match{title{query:Cach}}"
	q.Parse(s)
	search.Query = q
	err := es.Search(context.Background(), "test", *search, &docs)
	fmt.Println(err, docs)
}
