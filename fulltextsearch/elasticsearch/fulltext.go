package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldszeroterms"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/simplequerystringflag"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
	"github.com/luongduc1246/ultility/reqparams"
)

/*
Phân tích câu query intervals
câu query có dạng intervals{fields{all_of{...}}}
*/
func ParseIntervalsQuery(m reqparams.Querier) map[string]types.IntervalsQuery {
	intervals := make(map[string]types.IntervalsQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewIntervalsQuery()
			options, ok := value.(reqparams.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {
				case "all_of":
					allOf, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.AllOf = ParseAllOfQuery(allOf)
				case "any_of":
					anyOf, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.AnyOf = ParseAnyOfQuery(anyOf)
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
				case "fuzzy":
					fuzzy, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.Fuzzy = ParseIntervalsFuzzyQuery(fuzzy)
				case "match":
					match, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.Match = ParseIntervalsMatchQuery(match)
				case "prefix":
					prefix, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.Prefix = ParseIntervalsPrefixQuery(prefix)
				case "wildcard":
					wildcard, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.Wildcard = ParseIntervalsWildcardQuery(wildcard)
				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v
				}
			}
			intervals[string(key)] = *query
		}
	default:
		return nil
	}
	return intervals
}

/*
Phân tích câu query all_of
câu query có all_of{filter{...},intervals[{...},{...}]}
*/
func ParseAllOfQuery(m reqparams.Querier) *types.IntervalsAllOf {
	allOf := types.NewIntervalsAllOf()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "filter":
				filter, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				allOf.Filter = ParseIntervalsFilter(filter)
			case "max_gaps":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				allOf.MaxGaps = &v
			case "intervals":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceIntervals := []types.Intervals{}
				for _, query := range options {
					v, ok := query.(reqparams.Querier)
					if !ok {
						break
					}
					i := ParseIntervals(v)
					sliceIntervals = append(sliceIntervals, *i)
				}
				allOf.Intervals = sliceIntervals
			case "ordered":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				allOf.Ordered = &v
			}
		}
	}
	return allOf
}

/*
Phân tích câu query intervals trong all_of,after,...

	...{all_of{...}}
*/
func ParseIntervals(m reqparams.Querier) *types.Intervals {
	query := types.NewIntervals()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "all_of":
				allOf, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.AllOf = ParseAllOfQuery(allOf)
			case "any_of":
				anyOf, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.AnyOf = ParseAnyOfQuery(anyOf)
			case "fuzzy":
				fuzzy, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Fuzzy = ParseIntervalsFuzzyQuery(fuzzy)
			case "match":
				match, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Match = ParseIntervalsMatchQuery(match)
			case "prefix":
				prefix, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Prefix = ParseIntervalsPrefixQuery(prefix)
			case "wildcard":
				wildcard, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Wildcard = ParseIntervalsWildcardQuery(wildcard)
			}
		}
	}
	return query
}

/*
Phân tích câu filter nằm trong macth của intervals,all_of...

	...{after{all_of{...},...}}
*/
func ParseIntervalsFilter(m reqparams.Querier) *types.IntervalsFilter {
	filter := types.NewIntervalsFilter()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "after":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.After = ParseIntervals(field)
			case "before":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.Before = ParseIntervals(field)
			case "contained_by":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.ContainedBy = ParseIntervals(field)
			case "containing":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.Containing = ParseIntervals(field)
			case "not_contained_by":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.NotContainedBy = ParseIntervals(field)
			case "not_containing":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.NotContaining = ParseIntervals(field)
			case "not_overlapping":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.NotOverlapping = ParseIntervals(field)
			case "overlapping":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.Overlapping = ParseIntervals(field)
			case "script":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				filter.Script = ParseScript(field)
			}
		}
	}
	return filter
}

/*
Phân tích câu query script
câu query có dạng ...{id:3,lang:vn,...}
*/
func ParseScript(m reqparams.Querier) *types.Script {
	query := types.NewScript()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "id":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Id = &v
			case "lang":
				v, ok := value.(string)
				if !ok {
					break
				}
				lang := scriptlanguage.ScriptLanguage{
					Name: v,
				}
				query.Lang = &lang
			case "options":
				field, ok := value.(reqparams.Querier)
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
				query.Options = options
			case "params":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				ps := make(map[string]json.RawMessage)
				pars, ok := field.GetParams().(map[string]interface{})
				if !ok {
					break
				}
				for k, v := range pars {
					s, ok := v.(string)
					if !ok {
						break
					}
					ps[k] = json.RawMessage(s)
				}
				query.Params = ps
			case "source":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Source = &v

			}
		}
	}
	return query
}

/*
Phân tích câu query any_of

	...{filter{}}
*/
func ParseAnyOfQuery(m reqparams.Querier) *types.IntervalsAnyOf {
	anyOf := types.NewIntervalsAnyOf()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "filter":
				filter, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				anyOf.Filter = ParseIntervalsFilter(filter)

			case "intervals":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceIntervals := []types.Intervals{}
				for _, query := range options {
					v, ok := query.(reqparams.Querier)
					if !ok {
						break
					}
					i := ParseIntervals(v)
					sliceIntervals = append(sliceIntervals, *i)
				}
				anyOf.Intervals = sliceIntervals
			}
		}
	}
	return anyOf
}

/*
Phân tích câu query fuzzy
câu query có dạng ...{...}
*/
func ParseIntervalsFuzzyQuery(m reqparams.Querier) *types.IntervalsFuzzy {
	fuzzy := types.NewIntervalsFuzzy()
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
				fuzzy.Analyzer = &v
			case "fuzziness":
				v, ok := value.(types.Fuzziness)
				if !ok {
					return nil
				}
				fuzzy.Fuzziness = v
			case "prefix_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				fuzzy.PrefixLength = &v
			case "term":
				v, ok := value.(string)
				if !ok {
					break
				}
				fuzzy.Term = v
			case "transpositions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				fuzzy.Transpositions = &v
			case "use_field":
				v, ok := value.(string)
				if !ok {
					break
				}
				fuzzy.UseField = &v
			}
		}
	}
	return fuzzy
}

/*
Phân tích câu query match trong intervals

	...{...}
*/
func ParseIntervalsMatchQuery(m reqparams.Querier) *types.IntervalsMatch {
	query := types.NewIntervalsMatch()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "analyzer":
				s, ok := value.(string)
				if !ok {
					break
				}
				query.Analyzer = &s
			case "filter":
				filter, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Filter = ParseIntervalsFilter(filter)
			case "max_gaps":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxGaps = &v
			case "ordered":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Ordered = &v
			case "query":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Query = v
			case "use_field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.UseField = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query prefix trong intervals

	...{...}
*/
func ParseIntervalsPrefixQuery(m reqparams.Querier) *types.IntervalsPrefix {
	query := types.NewIntervalsPrefix()
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
			case "prefix":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Prefix = v
			case "use_field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.UseField = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query wildcard trong intervals

	...{...}
*/
func ParseIntervalsWildcardQuery(m reqparams.Querier) *types.IntervalsWildcard {
	query := types.NewIntervalsWildcard()
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
			case "pattern":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Pattern = v
			case "use_field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.UseField = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query match
query có dạng match{anylyzer:true,...}
*/
func ParseMatchQuery(m reqparams.Querier) map[string]types.MatchQuery {
	matchQuery := make(map[string]types.MatchQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewMatchQuery()
			options, ok := value.(reqparams.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {
				case "analyzer":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Analyzer = &v
				case "auto_generate_synonyms_phrase_query":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.AutoGenerateSynonymsPhraseQuery = &v
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
				case "cutoff_frequency":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseFloat(s, 64)
					if err != nil {
						break
					}
					vFloat64 := types.Float64(v)
					query.CutoffFrequency = &vFloat64
				case "fuzziness":
					s, ok := value.(string)
					if !ok {
						break
					}
					v := types.Fuzziness(s)
					query.Fuzziness = &v
				case "fuzzy_rewrite":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.FuzzyRewrite = &v
				case "fuzzy_transpositions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.FuzzyTranspositions = &v
				case "lenient":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.Lenient = &v
				case "max_expansions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxExpansions = &v
				case "minimum_should_match":
					query.MinimumShouldMatch = value
				case "operator":
					op := operator.Operator{}
					v, ok := value.(string)
					if !ok {
						break
					}
					op.Name = v
					query.Operator = &op
				case "prefix_length":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.PrefixLength = &v
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
				case "zero_terms_query":
					zero := zerotermsquery.ZeroTermsQuery{}
					v, ok := value.(string)
					if !ok {
						break
					}
					zero.Name = v
					query.ZeroTermsQuery = &zero
				}
			}
			matchQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return matchQuery
}

/*
Phân tích câu query match_phrase
query có dạng match_phrase{anylyzer:true}
*/
func ParseMatchPhraseQuery(m reqparams.Querier) map[string]types.MatchPhraseQuery {
	matchQuery := make(map[string]types.MatchPhraseQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewMatchPhraseQuery()
			options, ok := value.(reqparams.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {
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
				case "zero_terms_query":
					zero := zerotermsquery.ZeroTermsQuery{}
					v, ok := value.(string)
					if !ok {
						break
					}
					zero.Name = v
					query.ZeroTermsQuery = &zero
				case "slop":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.Slop = &v
				}
			}
			matchQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return matchQuery
}

/*
Phân tích câu query match_phrase_prefix
query có dạng match_phrase_prefix{anylyzer:true}
*/
func ParseMatchPhrasePrefixQuery(m reqparams.Querier) map[string]types.MatchPhrasePrefixQuery {
	matchQuery := make(map[string]types.MatchPhrasePrefixQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewMatchPhrasePrefixQuery()
			options, ok := value.(reqparams.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {
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
				case "max_expansions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxExpansions = &v
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
				case "zero_terms_query":
					zero := zerotermsquery.ZeroTermsQuery{}
					v, ok := value.(string)
					if !ok {
						break
					}
					zero.Name = v
					query.ZeroTermsQuery = &zero
				case "slop":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.Slop = &v
				}
			}
			matchQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return matchQuery
}

/*
Phân tích câu query match_bool_prefix
query có dạng match_bool_prefix{anylyzer:true}
*/
func ParseMatchBoolPrefixQuery(m reqparams.Querier) map[string]types.MatchBoolPrefixQuery {
	matchQuery := make(map[string]types.MatchBoolPrefixQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewMatchBoolPrefixQuery()
			options, ok := value.(reqparams.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})

			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {
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

				case "fuzziness":
					s, ok := value.(string)
					if !ok {
						break
					}
					v := types.Fuzziness(s)
					query.Fuzziness = &v
				case "fuzzy_rewrite":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.FuzzyRewrite = &v
				case "fuzzy_transpositions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.FuzzyTranspositions = &v

				case "max_expansions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxExpansions = &v
				case "minimum_should_match":
					v := value.(types.MinimumShouldMatch)
					query.MinimumShouldMatch = &v
				case "operator":
					op := operator.Operator{}
					s, ok := value.(string)
					if !ok {
						break
					}
					op.Name = s
					query.Operator = &op
				case "prefix_length":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.PrefixLength = &v
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
			matchQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return matchQuery
}

/*
Phân tích câu query combined_fields

	combined_fields{boost:3,fields[a,b,d]}
*/
func ParseCombinedFieldsQuery(m reqparams.Querier) *types.CombinedFieldsQuery {
	query := types.NewCombinedFieldsQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {

			case "auto_generate_synonyms_phrase_query":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AutoGenerateSynonymsPhraseQuery = &v
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
			case "minimum_should_match":
				query.MinimumShouldMatch = value
			case "operator":
				op := combinedfieldsoperator.CombinedFieldsOperator{}
				v, ok := value.(string)
				if !ok {
					break
				}
				op.Name = v
				query.Operator = &op
			case "fields":
				field, ok := value.(reqparams.Querier)
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
			case "zero_terms_query":
				zero := combinedfieldszeroterms.CombinedFieldsZeroTerms{}
				v, ok := value.(string)
				if !ok {
					break
				}
				zero.Name = v
				query.ZeroTermsQuery = &zero
			}

		}
	}

	return query
}

/*
Phân tích câu query multi_match

	multi_match{boost:3,fields=[a,b,d]}
*/
func ParseMultiMatchQuery(m reqparams.Querier) *types.MultiMatchQuery {
	query := types.NewMultiMatchQuery()
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
			case "auto_generate_synonyms_phrase_query":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AutoGenerateSynonymsPhraseQuery = &v
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
			case "cutoff_frequency":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.CutoffFrequency = &vFloat64
			case "fields":
				field, ok := value.(reqparams.Querier)
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
			case "fuzziness":
				query.Fuzziness = value
			case "fuzzy_rewrite":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.FuzzyRewrite = &v
			case "fuzzy_transpositions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.FuzzyTranspositions = &v
			case "lenient":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Lenient = &v
			case "max_expansions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxExpansions = &v
			case "minimum_should_match":
				query.MinimumShouldMatch = value
			case "operator":
				op := operator.Operator{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := string(s)
				op.Name = v
				query.Operator = &op
			case "prefix_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.PrefixLength = &v
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
			case "slop":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.Slop = &v
			case "tie_breaker":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.TieBreaker = &vFloat64
			case "type":
				t := textquerytype.TextQueryType{}
				v, ok := value.(string)
				if !ok {
					break
				}
				t.Name = v
				query.Type = &t
			case "zero_terms_query":
				zero := zerotermsquery.ZeroTermsQuery{}
				v, ok := value.(string)
				if !ok {
					break
				}
				zero.Name = v
				query.ZeroTermsQuery = &zero
			}
		}
	}
	return query
}

/*
Phân tích câu query query_string
câu query có dạng query_string{boost:3,fields[a,b,d]}
*/
func ParseQueryStringQuery(m reqparams.Querier) *types.QueryStringQuery {
	query := types.NewQueryStringQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "allow_leading_wildcard":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AllowLeadingWildcard = &v
			case "analyze_wildcard":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AnalyzeWildcard = &v
			case "analyzer":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Analyzer = &v
			case "auto_generate_synonyms_phrase_query":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AutoGenerateSynonymsPhraseQuery = &v
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
			case "default_field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.DefaultField = &v
			case "default_operator":
				op := operator.Operator{}
				v, ok := value.(string)
				if !ok {
					break
				}
				op.Name = v
				query.DefaultOperator = &op
			case "enable_position_increments":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.EnablePositionIncrements = &v
			case "escape":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Escape = &v
			case "fields":
				field, ok := value.(reqparams.Querier)
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
			case "fuzziness":
				query.Fuzziness = value
			case "fuzzy_max_expansions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.FuzzyMaxExpansions = &v
			case "fuzzy_prefix_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.FuzzyPrefixLength = &v
			case "fuzzy_rewrite":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.FuzzyRewrite = &v
			case "fuzzy_transpositions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.FuzzyTranspositions = &v
			case "lenient":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Lenient = &v
			case "max_determinized_states":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxDeterminizedStates = &v

			case "minimum_should_match":
				query.MinimumShouldMatch = value
			case "phrase_slop":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.PhraseSlop = &vFloat64
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
			case "quote_analyzer":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QuoteAnalyzer = &v
			case "quote_field_suffix":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QuoteFieldSuffix = &v
			case "rewrite":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Rewrite = &v

			case "tie_breaker":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.TieBreaker = &vFloat64
			case "time_zone":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.TimeZone = &v
			case "type":
				t := textquerytype.TextQueryType{}
				v, ok := value.(string)
				if !ok {
					break
				}
				t.Name = v
				query.Type = &t

			}
		}
	}
	return query
}

/*
Phân tích câu query simple_query_string

	simple_query_string{boost=3,...}
*/
func ParseSimpleQueryStringQuery(m reqparams.Querier) *types.SimpleQueryStringQuery {
	query := types.NewSimpleQueryStringQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "analyze_wildcard":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AnalyzeWildcard = &v
			case "analyzer":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Analyzer = &v
			case "auto_generate_synonyms_phrase_query":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.AutoGenerateSynonymsPhraseQuery = &v
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
			case "default_operator":
				op := operator.Operator{}
				v, ok := value.(string)
				if !ok {
					break
				}
				op.Name = v
				query.DefaultOperator = &op

			case "fields":
				field, ok := value.(reqparams.Querier)
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

			case "fuzzy_max_expansions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.FuzzyMaxExpansions = &v
			case "fuzzy_prefix_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.FuzzyPrefixLength = &v

			case "fuzzy_transpositions":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.FuzzyTranspositions = &v
			case "lenient":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Lenient = &v
			case "minimum_should_match":

				query.MinimumShouldMatch = value

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

			case "quote_field_suffix":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QuoteFieldSuffix = &v
			case "flags":
				flags := simplequerystringflag.SimpleQueryStringFlag{}
				v, ok := value.(string)
				if !ok {
					break
				}
				flags.Name = v
				query.Flags = flags
			}
		}
	}
	return query
}
