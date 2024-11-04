package elasticsearch

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

/* Intervals */
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
				options := field.GetParams().(map[string]string)
				query.Options = options
			case fulltextsearch.PARAMS:
				field := value.(fulltextsearch.Querier)
				params := field.GetParams().(map[string]json.RawMessage)
				query.Params = params
			case fulltextsearch.SOURCE:
				v := string(value.(fulltextsearch.Source))
				query.Source = &v

			}
		}
	}
	return query
}

func ParseAnyOfQuery(m fulltextsearch.Querier) *types.IntervalsAnyOf {
	anyOf := types.NewIntervalsAnyOf()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.FILTER:
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
				v := string(value.(fulltextsearch.Fuzziness))
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
			case fulltextsearch.FILTER:
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

/* match */
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
					var op *operator.Operator
					switch value {
					case "and":
						op = &operator.And
					case "or":
						op = &operator.And
					}
					query.Operator = op
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
					var zero *zerotermsquery.ZeroTermsQuery
					switch value {
					case "none":
						zero = &zerotermsquery.None
					case "all":
						zero = &zerotermsquery.All
					}
					query.ZeroTermsQuery = zero
				}
			}
			matchQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return matchQuery
}
