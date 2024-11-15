package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

/*
# distance_feature{parameters}

	parameters :
	1. untyped: queries
	2. date: queries
	3. geo: queries
*/
func ParseDistanceFeatureQuery(m fulltextsearch.Querier) types.DistanceFeatureQuery {
	var query types.DistanceFeatureQuery
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "untyped":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query = ParseUntypedDistanceFeatureQuery(field)
			case "date":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query = ParseDateDistanceFeatureQuery(field)
			case "geo":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query = ParseGeoDistanceFeatureQuery(field)

			}
		}
	}
	return query
}

/*
#Phân tích câu query untyped cua range

	{untyped{boost:float32,field:string,origin:json.RawMessage,pivot:string,_name:string}}
*/
func ParseUntypedDistanceFeatureQuery(m fulltextsearch.Querier) types.UntypedDistanceFeatureQuery {
	query := types.UntypedDistanceFeatureQuery{}
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			case "origin":
				v, ok := value.(string)
				if !ok {
					break
				}
				j := json.RawMessage(v)
				query.Origin = j
			case "pivot":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Pivot = json.RawMessage(v)

			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v

			}
		}
	}
	return query
}

/*
Phân tích câu query date cua range

	{date{boost:float32,field:string,origin:string,pivot:string,_name:string}}
*/
func ParseDateDistanceFeatureQuery(m fulltextsearch.Querier) types.DateDistanceFeatureQuery {
	query := types.DateDistanceFeatureQuery{}
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			case "origin":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Origin = v
			case "pivot":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Pivot = v

			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v

			}
		}
	}
	return query
}

/*
Phân tích câu query geo cua range

	{geo{boost:float32,field:string,origin:geolocation,pivot:string,_name:string}}
	- geolocation :
		* {geohash:string}
		* {lat:float64,lon:float64}
		* [float64,float64...]
		* string
*/
func ParseGeoDistanceFeatureQuery(m fulltextsearch.Querier) types.GeoDistanceFeatureQuery {
	query := types.GeoDistanceFeatureQuery{}
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			case "origin":
				query.Origin = ParseGeoLocale(value)
			case "pivot":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Pivot = v

			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v

			}
		}
	}
	return query
}

/*
more_like_this{...parameters}

	parameters:
	- anylyzer : string
	- boost : float32
	- boost_terms : float64
	- fail_on_unsupported_field : bool
	- fields : []string * [david,madam]
	- include : bool
	- like : []interface{}
	- interface{} có thể là string hoặc querier (LikeDocument)
	- max_doc_freq : int
	- max_query_terms : int
	- max_word_length : int
	- min_doc_freq : int
	- min_term_freq : int
	- min_word_length : int
	- minimum_should_match : string
	- _name : string
	- routing : string
	- stop_words : []string * [david,madam]
	- unlike : []interface{}
	- interface{} có thể là string hoặc querier (LikeDocument)
	- version : int64
	- version_type : string
*/
func ParseMoreLikeThisQuery(m fulltextsearch.Querier) *types.MoreLikeThisQuery {
	query := types.NewMoreLikeThisQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "analyzer":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Analyzer = &v
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "boost_terms":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.BoostTerms = &vFloat64
			case "fail_on_unsupported_field":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.FailOnUnsupportedField = &v
			case "fields":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]string, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						fields = append(fields, s)
					}
				}
				query.Fields = fields
			case "include":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Include = &v
			case "like":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Like = ParseLike(field)

			case "max_doc_freq":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxDocFreq = &v
			case "max_query_terms":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxQueryTerms = &v
			case "max_word_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxWordLength = &v
			case "min_doc_freq":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MinDocFreq = &v
			case "min_term_freq":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MinTermFreq = &v
			case "min_word_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MinWordLength = &v
			case "minimum_should_match":
				s, ok := value.(string)
				if !ok {
					break
				}
				query.MinimumShouldMatch = s
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "routing":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Routing = &v
			case "stop_words":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]string, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						fields = append(fields, s)
					}
				}
				query.StopWords = fields
			case "unlike":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Unlike = ParseLike(field)
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					break
				}
				query.Version = &v
			case "version_type":
				model := versiontype.VersionType{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.VersionType = &model
			}
		}
	}
	return query
}

/*
# like[index]

	index:
	  - string
	  - query(like document)
*/
func ParseLike(m fulltextsearch.Querier) []types.Like {
	likes := make([]types.Like, 0)
	pars, ok := m.GetParams().([]interface{})
	if !ok {
		return nil
	}
	for _, v := range pars {
		switch t := v.(type) {
		case fulltextsearch.Query:
			doc := ParseLikeDocument(&t)
			likes = append(likes, doc)
		case string:
			likes = append(likes, t)
		}
	}
	return likes
}

/*
{...parameters}
*parameters:
  - doc : string (json.RawMessage)
  - fields : []string *fields[nam,nu]
  - _id : string
  - _index : string
  - per_field_analyzer : map[string]string *per_field_analyzer{field:ba,fell:bon}
  - version : int64
  - version_type : string
*/
func ParseLikeDocument(m fulltextsearch.Querier) types.LikeDocument {
	query := types.LikeDocument{}
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "doc":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Doc = json.RawMessage(v)
			case "fields":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]string, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						fields = append(fields, s)
					}
				}
				query.Fields = fields
			case "_id":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Id_ = &v
			case "_index":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Index_ = &v
			case "per_field_analyzer":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options := make(map[string]string)
				pars, ok := field.GetParams().(map[string]interface{})
				if !ok {
					break
				}
				for k, v := range pars {
					s, ok := v.(string)
					if !ok {
						break
					}
					options[k] = s
				}
				query.PerFieldAnalyzer = options
			case "routing":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Routing = &v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					break
				}
				query.Version = &v
			case "version_type":
				model := versiontype.VersionType{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.VersionType = &model
			}
		}
	}
	return query
}

/*
percolate{...parameters}
*parameters:
  - boost : float32
  - document : string (json.RawMessage)
  - documents : []string(json.RawMessage) *documents[nam,nu]
  - id : string
  - field : string
  - index : string
  - name : string
  - preference : string
  - _name : string
  - version : int64
  - version_type : string
*/
func ParsePercolateQuery(m fulltextsearch.Querier) *types.PercolateQuery {
	query := types.NewPercolateQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "document":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Document = json.RawMessage(v)
			case "documents":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]json.RawMessage, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						doc := json.RawMessage(s)
						fields = append(fields, doc)
					}
				}
				query.Documents = fields
			case "id":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Id = &v
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			case "index":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Index = &v
			case "name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Name = &v
			case "preference":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Preference = &v
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "routing":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Routing = &v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					break
				}
				query.Version = &v
			}
		}
	}
	return query
}

/*
rank_feature{...parameters}

	parameters:
	- boost : float32
	- field : string
	- linear : query linear{}
	- log : query *log{scaling_factor:float32}
	- _name : string
	- saturation : query *saturation{pivot:float32}
	- sigmoid : query *sigmoid{pivot:float32,exponent:float32}
*/
func ParseRankFeatureQuery(m fulltextsearch.Querier) *types.RankFeatureQuery {
	query := types.NewRankFeatureQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			case "linear":
				_, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Linear = types.NewRankFeatureFunctionLinear()
			case "log":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Log = ParseLog(field)
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "saturation":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Saturation = ParseSaturation(field)
			case "sigmoid":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Sigmoid = ParseSigmoid(field)
			}
		}
	}
	return query
}

func ParseLog(m fulltextsearch.Querier) *types.RankFeatureFunctionLogarithm {
	query := types.NewRankFeatureFunctionLogarithm()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "scaling_factor":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.ScalingFactor = vFloat32
			}
		}
	}
	return query
}

func ParseSaturation(m fulltextsearch.Querier) *types.RankFeatureFunctionSaturation {
	query := types.NewRankFeatureFunctionSaturation()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "pivot":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Pivot = &vFloat32
			}
		}
	}
	return query
}
func ParseSigmoid(m fulltextsearch.Querier) *types.RankFeatureFunctionSigmoid {
	query := types.NewRankFeatureFunctionSigmoid()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "pivot":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Pivot = vFloat32
			case "exponent":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Exponent = vFloat32
			}
		}
	}
	return query
}

/*
script{...parameters}

	parameters:
	- boost : float32
	- _name : string
	- script : query *script{id:string,...}
*/
func ParseScriptQuery(m fulltextsearch.Querier) *types.ScriptQuery {
	query := types.NewScriptQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v

			case "script":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Script = *ParseScript(field)
			}
		}
	}
	return query
}

/*
script_score{...parameters}

	 parameters:
	- boost : float32
	- min_score : float32
	- query : query
	- _name : string
	- script : query *script{id:string,...}
*/
func ParseScriptScoreQuery(m fulltextsearch.Querier) *types.ScriptScoreQuery {
	query := types.NewScriptScoreQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "min_score":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.MinScore = &vFloat32
			case "query":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Query = ParseQueryToSearch(field)
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "script":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Script = *ParseScript(field)
			}
		}
	}
	return query
}

/*
script_score{...parameters}

  parameters:
  - boost : float32
  - query : string
  - _name : string
*/

func ParseWrapperQuery(m fulltextsearch.Querier) *types.WrapperQuery {
	query := types.NewWrapperQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "query":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Query = v
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			}
		}
	}
	return query
}

/*
pinned{...parameters}
*parameters:
  - boost : float32
  - docs : []query (PinnedDoc)
  - ids : []string
  - organic: query
  - _name : string
*/

func ParsePinnedQuery(m fulltextsearch.Querier) *types.PinnedQuery {
	query := types.NewPinnedQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "ids":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]string, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						fields = append(fields, s)
					}
				}
				query.Ids = fields
			case "docs":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]types.PinnedDoc, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					quies, ok := v.(fulltextsearch.Querier)
					if ok {
						fields = append(fields, ParsePinnedDoc(quies))
					}
				}
				query.Docs = fields
			case "organic":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Organic = ParseQueryToSearch(field)
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			}
		}
	}
	return query
}

func ParsePinnedDoc(m fulltextsearch.Querier) types.PinnedDoc {
	query := types.NewPinnedDoc()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "_id":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Id_ = v
			case "_index":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Index_ = v
			}
		}
	}
	return *query
}

/*
rule{...parameters}

  parameters:
  - boost : float32
  - match_criteria : string (json.RawMessage)
  - ruleset_ids : []string
  - organic: query
  - _name : string
*/

func ParseRuleQuery(m fulltextsearch.Querier) *types.RuleQuery {
	query := types.NewRuleQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "match_criteria":
				v, ok := value.(string)
				if !ok {
					break
				}
				j := json.RawMessage(v)
				query.MatchCriteria = j
			case "organic":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Organic = ParseQueryToSearch(field)
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "ruleset_ids":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]string, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						fields = append(fields, s)
					}
				}
				query.RulesetIds = fields
			}
		}
	}
	return query
}
