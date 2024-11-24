package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/luongduc1246/ultility/reqparams"
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
func (e Elasticsearch) DeleteIndex(ctx context.Context, name string) error {
	_, err := e.client.Indices.Delete(name).Do(ctx)
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
	if err != nil {
		return err
	}
	m, _ := res.Source_.MarshalJSON()
	err = json.Unmarshal(m, model)
	if err != nil {
		return err
	}
	return nil
}
func (e Elasticsearch) Search(ctx context.Context, index string, param reqparams.Search, models interface{}) (err error) {
	modelValue := reflect.ValueOf(models)
	if modelValue.Kind() != reflect.Pointer {
		return errors.New("parameter models must be pointer")
	}
	modelValue = reflect.Indirect(modelValue)
	elemType := modelValue.Type().Elem()
	search := e.client.Search()

	if param.Query != nil {
		query := ParseQueryToSearch(param.Query)
		search = search.Index(index).Query(query)
	}
	if param.Sort != nil {
		sort := ParseSortQueryToSort(param.Sort)
		search = search.Sort(sort...)
	}
	from, size := ParseFromAndSize(param.Page, param.Limit)
	search = search.From(from).Size(size)
	res, err := search.Do(ctx)
	if err != nil {
		return err
	}
	if res.Hits.Total.Value > 0 {
		for _, hit := range res.Hits.Hits {
			newElem := reflect.New(elemType)
			err := json.Unmarshal(hit.Source_, newElem.Interface())
			if err == nil {
				if modelValue.CanSet() {
					modelValue.Set(reflect.Append(modelValue, reflect.Indirect(newElem)))
				}
			}
		}
	}
	return nil
}

func ParseFromAndSize(page int, limit int) (from, size int) {
	switch true {
	case limit > MAXSIZE:
		size = MAXSIZE
	case limit < MINSIZE:
		size = MINSIZE
	default:
		size = limit
	}
	if page > 0 {
		from = (page - 1) * size
	}
	return
}

func ParseQueryToSearch(q reqparams.Querier) *types.Query {
	query := types.NewQuery()
	params := q.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			/* compound */
			case "bool":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				b := ParseBoolQuery(quies)
				query.Bool = b
			case "boosting":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				b := ParseBoostingQuery(quies)
				query.Boosting = b
			case "constant_score":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				c := ParseConstantScoreQuery(quies)
				query.ConstantScore = c
			case "dis_max":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				d := ParseDisMaxQuery(quies)
				query.DisMax = d
			case "function_score":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				d := ParseFunctionScoreQuery(quies)
				query.FunctionScore = d
			/* fulltext */
			case "match":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				match := ParseMatchQuery(quies)
				query.Match = match
			case "match_all":
				query.MatchAll = types.NewMatchAllQuery()
			case "intervals":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseIntervalsQuery(quies)
				query.Intervals = in
			case "match_bool_prefix":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseMatchBoolPrefixQuery(quies)
				query.MatchBoolPrefix = in
			case "match_phrase":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseMatchPhraseQuery(quies)
				query.MatchPhrase = in
			case "match_phrase_prefix":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseMatchPhrasePrefixQuery(quies)
				query.MatchPhrasePrefix = in
			case "combined_fields":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseCombinedFieldsQuery(quies)
				query.CombinedFields = in
			case "multi_match":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseMultiMatchQuery(quies)
				query.MultiMatch = in
			case "query_string":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseQueryStringQuery(quies)
				query.QueryString = in
			case "simple_query_string":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseSimpleQueryStringQuery(quies)
				query.SimpleQueryString = in
				/* joining */
			case "nested":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseNestedQuery(quies)
				query.Nested = in
			case "has_child":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseHasChildQuery(quies)
				query.HasChild = in
			case "has_parent":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseHasParentQuery(quies)
				query.HasParent = in
			case "parent_id":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseParentIdQuery(quies)
				query.ParentId = in
				/* term */
			case "exists":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseExistsQuery(quies)
				query.Exists = in
			case "fuzzy":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseFuzzyQuery(quies)
				query.Fuzzy = in
			case "ids":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseIdsQuery(quies)
				query.Ids = in
			case "prefix":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParsePrefixQuery(quies)
				query.Prefix = in
			case "range":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseRangeQuery(quies)
				query.Range = in
			case "regexp":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseRegexpQuery(quies)
				query.Regexp = in
			case "term":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseTermQuery(quies)
				query.Term = in
			case "terms":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseTermsQuery(quies)
				query.Terms = in
			case "terms_set":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseTermsQuery(quies)
				query.Terms = in
			case "wildcard":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseWildcardQuery(quies)
				query.Wildcard = in
				/* special */
			case "distance_feature":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseDistanceFeatureQuery(quies)
				query.DistanceFeature = in
			case "more_like_this":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseMoreLikeThisQuery(quies)
				query.MoreLikeThis = in
			case "percolate":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParsePercolateQuery(quies)
				query.Percolate = in
			case "rank_feature":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseRankFeatureQuery(quies)
				query.RankFeature = in
			case "script":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseScriptQuery(quies)
				query.Script = in
			case "script_score":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseScriptScoreQuery(quies)
				query.ScriptScore = in
			case "wrapper":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseWrapperQuery(quies)
				query.Wrapper = in
			case "pinned":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParsePinnedQuery(quies)
				query.Pinned = in
			case "rule":
				quies, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				in := ParseRuleQuery(quies)
				query.Rule = in
			}

		}
	}

	return query
}
