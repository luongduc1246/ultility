package fulltextsearch

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/luongduc1246/ultility/structure"
)

const (
	/* nested */
	NESTED         QueryKey = "nested"
	IGNOREUNMAPPED QueryKey = "ignore_unmapped"
	INNERHITS      QueryKey = "inner_hits"
	/* innerhits */
	COLLAPSE                   QueryKey = "collapse"
	MAXCONCURRENTGROUPSEARCHES QueryKey = "max_concurrent_group_searches"
	SLICEINNERHITS             QueryKey = "slice_inner_hits"
	DOCVALUEFIELDS             QueryKey = "doc_value_fields"
	FIELDANDFORMAT             QueryKey = "field_and_format"
	INCLUDEUNMAPPED            QueryKey = "include_unmapped"
	EXPLAIN                    QueryKey = "explain"
	FROM                       QueryKey = "from"
	/* highlight */
	HIGHLIGHT         QueryKey = "highlight"
	HIGHLIGHTANALYZER QueryKey = "highlight_analyzer"
	/* highlight analyzer */
	CUSTOMANALYZER QueryKey = "custom_analyzer"
	CHARFILTER     QueryKey = "char_filter"
	// thay thế cho filter trong custom_analyzer
	ANALYZERFILTER       QueryKey = "analyzer_filter"
	POSITIONINCREMENTGAP QueryKey = "position_increment_gap"
	POSITIONOFFSETGAP    QueryKey = "position_offset_gap"
	TOKENIZER            QueryKey = "tokenizer"

	FINGERPRINTANALYZER QueryKey = "finger_print_analyzer"
	MAXOUTPUTSIZE       QueryKey = "max_output_size"
	PRESERVEORIGINAL    QueryKey = "preserve_original"
	SEPARATOR           QueryKey = "separator"
	STOPWORDSPATH       QueryKey = "stopwords_path"
	ANALYZERVERSION     QueryKey = "analyzer_version"

	KEYWORDANALYZER QueryKey = "keyword_analyzer"

	LANGUAGEANALYZER QueryKey = "language_analyzer"
	LANGUAGE         QueryKey = "language"
	STEMEXCLUSION    QueryKey = "stem_exclusion"

	NORIANALYZER   QueryKey = "nori_analyzer"
	DECOMPOUNDMODE QueryKey = "decompound_mode"
	STOPTAGS       QueryKey = "stoptags"
	USERDICTIONARY QueryKey = "user_dictionary"

	PATTERNANALYZER QueryKey = "pattern_analyzer"
	LOWERCASE       QueryKey = "lowercase"

	SIMPLEANALYZER  QueryKey = "simple_analyzer"
	STANDARANALYZER QueryKey = "standar_analyzer"
	MAXTOKENLENGTH  QueryKey = "max_token_length"

	STOPANALYZER       QueryKey = "stop_analyzer"
	WHITESPACEANALYZER QueryKey = "white_space_analyzer"

	ICUANALYZER QueryKey = "icu_analyzer"
	METHOD      QueryKey = "method"
	MODE        QueryKey = "mode"

	KUROMOJIANALYZER QueryKey = "kuromoji_analyzer"
	SNOWBALLANALYZER QueryKey = "snowball_analyzer"
	DUTCHANALYZER    QueryKey = "dutch_analyzer"
	/*  */
	BOUNDARYCHARS         QueryKey = "boundary_chars"
	BOUNDARYMAXSCAN       QueryKey = "boundary_max_scan"
	BOUNDARYSCANNER       QueryKey = "boundary_scanner"
	BOUNDARYSCANNERLOCALE QueryKey = "boundary_scanner_locale"
	ENCODER               QueryKey = "encoder"
	HIGHLIGHTFIELDS       QueryKey = "highlight_fields"
	FORCESOURCE           QueryKey = "force_source"
	FRAGMENTSIZE          QueryKey = "fragment_size"
	FRAGMENTER            QueryKey = "fragmenter"
	HIGHLIGHTFILTER       QueryKey = "highlight_filter"
	HIGHLIGHTQUERY        QueryKey = "highlight_query"
	MAXANALYZEDOFFSET     QueryKey = "max_analyzed_offset"
	MAXFRAGMENTLENGTH     QueryKey = "max_fragment_length"
	NOMATCHSIZE           QueryKey = "no_match_size"
	NUMBEROFFRAGMENTS     QueryKey = "number_of_fragments"
	HIGHLIGHTORDER        QueryKey = "highlight_order"
	PHRASELIMIT           QueryKey = "phrase_limit"
	POSTTAGS              QueryKey = "post_tags"
	PRETAGS               QueryKey = "pre_tags"
	REQUIREFIELDMATCH     QueryKey = "require_field_match"
	TAGSSCHEMA            QueryKey = "tags_schema"
	/*  */
	SCRIPTFIELDS     QueryKey = "script_fields"
	IGNOREFAILURE    QueryKey = "ignore_failure"
	SEQNOPRIMARYTERM QueryKey = "seq_no_primary_term"
	SIZE             QueryKey = "size"

	SORT             QueryKey = "sort"
	SORTCOMBINATIONS QueryKey = "sort_combinations"

	DOC_  QueryKey = "_doc"
	ORDER QueryKey = "order"

	GEODISTANCE_    QueryKey = "_geo_distance"
	GEODISTANCESORT QueryKey = "geo_distance_sort"
	SLICEGEOLOCATION

	SCORE_      QueryKey = "_score"
	SCRIPT_     QueryKey = "_score"
	SORTOPTIONS QueryKey = "sort_options"

	SOURCE_          QueryKey = "_source"
	STOREDFIELDS     QueryKey = "stored_fields"
	TRACKSCORES      QueryKey = "track_scores"
	INNERHITSVERSION QueryKey = "innerhit_version"
	/*  */
	/*  */

	HASCHILD    QueryKey = "has_child"
	MAXCHILDREN QueryKey = "max_children"
	MINCHILDREN QueryKey = "min_children"

	HASPARENT  QueryKey = "has_parent"
	SCORE      QueryKey = "score"
	PARENTTYPE QueryKey = "parent_type"
	PARENTID   QueryKey = "parent_id"
	ID         QueryKey = "id"
)

/* innerhit */
type Collapse struct {
	Querier
}
type MaxConcurrentGroupSearches int
type FieldAndFormat struct {
	Querier
}
type IncludeUnmapped bool
type Explain bool
type From int

/* highlight */
type Highlight struct {
	Querier
}
type HighlightAnalyzer struct {
	Querier
}

type CustomAnalyzer struct {
	Querier
}
type CharFilter []string
type AnalyzerFilter []string
type PositionIncrementGap int
type PositionOffsetGap int
type Tokenizer string

type FingerprintAnalyzer struct {
	Querier
}
type MaxOutputSize int
type PreserveOriginal bool
type Separator string
type StopwordsPath string
type AnalyzerVersion string

type KeywordAnalyzer struct {
	Querier
}

type LanguageAnalyzer struct {
	Querier
}
type StemExclusion []string
type Language string

type NoriAnalyzer struct {
	Querier
}
type DecompoundMode string
type Stoptags []string
type UserDictionary string

type PatternAnalyzer struct {
	Querier
}
type Lowercase bool

type SimpleAnalyzer struct {
	Querier
}

type StandardAnalyzer struct {
	Querier
}
type MaxTokenLength int

type StopAnalyzer struct {
	Querier
}
type WhitespaceAnalyzer struct {
	Querier
}

type IcuAnalyzer struct {
	Querier
}
type Method string
type Mode string

type KuromojiAnalyzer struct {
	Querier
}
type SnowballAnalyzer struct {
	Querier
}
type DutchAnalyzer struct {
	Querier
}

type BoundaryChars string
type BoundaryMaxScan int
type BoundaryScanner string
type BoundaryScannerLocale string
type Encoder string

type HighlightFields struct {
	Querier
}
type ForceSource bool
type FragmentSize int
type Fragmenter string
type HighlightFilter bool

type HighlightQuery struct {
	Querier
}

type MaxAnalyzedOffset int
type MaxFragmentLength int
type NoMatchSize int
type NumberOfFragments int
type HighlightOrder string
type PhraseLimit int
type PostTags []string
type PreTags []string
type RequireFieldMatch bool
type TagsSchema string

/*  */
type ScriptFields struct {
	Querier
}
type IgnoreFailure bool
type SeqNoPrimaryTerm bool
type Size int
type SortCombinations struct {
	Querier
}
type Doc_ struct {
	Querier
}
type Order string
type Source_ struct {
	Querier
}
type StoredFields []string
type TrackScores bool
type InnerhitsVersion bool

/*  */

type Nested struct {
	Querier
}
type InnerHits Nested

type IgnoreUnmapped bool

type HasChild struct {
	Querier
}

type MaxChildren int
type MinChildren int

type HasParent struct {
	Querier
}

type Score bool
type ParentType string

type ParentId struct {
	Querier
}

type Id string

type SliceInnerHits struct {
	Params []Querier
}

func NewSliceInnerHits() *SliceInnerHits {
	return &SliceInnerHits{
		Params: make([]Querier, 0),
	}
}

func (f *SliceInnerHits) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *SliceInnerHits) GetParams() interface{} {
	return f.Params
}

/* chứa mảng []SortCombinations */
type Sort struct {
	Params []Querier
}

func NewSort() *Sort {
	return &Sort{
		Params: make([]Querier, 0),
	}
}

func (f *Sort) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Sort) GetParams() interface{} {
	return f.Params
}

/* DocValueFields là Slice FieldAndFormat  */
type DocvalueFields struct {
	Params []Querier
}

func NewDocvalueFields() *DocvalueFields {
	return &DocvalueFields{
		Params: make([]Querier, 0),
	}
}

func (f *DocvalueFields) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *DocvalueFields) GetParams() interface{} {
	return f.Params
}

func stringToJoiningQuerier(stack *structure.Stack[Querier], key string, i int, indexStart *int) error {
	switch QueryKey(key) {
	case HIGHLIGHTFIELDS:
		match := HighlightFields{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(HIGHLIGHTFIELDS, match)
		stack.Push(match)
		*indexStart = i + 1
	case HIGHLIGHTQUERY:
		match := HighlightQuery{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(HIGHLIGHTQUERY, match)
		stack.Push(match)
		*indexStart = i + 1
	case HIGHLIGHTANALYZER:
		match := HighlightAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(HIGHLIGHTANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case CUSTOMANALYZER:
		match := CustomAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(CUSTOMANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case FINGERPRINTANALYZER:
		match := FingerprintAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(FINGERPRINTANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case KEYWORDANALYZER:
		match := FingerprintAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(KEYWORDANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case LANGUAGEANALYZER:
		match := LanguageAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(LANGUAGEANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case NORIANALYZER:
		match := NoriAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(NORIANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case PATTERNANALYZER:
		match := PatternAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(PATTERNANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case SIMPLEANALYZER:
		match := SimpleAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(SIMPLEANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case STANDARANALYZER:
		match := StandardAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(STANDARANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case STOPANALYZER:
		match := StopAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(STOPANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case WHITESPACEANALYZER:
		match := WhitespaceAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(WHITESPACEANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case ICUANALYZER:
		match := IcuAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(ICUANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case KUROMOJIANALYZER:
		match := KuromojiAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(KUROMOJIANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case SNOWBALLANALYZER:
		match := SnowballAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(SNOWBALLANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case DUTCHANALYZER:
		match := DutchAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(DUTCHANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case SCRIPTFIELDS:
		match := DutchAnalyzer{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(DUTCHANALYZER, match)
		stack.Push(match)
		*indexStart = i + 1
	case SORT:
		q := NewSort()
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(SORT, q)
		stack.Push(q)
		*indexStart = i + 1
	case SORTCOMBINATIONS:
		q := SortCombinations{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(SORTCOMBINATIONS, q)
		stack.Push(q)
		*indexStart = i + 1
	case DOC_:
		q := Doc_{NewQuerySearch()}
		peek, err := stack.Peek()
		if err != nil {
			return err
		}
		peek.AddParam(DOC_, q)
		stack.Push(q)
		*indexStart = i + 1
	}
	return nil
}

func stringToJoining(key string, value string) interface{} {
	highlight := stringToHighLightQuery(key, value)
	if highlight != nil {
		return highlight
	}
	custom := stringToCustomAnalyzer(key, value)
	if custom != nil {
		return custom
	}
	finger := stringToFingerprintAnalyzer(key, value)
	if finger != nil {
		return finger
	}
	lang := stringToLanguageAnalyzer(key, value)
	if lang != nil {
		return lang
	}
	nori := stringToNoriAnalyzer(key, value)
	if nori != nil {
		return nori
	}
	pattern := stringToPatternAnalyzer(key, value)
	if pattern != nil {
		return pattern
	}
	standard := stringToStandardAnalyzer(key, value)
	if standard != nil {
		return standard
	}
	icu := stringToIcuAnalyzer(key, value)
	if icu != nil {
		return icu
	}
	scriptField := stringToScriptField(key, value)
	if scriptField != nil {
		return scriptField
	}
	doc_ := stringToDoc_(key, value)
	if doc_ != nil {
		return doc_
	}
	return nil
}

func stringToDoc_(key string, value string) interface{} {
	switch QueryKey(key) {
	case ORDER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Order(value)
	}
	return nil
}
func stringToInnerHits(key string, value string) interface{} {
	switch QueryKey(key) {
	case FROM:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return From(v)
	case SIZE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return Size(v)
	case EXPLAIN:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return Explain(boolValue)
	case IGNOREUNMAPPED:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return IgnoreUnmapped(boolValue)
	case SEQNOPRIMARYTERM:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return SeqNoPrimaryTerm(boolValue)
	}
	return nil
}

func stringToScriptField(key string, value string) interface{} {
	switch QueryKey(key) {
	case IGNOREFAILURE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return IgnoreFailure(boolValue)
	}
	return nil
}

func stringToStandardAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case MAXTOKENLENGTH:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return MaxTokenLength(v)
	}
	return nil
}
func stringToLanguageAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case LANGUAGE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Language(value)
	case STEMEXCLUSION:
		fields := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return nil
		}
		for _, v := range vals {
			val, err := (url.QueryUnescape(v))
			if err == nil {
				fields = append(fields, val)
			}
		}
		return StemExclusion(fields)
	}
	return nil
}
func stringToIcuAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case METHOD:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Method(value)
	case MODE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Mode(value)
	}
	return nil
}
func stringToPatternAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case LOWERCASE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return Lowercase(boolValue)
	}
	return nil
}
func stringToNoriAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case DECOMPOUNDMODE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return DecompoundMode(value)
	case USERDICTIONARY:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return UserDictionary(value)
	case STOPTAGS:
		fields := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return nil
		}
		for _, v := range vals {
			val, err := (url.QueryUnescape(v))
			if err == nil {
				fields = append(fields, val)
			}
		}
		return Stoptags(fields)
	}
	return nil
}

func stringToFingerprintAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case MAXOUTPUTSIZE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return MaxOutputSize(v)
	case PRESERVEORIGINAL:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return PreserveOriginal(boolValue)
	case SEPARATOR:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Separator(value)
	case STOPWORDSPATH:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return StopwordsPath(value)
	case ANALYZERVERSION:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return AnalyzerVersion(value)
	}
	return nil
}

func stringToCustomAnalyzer(key string, value string) interface{} {
	switch QueryKey(key) {
	case CHARFILTER:
		fields := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return nil
		}
		for _, v := range vals {
			val, err := (url.QueryUnescape(v))
			if err == nil {
				fields = append(fields, val)
			}
		}
		return CharFilter(fields)
	case ANALYZERFILTER:
		fields := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return nil
		}
		for _, v := range vals {
			val, err := (url.QueryUnescape(v))
			if err == nil {
				fields = append(fields, val)
			}
		}
		return AnalyzerFilter(fields)
	case POSITIONINCREMENTGAP:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return PositionIncrementGap(v)
	case POSITIONOFFSETGAP:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return PositionOffsetGap(v)
	case TYPE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Type(value)
	case TOKENIZER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Tokenizer(value)
	}
	return nil
}

func stringToHighLightQuery(key string, value string) interface{} {

	switch QueryKey(key) {
	case BOUNDARYCHARS:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return QueryName(value)
	case BOUNDARYMAXSCAN:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return BoundaryMaxScan(v)
	case BOUNDARYSCANNER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return BoundaryScanner(value)
	case BOUNDARYSCANNERLOCALE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return BoundaryScannerLocale(value)
	case ENCODER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Encoder(value)
	case FORCESOURCE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return ForceSource(boolValue)
	case FRAGMENTSIZE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return FragmentSize(v)
	case FRAGMENTER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return Fragmenter(value)
	case HIGHLIGHTFILTER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return HighlightFilter(boolValue)
	case MAXANALYZEDOFFSET:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return MaxAnalyzedOffset(v)
	case MAXFRAGMENTLENGTH:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return MaxFragmentLength(v)
	case NOMATCHSIZE:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return NoMatchSize(v)
	case NUMBEROFFRAGMENTS:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return NumberOfFragments(v)
	case HIGHLIGHTORDER:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return HighlightOrder(value)
	case PHRASELIMIT:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return PhraseLimit(v)
	case POSTTAGS:
		fields := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return nil
		}
		for _, v := range vals {
			val, err := (url.QueryUnescape(v))
			if err == nil {
				fields = append(fields, val)
			}
		}
		return PostTags(fields)
	case PRETAGS:
		fields := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return nil
		}
		for _, v := range vals {
			val, err := (url.QueryUnescape(v))
			if err == nil {
				fields = append(fields, val)
			}
		}
		return PreTags(fields)
	case REQUIREFIELDMATCH:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		return RequireFieldMatch(boolValue)
	case TAGSSCHEMA:
		value, err := url.QueryUnescape(value)
		if err != nil {
			return nil
		}
		return TagsSchema(value)
	}
	return nil
}
