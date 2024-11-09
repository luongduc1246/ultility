package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/luongduc1246/ultility/reqparams"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

var instance *Elasticsearch

type Elasticsearch struct {
	client *elasticsearch.TypedClient
}

func NewElasticSearch(conf elasticsearch.Config) *Elasticsearch {
	if instance != nil {
		return instance
	}
	cli, err := elasticsearch.NewTypedClient(conf)
	if err != nil {
		panic(err)
	}
	return &Elasticsearch{
		client: cli,
	}
}

func (e Elasticsearch) CreateIndex(ctx context.Context, name string) error {
	_, err := e.client.Indices.Create(name).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e Elasticsearch) Insert(ctx context.Context, index string, id string, body interface{}) (err error) {
	_, err = e.client.Index(index).Id(id).Request(body).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e Elasticsearch) Update(ctx context.Context, index string, id string, body interface{}) (err error) {
	_, err = e.client.Update(index, id).Doc(body).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e Elasticsearch) Delete(ctx context.Context, index string, id string) (err error) {
	_, err = e.client.Delete(index, id).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e Elasticsearch) Get(ctx context.Context, index string, id string, model interface{}) (err error) {
	res, err := e.client.Get(index, id).Do(ctx)
	m, _ := res.Source_.MarshalJSON()
	err = json.Unmarshal(m, model)
	if err != nil {
		return err
	}
	return nil
}
func (e Elasticsearch) Search(ctx context.Context, index string, param reqparams.Search, models interface{}) (err error) {
	query := ParseQueryToSearch(param.Query)

	c := e.client.Search().Index(index).Query(query)
	res, err := c.Do(ctx)
	fmt.Println(res.Hits.Hits)
	return nil
}

func ParseQueryToSearch(q fulltextsearch.Querier) *types.Query {
	query := types.NewQuery()
	params := q.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			/* compound */
			case fulltextsearch.BOOL:
				b := ParseBoolQuery(value.(fulltextsearch.Querier))
				query.Bool = b
			case fulltextsearch.BOOSTING:
				b := ParseBoostingQuery(value.(fulltextsearch.Querier))
				query.Boosting = b
			case fulltextsearch.CONSTANTSCORE:
				c := ParseConstantScoreQuery(value.(fulltextsearch.Querier))
				query.ConstantScore = c
			case fulltextsearch.DISMAX:
				d := ParseDisMaxQuery(value.(fulltextsearch.Querier))
				query.DisMax = d
			case fulltextsearch.FUNCTIONSCORE:
				d := ParseFunctionScoreQuery(value.(fulltextsearch.Querier))
				query.FunctionScore = d
			/* fulltext */
			case fulltextsearch.MATCH:
				match := ParseMatchQuery(value.(fulltextsearch.Querier))
				query.Match = match
			case fulltextsearch.INTERVALS:
				in := ParseIntervalsQuery(value.(fulltextsearch.Querier))
				query.Intervals = in
			case fulltextsearch.MATCHBOOLPREFIX:
				in := ParseMatchBoolPrefixQuery(value.(fulltextsearch.Querier))
				query.MatchBoolPrefix = in
			case fulltextsearch.MATCHPHRASE:
				in := ParseMatchPhraseQuery(value.(fulltextsearch.Querier))
				query.MatchPhrase = in
			case fulltextsearch.MATCHPHRASEPREFIX:
				in := ParseMatchPhrasePrefixQuery(value.(fulltextsearch.Querier))
				query.MatchPhrasePrefix = in
			case fulltextsearch.COMBINEDFIELDS:
				in := ParseCombinedFieldsQuery(value.(fulltextsearch.Querier))
				query.CombinedFields = in
			case fulltextsearch.MULTIMATCH:
				in := ParseMultiMatchQuery(value.(fulltextsearch.Querier))
				query.MultiMatch = in
			case fulltextsearch.QUERYSTRING:
				in := ParseQueryStringQuery(value.(fulltextsearch.Querier))
				query.QueryString = in
			case fulltextsearch.SIMPLEQUERYSTRING:
				in := ParseSimpleQueryStringQuery(value.(fulltextsearch.Querier))
				query.SimpleQueryString = in
				/* joining */
			case fulltextsearch.NESTED:
				in := ParseNestedQuery(value.(fulltextsearch.Querier))
				query.Nested = in
			}

		}
	}
	// switch q.(type) {
	// case *fulltextsearch.QuerySearch:
	// 	params := q.GetParams().(map[fulltextsearch.QueryKey]interface{})
	// 	for key, value := range params {
	// 		q := ParseQueryToSearch(value.(fulltextsearch.Querier))

	// 	}
	// case fulltextsearch.Bool:
	// 	bo := ParseBoolQuery(q)

	// 	query.Bool = bo
	// case fulltextsearch.Match:
	// 	match := ParseMatchQuery(q)
	// 	query.Match = match
	// case fulltextsearch.Intervals:
	// 	i := ParseIntervalsQuery(q)
	// 	query.Intervals = i
	// }
	return query
}

// func (e Elasticsearch) Delete(index string, id string, otps ...fulltextsearch.Optioner) (res io.Reader, err error) {
// 	o := arrays.ConvertSliceTypeToSliceType[func(*esapi.DeleteRequest), fulltextsearch.Optioner](otps)
// 	es, err := e.client.Delete(index, id, o...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return es.Body, nil
// }
// func (e Elasticsearch) Search(otps ...fulltextsearch.Optioner) (res io.Reader, err error) {
// 	o := arrays.ConvertSliceTypeToSliceType[func(*esapi.Search), fulltextsearch.Optioner](otps)
// 	es, err := e.client.Search(o...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return es.Body, nil
// }
