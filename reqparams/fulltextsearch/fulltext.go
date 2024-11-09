package fulltextsearch

const (
	INTERVALS                       QueryKey = "intervals"
	ALLOF                           QueryKey = "all_of"
	ANYOF                           QueryKey = "any_of"
	MATCH                           QueryKey = "match"
	MATCHALL                        QueryKey = "match_all"
	AUTOGENERATESYNONYMSPHRASEQUERY QueryKey = "auto_generate_synonyms_phrase_query"
	CUTOFFFREQUENCY                 QueryKey = "cutoff_frequency"
	FUZZINESS                       QueryKey = "fuzziness"
	FUZZYREWRITE                    QueryKey = "fuzzy_rewrite"
	FUZZYTRANSPOSITIONS             QueryKey = "fuzzy_transpositions"
	LENIENT                         QueryKey = "lenient"
	MAXEXPANSIONS                   QueryKey = "max_expansions"
	OPERATOR                        QueryKey = "operator"
	PREFIXLENGTH                    QueryKey = "prefix_length"
	ZEROTERMSQUERY                  QueryKey = "zero_terms_query"
	MATCHBOOLPREFIX                 QueryKey = "match_bool_prefix"
	MATCHPHRASE                     QueryKey = "match_phrase"
	MATCHPHRASEPREFIX               QueryKey = "match_phrase_prefix"
	SLOP                            QueryKey = "slop"
	COMBINEDFIELDS                  QueryKey = "combined_fields"
	FIELDS                          QueryKey = "fields"
	MULTIMATCH                      QueryKey = "multi_match"
	REWRITE                         QueryKey = "rewrite"
	TYPE                            QueryKey = "type"

	INTERVALSFILTER QueryKey = "intervals_filter"

	AFTER          QueryKey = "after"
	BEFORE         QueryKey = "before"
	CONTAINEDBY    QueryKey = "contained_by"
	CONTAINING     QueryKey = "containing"
	NOTCONTAINEDBY QueryKey = "not_contained_by"
	NOTCONTAINING  QueryKey = "not_containing"
	NOTOVERLAPPING QueryKey = "not_overlapping"
	OVERLAPPING    QueryKey = "overlapping"
	USEFIELD       QueryKey = "use_field"
	PATTERN        QueryKey = "pattern"

	SLICEINTERVALS QueryKey = "slice_intervals" /* thay thế cho intervals trong AllOf và AnyOff */

	MAXGAPS QueryKey = "max_gaps"
	ORDERED QueryKey = "ordered"

	QUERYSTRING              QueryKey = "query_string"
	SIMPLEQUERYSTRING        QueryKey = "simple_query_string"
	ALLOWLEADINGWILDCARD     QueryKey = "allow_leading_wildcard"
	ANALYZEWILDCARD          QueryKey = "analyze_wildcard"
	DEFAULTFIELD             QueryKey = "default_field"
	DEFAULTOPERATOR          QueryKey = "default_operator"
	ENABLEPOSITIONINCREMENTS QueryKey = "enable_position_increments"
	ESCAPE                   QueryKey = "escape"
	FUZZYMAXEXPANSIONS       QueryKey = "fuzzy_max_expansions"
	FUZZYPREFIXLENGTH        QueryKey = "fuzzy_prefix_length"
	PHRASESLOP               QueryKey = "phrase_slop"
	QUOTEANALYZER            QueryKey = "quote_analyzer"
	QUOTEFIELDSUFFIX         QueryKey = "quote_field_suffix"
	TIMEZONE                 QueryKey = "time_zone"
	MAXDETERMINIZEDSTATES    QueryKey = "max_determinized_states"
)

type Intervals struct {
	Querier
}
type AllOf Intervals
type AnyOf Intervals

type After Intervals
type Before Intervals
type ContainedBy Intervals
type Containing Intervals
type NotContainedBy Intervals
type NotContaining Intervals
type NotOverlapping Intervals
type Overlapping Intervals

type Match struct {
	Querier
}
type MatchBoolPrefix Match
type MatchPhrase Match
type MatchPhrasePrefix Match
type CombinedFields Match
type MultiMatch Match
type QueryString Match
type SimpleQueryString Match
type MatchAll Match

type AutoGenerateSynonymsPhraseQuery bool
type FuzzyTranspositions bool
type Lenient bool
type CutoffFrequency float64
type Fuzziness interface{}
type FuzzyRewrite string
type MaxExpansions int
type Operator string
type PrefixLength int
type ZeroTermsQuery string
type Slop int
type Fields []string
type Type string
type Rewrite string
type Pattern string

type AllowLeadingWildcard bool
type AnalyzeWildcard bool
type EnablePositionIncrements bool
type Escape bool
type FuzzyMaxExpansions int
type FuzzyPrefixLength int
type PhraseSlop float64
type MaxDeterminizedStates int

type DefaultField string
type UseField string
type DefaultOperator string
type QuoteAnalyzer string
type QuoteFieldSuffix string
type TimeZone string
type MaxGaps int
type Ordered bool

type IntervalsFilter struct {
	Querier
}

type SliceIntervals struct {
	Params []Querier
}

func NewSliceIntervals() *SliceIntervals {
	return &SliceIntervals{
		Params: make([]Querier, 0),
	}
}

func (f *SliceIntervals) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *SliceIntervals) GetParams() interface{} {
	return f.Params
}
