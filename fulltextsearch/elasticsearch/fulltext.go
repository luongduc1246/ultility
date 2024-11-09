package elasticsearch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldszeroterms"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

/*
Phân tích câu query intervals
câu query có dạng intervals[fields[all_of[...]]]
*/
func ParseIntervalsQuery(m fulltextsearch.Querier) map[string]types.IntervalsQuery {
	intervals := make(map[string]types.IntervalsQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewIntervalsQuery()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.ALLOF:
					allOf := value.(fulltextsearch.Querier)
					query.AllOf = ParseAllOfQuery(allOf)
				case fulltextsearch.ANYOF:
					anyOf := value.(fulltextsearch.Querier)
					query.AnyOf = ParseAnyOfQuery(anyOf)
				case fulltextsearch.BOOST:
					v := float32(value.(fulltextsearch.Boost))
					query.Boost = &v
				case fulltextsearch.FUZZY:
					fuzzy := value.(fulltextsearch.Querier)
					query.Fuzzy = ParseIntervalsFuzzyQuery(fuzzy)
				case fulltextsearch.MATCH:
					match := value.(fulltextsearch.Querier)
					query.Match = ParseIntervalsMatchQuery(match)
				case fulltextsearch.PREFIX:
					prefix := value.(fulltextsearch.Querier)
					query.Prefix = ParseIntervalsPrefixQuery(prefix)
				case fulltextsearch.WILDCARD:
					wildcard := value.(fulltextsearch.Querier)
					query.Wildcard = ParseIntervalsWildcardQuery(wildcard)
				case fulltextsearch.QUERYNAME:
					v := string(value.(fulltextsearch.QueryName))
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
câu query có dạng intervals[fields[all_of[intervals_filter[...],slice_intervals[intervals[...]]]]]
  - intervals_filter thay cho filter
  - slice_intervals thay cho mảng intervals
*/
func ParseAllOfQuery(m fulltextsearch.Querier) *types.IntervalsAllOf {
	allOf := types.NewIntervalsAllOf()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.FILTER:
				filter := value.(fulltextsearch.Querier)
				allOf.Filter = ParseIntervalsFilter(filter)
			case fulltextsearch.MAXGAPS:
				v := int(value.(fulltextsearch.MaxGaps))
				allOf.MaxGaps = &v
			case fulltextsearch.SLICEINTERVALS:
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().([]fulltextsearch.Querier)
				sliceIntervals := []types.Intervals{}
				for _, query := range options {
					i := ParseIntervals(query)
					sliceIntervals = append(sliceIntervals, *i)
				}
				allOf.Intervals = sliceIntervals
			case fulltextsearch.ORDERED:
				v := bool(value.(fulltextsearch.Ordered))
				allOf.Ordered = &v
			}
		}
	}
	return allOf
}

/*
Phân tích câu query intervals trong all_off
câu query có dạng ...[all_of[slice_intervals[intervals[all_of[...],any_of[...],...]]]]]
*/
func ParseIntervals(m fulltextsearch.Querier) *types.Intervals {
	query := types.NewIntervals()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ALLOF:
				allOf := value.(fulltextsearch.Querier)
				query.AllOf = ParseAllOfQuery(allOf)
			case fulltextsearch.ANYOF:
				anyOf := value.(fulltextsearch.Querier)
				query.AnyOf = ParseAnyOfQuery(anyOf)
			case fulltextsearch.FUZZY:
				fuzzy := value.(fulltextsearch.Querier)
				query.Fuzzy = ParseIntervalsFuzzyQuery(fuzzy)
			case fulltextsearch.MATCH:
				match := value.(fulltextsearch.Querier)
				query.Match = ParseIntervalsMatchQuery(match)
			case fulltextsearch.PREFIX:
				prefix := value.(fulltextsearch.Querier)
				query.Prefix = ParseIntervalsPrefixQuery(prefix)
			case fulltextsearch.WILDCARD:
				wildcard := value.(fulltextsearch.Querier)
				query.Wildcard = ParseIntervalsWildcardQuery(wildcard)
			}
		}
	}
	return query
}

/*
Phân tích câu query intervals_filter nằm trong macth của intervals,all_of...
câu query có dạng intervals_filter[after[...],script[id=3]]
*/
func ParseIntervalsFilter(m fulltextsearch.Querier) *types.IntervalsFilter {
	filter := types.NewIntervalsFilter()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.AFTER:
				field := value.(fulltextsearch.Querier)
				filter.After = ParseIntervals(field)
			case fulltextsearch.BEFORE:
				field := value.(fulltextsearch.Querier)
				filter.Before = ParseIntervals(field)
			case fulltextsearch.CONTAINEDBY:
				field := value.(fulltextsearch.Querier)
				filter.ContainedBy = ParseIntervals(field)
			case fulltextsearch.CONTAINING:
				field := value.(fulltextsearch.Querier)
				filter.Containing = ParseIntervals(field)
			case fulltextsearch.NOTCONTAINEDBY:
				field := value.(fulltextsearch.Querier)
				filter.NotContainedBy = ParseIntervals(field)
			case fulltextsearch.NOTCONTAINING:
				field := value.(fulltextsearch.Querier)
				filter.NotContaining = ParseIntervals(field)
			case fulltextsearch.NOTOVERLAPPING:
				field := value.(fulltextsearch.Querier)
				filter.NotOverlapping = ParseIntervals(field)
			case fulltextsearch.SCRIPT:
				field := value.(fulltextsearch.Querier)
				filter.Script = ParseScript(field)
			}
		}
	}
	return filter
}

/*
Phân tích câu query script
câu query có dạng script[id=3,lang=vn,...]
*/
func ParseScript(m fulltextsearch.Querier) *types.Script {
	query := types.NewScript()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ID:
				v := string(value.(fulltextsearch.Id))
				query.Id = &v
			case fulltextsearch.LANG:
				v := string(value.(fulltextsearch.Lang))
				lang := scriptlanguage.ScriptLanguage{
					Name: v,
				}
				query.Lang = &lang
			case fulltextsearch.OPTIONS:
				field := value.(fulltextsearch.Querier)
				options := make(map[string]string)
				pars := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				for k, v := range pars {
					strKey := fmt.Sprintf("%v", k)
					strValue := fmt.Sprintf("%v", v)
					options[strKey] = strValue
				}
				query.Options = options
			case fulltextsearch.PARAMS:
				field := value.(fulltextsearch.Querier)
				ps := make(map[string]json.RawMessage)
				pars := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				for k, v := range pars {
					strKey := fmt.Sprintf("%v", k)
					strValue := fmt.Sprintf("%v", v)
					ps[strKey] = json.RawMessage(strValue)
				}
				query.Params = ps
			case fulltextsearch.SOURCE:
				v := string(value.(fulltextsearch.Source))
				query.Source = &v

			}
		}
	}
	return query
}

/*
Phân tích câu query any_of
câu query có dạng intervals[fields[any_of[intervals_filter[...],slice_intervals[intervals[...]]]]]
  - intervals_filter thay cho filter
  - slice_intervals thay cho mảng intervals
*/
func ParseAnyOfQuery(m fulltextsearch.Querier) *types.IntervalsAnyOf {
	anyOf := types.NewIntervalsAnyOf()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.INTERVALSFILTER:
				filter := value.(fulltextsearch.Querier)
				anyOf.Filter = ParseIntervalsFilter(filter)

			case fulltextsearch.SLICEINTERVALS:
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().([]fulltextsearch.Querier)
				sliceIntervals := []types.Intervals{}
				for _, query := range options {
					i := ParseIntervals(query)
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
câu query có dạng intervals[fields[fuzzy[term_string=test,analyzer=atest]]]]
  - term_string thay cho term
*/
func ParseIntervalsFuzzyQuery(m fulltextsearch.Querier) *types.IntervalsFuzzy {
	fuzzy := types.NewIntervalsFuzzy()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				fuzzy.Analyzer = &v
			case fulltextsearch.FUZZINESS:
				v, ok := value.(types.Fuzziness)
				if !ok {
					return nil
				}
				fuzzy.Fuzziness = v
			case fulltextsearch.PREFIXLENGTH:
				v := int(value.(fulltextsearch.PrefixLength))
				fuzzy.PrefixLength = &v
			case fulltextsearch.TERMSTRING:
				v := string(value.(fulltextsearch.TermString))
				fuzzy.Term = v
			case fulltextsearch.TRANSPOSITIONS:
				v := bool(value.(fulltextsearch.Transpositions))
				fuzzy.Transpositions = &v
			case fulltextsearch.USEFIELD:
				v := string(value.(fulltextsearch.UseField))
				fuzzy.UseField = &v
			}
		}
	}
	return fuzzy
}

/*
Phân tích câu query match trong intervals
câu query có dạng intervals[fields[match[analyzer=atest,intervals_filter[after[...],script[id=3]]]]]]
  - intervals_filter thay the cho filter
*/
func ParseIntervalsMatchQuery(m fulltextsearch.Querier) *types.IntervalsMatch {
	query := types.NewIntervalsMatch()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				query.Analyzer = &v
			case fulltextsearch.INTERVALSFILTER:
				filter := value.(fulltextsearch.Querier)
				query.Filter = ParseIntervalsFilter(filter)
			case fulltextsearch.MAXGAPS:
				v := int(value.(fulltextsearch.MaxGaps))
				query.MaxGaps = &v
			case fulltextsearch.ORDERED:
				v := bool(value.(fulltextsearch.Ordered))
				query.Ordered = &v
			case fulltextsearch.QUERY:
				v := string(value.(fulltextsearch.Query))
				query.Query = v
			case fulltextsearch.USEFIELD:
				v := string(value.(fulltextsearch.UseField))
				query.UseField = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query prefix trong intervals
query có dạng prefix[prefix_string=test]
  - prefix_string thay thế cho prefix trong params
*/
func ParseIntervalsPrefixQuery(m fulltextsearch.Querier) *types.IntervalsPrefix {
	query := types.NewIntervalsPrefix()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				query.Analyzer = &v
			case fulltextsearch.PREFIXSTRING:
				v := string(value.(fulltextsearch.PrefixString))
				query.Prefix = v
			case fulltextsearch.USEFIELD:
				v := string(value.(fulltextsearch.UseField))
				query.UseField = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query wildcard trong intervals
query có dạng wildcard[anylyzer=test]
*/
func ParseIntervalsWildcardQuery(m fulltextsearch.Querier) *types.IntervalsWildcard {
	query := types.NewIntervalsWildcard()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				query.Analyzer = &v
			case fulltextsearch.PATTERN:
				v := string(value.(fulltextsearch.Pattern))
				query.Pattern = v
			case fulltextsearch.USEFIELD:
				v := string(value.(fulltextsearch.UseField))
				query.UseField = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query match
query có dạng match[anylyzer=true]
*/
func ParseMatchQuery(m fulltextsearch.Querier) map[string]types.MatchQuery {
	matchQuery := make(map[string]types.MatchQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewMatchQuery()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.ANALYZER:
					v := string(value.(fulltextsearch.Analyzer))
					query.Analyzer = &v
				case fulltextsearch.AUTOGENERATESYNONYMSPHRASEQUERY:
					v := bool(value.(fulltextsearch.AutoGenerateSynonymsPhraseQuery))
					query.AutoGenerateSynonymsPhraseQuery = &v
				case fulltextsearch.BOOST:
					v := float32(value.(fulltextsearch.Boost))
					query.Boost = &v
				case fulltextsearch.CUTOFFFREQUENCY:
					v := types.Float64(value.(fulltextsearch.CutoffFrequency))
					query.CutoffFrequency = &v
				case fulltextsearch.FUZZINESS:
					v := types.Fuzziness(value.(fulltextsearch.Fuzziness))
					query.Fuzziness = &v
				case fulltextsearch.FUZZYREWRITE:
					v := string(value.(fulltextsearch.FuzzyRewrite))
					query.FuzzyRewrite = &v
				case fulltextsearch.FUZZYTRANSPOSITIONS:
					v := bool(value.(fulltextsearch.FuzzyTranspositions))
					query.FuzzyTranspositions = &v
				case fulltextsearch.LENIENT:
					v := bool(value.(fulltextsearch.Lenient))
					query.Lenient = &v
				case fulltextsearch.MAXEXPANSIONS:
					v := int(value.(fulltextsearch.MaxExpansions))
					query.MaxExpansions = &v
				case fulltextsearch.MINIMUMSHOULDMATCH:
					v := types.MinimumShouldMatch(value.(fulltextsearch.MinimumShouldMatch))
					query.MinimumShouldMatch = &v
				case fulltextsearch.OPERATOR:
					op := operator.Operator{}
					v := string(value.(fulltextsearch.Operator))
					op.Name = v
					query.Operator = &op
				case fulltextsearch.PREFIXLENGTH:
					v := int(value.(fulltextsearch.PrefixLength))
					query.PrefixLength = &v
				case fulltextsearch.QUERY:
					v := string(value.(fulltextsearch.Query))
					query.Query = v
				case fulltextsearch.QUERYNAME:
					v := string(value.(fulltextsearch.QueryName))
					query.QueryName_ = &v
				case fulltextsearch.ZEROTERMSQUERY:
					zero := zerotermsquery.ZeroTermsQuery{}
					v := string(value.(fulltextsearch.ZeroTermsQuery))
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
query có dạng match_phrase[anylyzer=true]
*/
func ParseMatchPhraseQuery(m fulltextsearch.Querier) map[string]types.MatchPhraseQuery {
	matchQuery := make(map[string]types.MatchPhraseQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewMatchPhraseQuery()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.ANALYZER:
					v := string(value.(fulltextsearch.Analyzer))
					query.Analyzer = &v
				case fulltextsearch.BOOST:
					v := float32(value.(fulltextsearch.Boost))
					query.Boost = &v
				case fulltextsearch.QUERY:
					v := string(value.(fulltextsearch.Query))
					query.Query = v
				case fulltextsearch.QUERYNAME:
					v := string(value.(fulltextsearch.QueryName))
					query.QueryName_ = &v
				case fulltextsearch.SLOP:
					v := int(value.(fulltextsearch.Slop))
					query.Slop = &v
				case fulltextsearch.ZEROTERMSQUERY:
					zero := zerotermsquery.ZeroTermsQuery{}
					v := string(value.(fulltextsearch.ZeroTermsQuery))
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
Phân tích câu query match_phrase_prefix
query có dạng match_phrase_prefix[anylyzer=true]
*/
func ParseMatchPhrasePrefixQuery(m fulltextsearch.Querier) map[string]types.MatchPhrasePrefixQuery {
	matchQuery := make(map[string]types.MatchPhrasePrefixQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewMatchPhrasePrefixQuery()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.ANALYZER:
					v := string(value.(fulltextsearch.Analyzer))
					query.Analyzer = &v
				case fulltextsearch.BOOST:
					v := float32(value.(fulltextsearch.Boost))
					query.Boost = &v
				case fulltextsearch.QUERY:
					v := string(value.(fulltextsearch.Query))
					query.Query = v
				case fulltextsearch.QUERYNAME:
					v := string(value.(fulltextsearch.QueryName))
					query.QueryName_ = &v
				case fulltextsearch.SLOP:
					v := int(value.(fulltextsearch.Slop))
					query.Slop = &v
				case fulltextsearch.ZEROTERMSQUERY:
					zero := zerotermsquery.ZeroTermsQuery{}
					v := string(value.(fulltextsearch.ZeroTermsQuery))
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
Phân tích câu query match_bool_prefix
query có dạng match_bool_prefix[anylyzer=true]
*/
func ParseMatchBoolPrefixQuery(m fulltextsearch.Querier) map[string]types.MatchBoolPrefixQuery {
	matchQuery := make(map[string]types.MatchBoolPrefixQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewMatchBoolPrefixQuery()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.ANALYZER:
					v := string(value.(fulltextsearch.Analyzer))
					query.Analyzer = &v
				case fulltextsearch.BOOST:
					v := float32(value.(fulltextsearch.Boost))
					query.Boost = &v
				case fulltextsearch.FUZZINESS:
					v := types.Fuzziness(value.(fulltextsearch.Fuzziness))
					query.Fuzziness = &v
				case fulltextsearch.FUZZYREWRITE:
					v := string(value.(fulltextsearch.FuzzyRewrite))
					query.FuzzyRewrite = &v
				case fulltextsearch.FUZZYTRANSPOSITIONS:
					v := bool(value.(fulltextsearch.FuzzyTranspositions))
					query.FuzzyTranspositions = &v
				case fulltextsearch.MAXEXPANSIONS:
					v := int(value.(fulltextsearch.MaxExpansions))
					query.MaxExpansions = &v
				case fulltextsearch.MINIMUMSHOULDMATCH:
					v := types.MinimumShouldMatch(value.(fulltextsearch.MinimumShouldMatch))
					query.MinimumShouldMatch = &v
				case fulltextsearch.OPERATOR:
					op := operator.Operator{}
					v := string(value.(fulltextsearch.Operator))
					op.Name = v
					query.Operator = &op
				case fulltextsearch.PREFIXLENGTH:
					v := int(value.(fulltextsearch.PrefixLength))
					query.PrefixLength = &v
				case fulltextsearch.QUERY:
					v := string(value.(fulltextsearch.Query))
					query.Query = v
				case fulltextsearch.QUERYNAME:
					v := string(value.(fulltextsearch.QueryName))
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
câu query có dạng combined_fields[boost=3,fields=a;b;d]
*/
func ParseCombinedFieldsQuery(m fulltextsearch.Querier) *types.CombinedFieldsQuery {
	query := types.NewCombinedFieldsQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.AUTOGENERATESYNONYMSPHRASEQUERY:
				v := bool(value.(fulltextsearch.AutoGenerateSynonymsPhraseQuery))
				query.AutoGenerateSynonymsPhraseQuery = &v
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.MINIMUMSHOULDMATCH:
				v := types.MinimumShouldMatch(value.(fulltextsearch.MinimumShouldMatch))
				query.MinimumShouldMatch = &v
			case fulltextsearch.QUERY:
				v := string(value.(fulltextsearch.Query))
				query.Query = v
			case fulltextsearch.OPERATOR:
				op := combinedfieldsoperator.CombinedFieldsOperator{}
				v := string(value.(fulltextsearch.Operator))
				op.Name = v
				query.Operator = &op
			case fulltextsearch.FIELDS:
				v := []string(value.(fulltextsearch.Fields))
				query.Fields = v
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			case fulltextsearch.ZEROTERMSQUERY:
				zero := combinedfieldszeroterms.CombinedFieldsZeroTerms{}
				v := string(value.(fulltextsearch.ZeroTermsQuery))
				zero.Name = v
				query.ZeroTermsQuery = &zero
			}
		}
	}
	return query
}

/*
Phân tích câu query multi_match
câu query có dạng multi_match[boost=3,fields=a;b;d]
*/
func ParseMultiMatchQuery(m fulltextsearch.Querier) *types.MultiMatchQuery {
	query := types.NewMultiMatchQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				query.Analyzer = &v
			case fulltextsearch.AUTOGENERATESYNONYMSPHRASEQUERY:
				v := bool(value.(fulltextsearch.AutoGenerateSynonymsPhraseQuery))
				query.AutoGenerateSynonymsPhraseQuery = &v
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.CUTOFFFREQUENCY:
				v := types.Float64(value.(fulltextsearch.CutoffFrequency))
				query.CutoffFrequency = &v
			case fulltextsearch.FUZZINESS:
				v := types.Fuzziness(value.(fulltextsearch.Fuzziness))
				query.Fuzziness = &v
			case fulltextsearch.FUZZYREWRITE:
				v := string(value.(fulltextsearch.FuzzyRewrite))
				query.FuzzyRewrite = &v
			case fulltextsearch.FUZZYTRANSPOSITIONS:
				v := bool(value.(fulltextsearch.FuzzyTranspositions))
				query.FuzzyTranspositions = &v
			case fulltextsearch.LENIENT:
				v := bool(value.(fulltextsearch.Lenient))
				query.Lenient = &v
			case fulltextsearch.MAXEXPANSIONS:
				v := int(value.(fulltextsearch.MaxExpansions))
				query.MaxExpansions = &v
			case fulltextsearch.MINIMUMSHOULDMATCH:
				v := types.MinimumShouldMatch(value.(fulltextsearch.MinimumShouldMatch))
				query.MinimumShouldMatch = &v
			case fulltextsearch.OPERATOR:
				op := operator.Operator{}
				v := string(value.(fulltextsearch.Operator))
				op.Name = v
				query.Operator = &op
			case fulltextsearch.PREFIXLENGTH:
				v := int(value.(fulltextsearch.PrefixLength))
				query.PrefixLength = &v
			case fulltextsearch.QUERY:
				v := string(value.(fulltextsearch.Query))
				query.Query = v
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			case fulltextsearch.ZEROTERMSQUERY:
				zero := zerotermsquery.ZeroTermsQuery{}
				v := string(value.(fulltextsearch.ZeroTermsQuery))
				zero.Name = v
				query.ZeroTermsQuery = &zero
			case fulltextsearch.SLOP:
				v := int(value.(fulltextsearch.Slop))
				query.Slop = &v
			case fulltextsearch.FIELDS:
				v := []string(value.(fulltextsearch.Fields))
				query.Fields = v
			case fulltextsearch.TIEBREAKER:
				v := types.Float64(value.(fulltextsearch.TieBreaker))
				query.TieBreaker = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query query_string
câu query có dạng query_string[boost=3,fields=a;b;d]
*/
func ParseQueryStringQuery(m fulltextsearch.Querier) *types.QueryStringQuery {
	query := types.NewQueryStringQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				query.Analyzer = &v
			case fulltextsearch.AUTOGENERATESYNONYMSPHRASEQUERY:
				v := bool(value.(fulltextsearch.AutoGenerateSynonymsPhraseQuery))
				query.AutoGenerateSynonymsPhraseQuery = &v
			case fulltextsearch.ALLOWLEADINGWILDCARD:
				v := bool(value.(fulltextsearch.AllowLeadingWildcard))
				query.AllowLeadingWildcard = &v
			case fulltextsearch.ANALYZEWILDCARD:
				v := bool(value.(fulltextsearch.AnalyzeWildcard))
				query.AnalyzeWildcard = &v
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.DEFAULTFIELD:
				v := string(value.(fulltextsearch.DefaultField))
				query.DefaultField = &v
			case fulltextsearch.DEFAULTOPERATOR:
				op := operator.Operator{}
				v := string(value.(fulltextsearch.Operator))
				op.Name = v
				query.DefaultOperator = &op
			case fulltextsearch.ESCAPE:
				v := bool(value.(fulltextsearch.Escape))
				query.Escape = &v
			case fulltextsearch.ENABLEPOSITIONINCREMENTS:
				v := bool(value.(fulltextsearch.EnablePositionIncrements))
				query.EnablePositionIncrements = &v
			case fulltextsearch.FUZZINESS:
				v := types.Fuzziness(value.(fulltextsearch.Fuzziness))
				query.Fuzziness = &v
			case fulltextsearch.FUZZYREWRITE:
				v := string(value.(fulltextsearch.FuzzyRewrite))
				query.FuzzyRewrite = &v
			case fulltextsearch.FUZZYTRANSPOSITIONS:
				v := bool(value.(fulltextsearch.FuzzyTranspositions))
				query.FuzzyTranspositions = &v
			case fulltextsearch.LENIENT:
				v := bool(value.(fulltextsearch.Lenient))
				query.Lenient = &v
			case fulltextsearch.FUZZYMAXEXPANSIONS:
				v := int(value.(fulltextsearch.FuzzyMaxExpansions))
				query.FuzzyMaxExpansions = &v
			case fulltextsearch.MINIMUMSHOULDMATCH:
				v := types.MinimumShouldMatch(value.(fulltextsearch.MinimumShouldMatch))
				query.MinimumShouldMatch = &v
			case fulltextsearch.FUZZYPREFIXLENGTH:
				v := int(value.(fulltextsearch.FuzzyPrefixLength))
				query.FuzzyPrefixLength = &v
			case fulltextsearch.MAXDETERMINIZEDSTATES:
				v := int(value.(fulltextsearch.MaxDeterminizedStates))
				query.MaxDeterminizedStates = &v
			case fulltextsearch.QUERY:
				v := string(value.(fulltextsearch.Query))
				query.Query = v
			case fulltextsearch.QUOTEANALYZER:
				v := string(value.(fulltextsearch.QuoteAnalyzer))
				query.QuoteAnalyzer = &v
			case fulltextsearch.QUOTEFIELDSUFFIX:
				v := string(value.(fulltextsearch.QuoteFieldSuffix))
				query.QuoteFieldSuffix = &v
			case fulltextsearch.TIMEZONE:
				v := string(value.(fulltextsearch.TimeZone))
				query.TimeZone = &v
			case fulltextsearch.REWRITE:
				v := string(value.(fulltextsearch.Rewrite))
				query.Rewrite = &v
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			case fulltextsearch.PHRASESLOP:
				v := types.Float64(value.(fulltextsearch.PhraseSlop))
				query.PhraseSlop = &v
			case fulltextsearch.FIELDS:
				v := []string(value.(fulltextsearch.Fields))
				query.Fields = v
			case fulltextsearch.TIEBREAKER:
				v := types.Float64(value.(fulltextsearch.TieBreaker))
				query.TieBreaker = &v
			case fulltextsearch.TYPE:
				zero := textquerytype.TextQueryType{}
				v := string(value.(fulltextsearch.Type))
				zero.Name = v
				query.Type = &zero
			}
		}
	}
	return query
}

/*
Phân tích câu query simple_query_string
câu query có dạng simple_query_string[boost=3,...]
*/
func ParseSimpleQueryStringQuery(m fulltextsearch.Querier) *types.SimpleQueryStringQuery {
	query := types.NewSimpleQueryStringQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.ANALYZER:
				v := string(value.(fulltextsearch.Analyzer))
				query.Analyzer = &v
			case fulltextsearch.AUTOGENERATESYNONYMSPHRASEQUERY:
				v := bool(value.(fulltextsearch.AutoGenerateSynonymsPhraseQuery))
				query.AutoGenerateSynonymsPhraseQuery = &v

			case fulltextsearch.ANALYZEWILDCARD:
				v := bool(value.(fulltextsearch.AnalyzeWildcard))
				query.AnalyzeWildcard = &v
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v

			case fulltextsearch.DEFAULTOPERATOR:
				op := operator.Operator{}
				v := string(value.(fulltextsearch.Operator))
				op.Name = v
				query.DefaultOperator = &op

			case fulltextsearch.FUZZYTRANSPOSITIONS:
				v := bool(value.(fulltextsearch.FuzzyTranspositions))
				query.FuzzyTranspositions = &v
			case fulltextsearch.LENIENT:
				v := bool(value.(fulltextsearch.Lenient))
				query.Lenient = &v
			case fulltextsearch.FUZZYMAXEXPANSIONS:
				v := int(value.(fulltextsearch.FuzzyMaxExpansions))
				query.FuzzyMaxExpansions = &v
			case fulltextsearch.MINIMUMSHOULDMATCH:
				v := types.MinimumShouldMatch(value.(fulltextsearch.MinimumShouldMatch))
				query.MinimumShouldMatch = &v
			case fulltextsearch.FUZZYPREFIXLENGTH:
				v := int(value.(fulltextsearch.FuzzyPrefixLength))
				query.FuzzyPrefixLength = &v

			case fulltextsearch.QUERY:
				v := string(value.(fulltextsearch.Query))
				query.Query = v

			case fulltextsearch.QUOTEFIELDSUFFIX:
				v := string(value.(fulltextsearch.QuoteFieldSuffix))
				query.QuoteFieldSuffix = &v

			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v

			case fulltextsearch.FIELDS:
				v := []string(value.(fulltextsearch.Fields))
				query.Fields = v

			}
		}
	}
	return query
}
