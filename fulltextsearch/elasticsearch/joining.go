package elasticsearch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/kuromojitokenizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/language"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snowballlanguage"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

/*
Phân tích câu query nested
câu query có dạng nested[boost=3,...]
  - query_search thay thế cho query
*/
func ParseNestedQuery(m fulltextsearch.Querier) *types.NestedQuery {
	query := types.NewNestedQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.IGNOREUNMAPPED:
				v := bool(value.(fulltextsearch.IgnoreUnmapped))
				query.IgnoreUnmapped = &v
			case fulltextsearch.QUERYSEARCH:
				field := value.(fulltextsearch.Querier)
				i := ParseQueryToSearch(field)
				query.Query = i
			case fulltextsearch.INNERHITS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseInnerHits(field)
				query.InnerHits = i
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			case fulltextsearch.SCOREMODE:
				scm := childscoremode.ChildScoreMode{}
				v := string(value.(fulltextsearch.ScoreMode))
				scm.Name = v
				query.ScoreMode = &scm
			}
		}
	}
	return query
}

/*
Phân tích câu query inner_hits
query có dạng inner_hits[...]
*/
func ParseInnerHits(m fulltextsearch.Querier) *types.InnerHits {
	query := types.NewInnerHits()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.COLLAPSE:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseCollapse(field)
				query.Collapse = i
			case fulltextsearch.DOCVALUEFIELDS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.DocvalueFields = ParseDocValueFields(field)
			case fulltextsearch.NAME:
				v := string(value.(fulltextsearch.Name))
				query.Name = &v
			case fulltextsearch.EXPLAIN:
				v := bool(value.(fulltextsearch.Explain))
				query.Explain = &v
			case fulltextsearch.FIELDS:
				v := []string(value.(fulltextsearch.Fields))
				query.Fields = v
			case fulltextsearch.FROM:
				v := int(value.(fulltextsearch.From))
				query.From = &v
			case fulltextsearch.HIGHLIGHT:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Highlight = ParseHighLight(field)
			case fulltextsearch.IGNOREUNMAPPED:
				v := bool(value.(fulltextsearch.IgnoreUnmapped))
				query.IgnoreUnmapped = &v
			case fulltextsearch.SCRIPTFIELDS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.ScriptFields = ParseScriptFields(field)
			case fulltextsearch.SEQNOPRIMARYTERM:
				v := bool(value.(fulltextsearch.SeqNoPrimaryTerm))
				query.SeqNoPrimaryTerm = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query script_fields
query có dạng script_fields[fields[...]]
*/
func ParseScriptFields(m fulltextsearch.Querier) map[string]types.ScriptField {
	scriptFields := make(map[string]types.ScriptField)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewScriptField()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.IGNOREFAILURE:
					v := bool(value.(fulltextsearch.IgnoreFailure))
					query.IgnoreFailure = &v
				case fulltextsearch.SCRIPT:
					field := value.(fulltextsearch.Querier)
					query.Script = *ParseScript(field)
				}
			}
			scriptFields[string(key)] = *query
		}
	default:
		return nil
	}
	return scriptFields
}

/*
phân tích câu query collapse
query có dạng collapse[collapse[...],slice_inner_hits[inner_hits[...]]]
  - slice_inner_hits thay thế cho inner_hits
*/
func ParseCollapse(m fulltextsearch.Querier) *types.FieldCollapse {
	query := types.NewFieldCollapse()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.COLLAPSE:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseCollapse(field)
				query.Collapse = i
			case fulltextsearch.SLICEINNERHITS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.InnerHits{}
				for _, q := range options {
					i := ParseInnerHits(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.InnerHits = sliceQuery
			case fulltextsearch.FIELD:
				v := string(value.(fulltextsearch.Field))
				query.Field = v
			case fulltextsearch.MAXCONCURRENTGROUPSEARCHES:
				v := int(value.(fulltextsearch.MaxConcurrentGroupSearches))
				query.MaxConcurrentGroupSearches = &v
			}
		}
	}
	return query
}

/*
phân tích câu query doc_value_fields
query có dạng doc_value_fields[field_and_format[field=abc,format=abc],field_and_format[...]]
*/

func ParseDocValueFields(m fulltextsearch.Querier) []types.FieldAndFormat {
	sliceQuery := []types.FieldAndFormat{}
	options, ok := m.GetParams().([]fulltextsearch.Querier)
	if !ok {
		return nil
	}
	for _, q := range options {
		i := ParseFieldAndFormat(q)
		sliceQuery = append(sliceQuery, *i)
	}
	return sliceQuery
}

/*
phân tích câu query field_and_format
query có dạng field_and_format[field=abc,format=abc]
*/
func ParseFieldAndFormat(m fulltextsearch.Querier) *types.FieldAndFormat {
	query := types.NewFieldAndFormat()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.FIELD:
				v := string(value.(fulltextsearch.Field))
				query.Field = v
			case fulltextsearch.FORMAT:
				v := string(value.(fulltextsearch.Format))
				query.Field = v
			case fulltextsearch.INCLUDEUNMAPPED:
				v := bool(value.(fulltextsearch.IncludeUnmapped))
				query.IncludeUnmapped = &v
			}
		}
	}
	return query
}

/*
phân tích câu query highlight
query có dạng highlight[field=abc,format=abc]
  - highlight_fields thay thế cho fields
  - highlight_order thay thế cho order
*/
func ParseHighLight(m fulltextsearch.Querier) *types.Highlight {
	query := types.NewHighlight()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOUNDARYCHARS:
				v := string(value.(fulltextsearch.BoundaryChars))
				query.BoundaryChars = &v
			case fulltextsearch.BOUNDARYMAXSCAN:
				v := int(value.(fulltextsearch.BoundaryMaxScan))
				query.BoundaryMaxScan = &v
			case fulltextsearch.BOUNDARYSCANNER:
				b := boundaryscanner.BoundaryScanner{}
				v := string(value.(fulltextsearch.BoundaryScanner))
				b.Name = v
				query.BoundaryScanner = &b
			case fulltextsearch.BOUNDARYSCANNERLOCALE:
				v := string(value.(fulltextsearch.BoundaryScannerLocale))
				query.BoundaryScannerLocale = &v
			case fulltextsearch.ENCODER:
				b := highlighterencoder.HighlighterEncoder{}
				v := string(value.(fulltextsearch.Encoder))
				b.Name = v
				query.Encoder = &b
			case fulltextsearch.HIGHLIGHTFIELDS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Fields = ParseHighLightFields(field)
			case fulltextsearch.FORCESOURCE:
				v := bool(value.(fulltextsearch.ForceSource))
				query.ForceSource = &v
			case fulltextsearch.FRAGMENTSIZE:
				v := int(value.(fulltextsearch.FragmentSize))
				query.FragmentSize = &v
			case fulltextsearch.FRAGMENTER:
				b := highlighterfragmenter.HighlighterFragmenter{}
				v := string(value.(fulltextsearch.Fragmenter))
				b.Name = v
				query.Fragmenter = &b
			case fulltextsearch.HIGHLIGHTFILTER:
				v := bool(value.(fulltextsearch.HighlightFilter))
				query.HighlightFilter = &v
			case fulltextsearch.HIGHLIGHTQUERY:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.HighlightQuery = ParseQueryToSearch(field)
			case fulltextsearch.MAXANALYZEDOFFSET:
				v := int(value.(fulltextsearch.MaxAnalyzedOffset))
				query.MaxAnalyzedOffset = &v
			case fulltextsearch.MAXFRAGMENTLENGTH:
				v := int(value.(fulltextsearch.MaxFragmentLength))
				query.MaxFragmentLength = &v
			case fulltextsearch.NOMATCHSIZE:
				v := int(value.(fulltextsearch.NoMatchSize))
				query.NoMatchSize = &v
			case fulltextsearch.NUMBEROFFRAGMENTS:
				v := int(value.(fulltextsearch.NumberOfFragments))
				query.NumberOfFragments = &v
			case fulltextsearch.OPTIONS:
				field := value.(fulltextsearch.Querier)
				options := make(map[string]json.RawMessage)
				pars := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				for k, v := range pars {
					strKey := fmt.Sprintf("%v", k)
					strValue := json.RawMessage(fmt.Sprintf("%v", v))
					options[strKey] = strValue
				}
				query.Options = options
			case fulltextsearch.HIGHLIGHTORDER:
				b := highlighterorder.HighlighterOrder{}
				v := string(value.(fulltextsearch.HighlightOrder))
				b.Name = v
				query.Order = &b
			case fulltextsearch.PHRASELIMIT:
				v := int(value.(fulltextsearch.PhraseLimit))
				query.PhraseLimit = &v
			case fulltextsearch.POSTTAGS:
				v := []string(value.(fulltextsearch.PostTags))
				query.PostTags = v
			case fulltextsearch.PRETAGS:
				v := []string(value.(fulltextsearch.PreTags))
				query.PreTags = v
			case fulltextsearch.REQUIREFIELDMATCH:
				v := bool(value.(fulltextsearch.RequireFieldMatch))
				query.RequireFieldMatch = &v
			case fulltextsearch.TAGSSCHEMA:
				b := highlightertagsschema.HighlighterTagsSchema{}
				v := string(value.(fulltextsearch.TagsSchema))
				b.Name = v
				query.TagsSchema = &b
			case fulltextsearch.TYPE:
				b := highlightertype.HighlighterType{}
				v := string(value.(fulltextsearch.Type))
				b.Name = v
				query.Type = &b
			}
		}
	}
	return query
}

/*
Phân tích câu query hightlight_fields
câu query có dạng intervals[fields[all_of[...]]]
  - highlight_analyzer thay the cho analyzer
*/
func ParseHighLightFields(m fulltextsearch.Querier) map[string]types.HighlightField {
	hight := make(map[string]types.HighlightField)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewHighlightField()
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.HIGHLIGHTANALYZER:
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query.Analyzer = ParseHighlightAnalyzer(field)
				case fulltextsearch.BOUNDARYCHARS:
					v := string(value.(fulltextsearch.BoundaryChars))
					query.BoundaryChars = &v
				case fulltextsearch.BOUNDARYMAXSCAN:
					v := int(value.(fulltextsearch.BoundaryMaxScan))
					query.BoundaryMaxScan = &v
				case fulltextsearch.BOUNDARYSCANNER:
					b := boundaryscanner.BoundaryScanner{}
					v := string(value.(fulltextsearch.BoundaryScanner))
					b.Name = v
					query.BoundaryScanner = &b
				case fulltextsearch.FORCESOURCE:
					v := bool(value.(fulltextsearch.ForceSource))
					query.ForceSource = &v
				case fulltextsearch.FRAGMENTSIZE:
					v := int(value.(fulltextsearch.FragmentSize))
					query.FragmentSize = &v
				case fulltextsearch.FRAGMENTER:
					b := highlighterfragmenter.HighlighterFragmenter{}
					v := string(value.(fulltextsearch.Fragmenter))
					b.Name = v
					query.Fragmenter = &b
				case fulltextsearch.HIGHLIGHTFILTER:
					v := bool(value.(fulltextsearch.HighlightFilter))
					query.HighlightFilter = &v
				case fulltextsearch.HIGHLIGHTQUERY:
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query.HighlightQuery = ParseQueryToSearch(field)
				case fulltextsearch.MAXANALYZEDOFFSET:
					v := int(value.(fulltextsearch.MaxAnalyzedOffset))
					query.MaxAnalyzedOffset = &v
				case fulltextsearch.MAXFRAGMENTLENGTH:
					v := int(value.(fulltextsearch.MaxFragmentLength))
					query.MaxFragmentLength = &v
				case fulltextsearch.NOMATCHSIZE:
					v := int(value.(fulltextsearch.NoMatchSize))
					query.NoMatchSize = &v
				case fulltextsearch.NUMBEROFFRAGMENTS:
					v := int(value.(fulltextsearch.NumberOfFragments))
					query.NumberOfFragments = &v
				case fulltextsearch.OPTIONS:
					field := value.(fulltextsearch.Querier)
					options := make(map[string]json.RawMessage)
					pars := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
					for k, v := range pars {
						strKey := fmt.Sprintf("%v", k)
						strValue := json.RawMessage(fmt.Sprintf("%v", v))
						options[strKey] = strValue
					}
					query.Options = options
				case fulltextsearch.HIGHLIGHTORDER:
					b := highlighterorder.HighlighterOrder{}
					v := string(value.(fulltextsearch.HighlightOrder))
					b.Name = v
					query.Order = &b
				case fulltextsearch.PHRASELIMIT:
					v := int(value.(fulltextsearch.PhraseLimit))
					query.PhraseLimit = &v
				case fulltextsearch.POSTTAGS:
					v := []string(value.(fulltextsearch.PostTags))
					query.PostTags = v
				case fulltextsearch.PRETAGS:
					v := []string(value.(fulltextsearch.PreTags))
					query.PreTags = v
				case fulltextsearch.REQUIREFIELDMATCH:
					v := bool(value.(fulltextsearch.RequireFieldMatch))
					query.RequireFieldMatch = &v
				case fulltextsearch.TAGSSCHEMA:
					b := highlightertagsschema.HighlighterTagsSchema{}
					v := string(value.(fulltextsearch.TagsSchema))
					b.Name = v
					query.TagsSchema = &b
				case fulltextsearch.TYPE:
					b := highlightertype.HighlighterType{}
					v := string(value.(fulltextsearch.Type))
					b.Name = v
					query.Type = &b

				}
			}
			hight[string(key)] = *query
		}
	default:
		return nil
	}
	return hight
}

/*
Phân tích highlight_analyzer
câu query dạng highlight_analyzer[custom_analyzer[...],...]

	các dạng highlight_analyzer
	- custom_analyzer
	- finger_print_analyzer
	- keyword_analyzer
	- language_analyzer
	- nori_analyzer
	- pattern_analyzer
	- simple_analyzer
	- standar_analyzer
	- stop_analyzer
	- white_space_analyzer
	- icu_analyzer
	- kuromoji_analyzer
	- snow_ball_analyzer
	- dutch_analyzer
*/
func ParseHighlightAnalyzer(m fulltextsearch.Querier) types.Analyzer {
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.CUSTOMANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseCustomAnalyzer(field)
				return i
			case fulltextsearch.FINGERPRINTANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseFingerPrintAnalyzer(field)
				return i
			case fulltextsearch.KEYWORDANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseKeywordAnalyzer(field)
				return i
			case fulltextsearch.LANGUAGEANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseLanguageAnalyzer(field)
				return i
			case fulltextsearch.NORIANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseNoriAnalyzer(field)
				return i
			case fulltextsearch.PATTERNANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParsePatternAnalyzer(field)
				return i
			case fulltextsearch.SIMPLEANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseSimpleAnalyzer(field)
				return i
			case fulltextsearch.STANDARANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseStandardAnalyzer(field)
				return i
			case fulltextsearch.STOPANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseStopAnalyzer(field)
				return i
			case fulltextsearch.WHITESPACEANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseWhiteSpaceAnalyzer(field)
				return i
			case fulltextsearch.ICUANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseIcuAnalyzer(field)
				return i
			case fulltextsearch.KUROMOJIANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseKuromojiAnalyzer(field)
				return i
			case fulltextsearch.SNOWBALLANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseSnowballAnalyzer(field)
				return i
			case fulltextsearch.DUTCHANALYZER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseDutchAnalyzer(field)
				return i
			}
		}
	}
	return nil
}

/*
phân tích câu query custom_analyzer
query có dạng custom_analyzer[...]
*/
func ParseCustomAnalyzer(m fulltextsearch.Querier) *types.CustomAnalyzer {
	query := types.NewCustomAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.CHARFILTER:
				v := []string(value.(fulltextsearch.CharFilter))
				query.CharFilter = v
			case fulltextsearch.ANALYZERFILTER:
				v := []string(value.(fulltextsearch.AnalyzerFilter))
				query.Filter = v
			case fulltextsearch.POSITIONINCREMENTGAP:
				v := int(value.(fulltextsearch.PositionIncrementGap))
				query.PositionIncrementGap = &v
			case fulltextsearch.POSITIONOFFSETGAP:
				v := int(value.(fulltextsearch.PositionOffsetGap))
				query.PositionOffsetGap = &v
			case fulltextsearch.TOKENIZER:
				v := string(value.(fulltextsearch.Tokenizer))
				query.Tokenizer = v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query stop_analyzer
query có dạng custom_analyzer[...]
  - analizer_version thay the version
*/
func ParseStopAnalyzer(m fulltextsearch.Querier) *types.StopAnalyzer {
	query := types.NewStopAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.STOPWORDSPATH:
				v := string(value.(fulltextsearch.StopwordsPath))
				query.StopwordsPath = &v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query finger_print_analyzer
query có dạng custom_analyzer[...]
  - analizer_version thay the version
*/
func ParseFingerPrintAnalyzer(m fulltextsearch.Querier) *types.FingerprintAnalyzer {
	query := types.NewFingerprintAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.MAXOUTPUTSIZE:
				v := int(value.(fulltextsearch.MaxOutputSize))
				query.MaxOutputSize = v
			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.PRESERVEORIGINAL:
				v := bool(value.(fulltextsearch.PreserveOriginal))
				query.PreserveOriginal = v
			case fulltextsearch.SEPARATOR:
				v := string(value.(fulltextsearch.Separator))
				query.Separator = v
			case fulltextsearch.STOPWORDSPATH:
				v := string(value.(fulltextsearch.StopwordsPath))
				query.StopwordsPath = &v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query language_analyzer
query có dạng language_analyzer[...]
  - analizer_version thay the version
*/
func ParseLanguageAnalyzer(m fulltextsearch.Querier) *types.LanguageAnalyzer {
	query := types.NewLanguageAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.LANGUAGE:
				lang := language.Language{}
				v := string(value.(fulltextsearch.Language))
				lang.Name = v
				query.Language = lang
			case fulltextsearch.STEMEXCLUSION:
				v := []string(value.(fulltextsearch.StemExclusion))
				query.StemExclusion = v
			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.STOPWORDSPATH:
				v := string(value.(fulltextsearch.StopwordsPath))
				query.StopwordsPath = &v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query snowball_analyzer
query có dạng snowball_analyzer[...]
  - analizer_version thay the version
*/
func ParseSnowballAnalyzer(m fulltextsearch.Querier) *types.SnowballAnalyzer {
	query := types.NewSnowballAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.LANGUAGE:
				lang := snowballlanguage.SnowballLanguage{}
				v := string(value.(fulltextsearch.Language))
				lang.Name = v
				query.Language = lang
			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query dutch_analyzer
query có dạng dutch_analyzer[...]
*/
func ParseDutchAnalyzer(m fulltextsearch.Querier) *types.DutchAnalyzer {
	query := types.NewDutchAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {

			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query pattern_analyzer
query có dạng pattern_analyzer[...]
  - analizer_version thay the version
*/
func ParsePatternAnalyzer(m fulltextsearch.Querier) *types.PatternAnalyzer {
	query := types.NewPatternAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.FLAGS:
				v := string(value.(fulltextsearch.Flags))
				query.Flags = &v
			case fulltextsearch.LOWERCASE:
				v := bool(value.(fulltextsearch.Lowercase))
				query.Lowercase = &v
			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.PATTERN:
				v := string(value.(fulltextsearch.Pattern))
				query.Pattern = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query standar_analyzer
query có dạng standar_analyzer[...]
  - analizer_version thay the version
*/
func ParseStandardAnalyzer(m fulltextsearch.Querier) *types.StandardAnalyzer {
	query := types.NewStandardAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.MAXTOKENLENGTH:
				v := int(value.(fulltextsearch.MaxTokenLength))
				query.MaxTokenLength = &v
			case fulltextsearch.STOPWORDS:
				v := []string(value.(fulltextsearch.StopWords))
				query.Stopwords = v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v

			}
		}
	}
	return query
}

/*
phân tích câu query nori_analyzer
query có dạng nori_analyzer[...]
  - analizer_version thay the version
*/
func ParseNoriAnalyzer(m fulltextsearch.Querier) *types.NoriAnalyzer {
	query := types.NewNoriAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.DECOMPOUNDMODE:
				decom := noridecompoundmode.NoriDecompoundMode{}
				v := string(value.(fulltextsearch.DecompoundMode))
				decom.Name = v
				query.DecompoundMode = &decom
			case fulltextsearch.STOPTAGS:
				v := []string(value.(fulltextsearch.Stoptags))
				query.Stoptags = v
			case fulltextsearch.USERDICTIONARY:
				v := string(value.(fulltextsearch.UserDictionary))
				query.UserDictionary = &v
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query white_space_analyzer
query có dạng white_space_analyzer[...]
  - analizer_version thay the version
*/
func ParseWhiteSpaceAnalyzer(m fulltextsearch.Querier) *types.WhitespaceAnalyzer {
	query := types.NewWhitespaceAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query icu_analyzer
query có dạng icu_analyzer[...]
*/
func ParseIcuAnalyzer(m fulltextsearch.Querier) *types.IcuAnalyzer {
	query := types.NewIcuAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.METHOD:
				method := icunormalizationtype.IcuNormalizationType{}
				v := string(value.(fulltextsearch.Method))
				method.Name = v
				query.Method = method
			case fulltextsearch.MODE:
				mod := icunormalizationmode.IcuNormalizationMode{}
				v := string(value.(fulltextsearch.Mode))
				mod.Name = v
				query.Mode = mod
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query icu_analyzer
query có dạng icu_analyzer[...]
*/
func ParseKuromojiAnalyzer(m fulltextsearch.Querier) *types.KuromojiAnalyzer {
	query := types.NewKuromojiAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.MODE:
				mod := kuromojitokenizationmode.KuromojiTokenizationMode{}
				v := string(value.(fulltextsearch.Mode))
				mod.Name = v
				query.Mode = mod
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.USERDICTIONARY:
				v := string(value.(fulltextsearch.UserDictionary))
				query.UserDictionary = &v
			}
		}
	}
	return query
}

/*
phân tích câu query keyword_analyzer
query có dạng keyword_analyzer[...]
  - analizer_version thay the version
*/
func ParseKeywordAnalyzer(m fulltextsearch.Querier) *types.KeywordAnalyzer {
	query := types.NewKeywordAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query simple_analyzer
query có dạng keyword_analyzer[...]
  - analizer_version thay the version
*/
func ParseSimpleAnalyzer(m fulltextsearch.Querier) *types.SimpleAnalyzer {
	query := types.NewSimpleAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.TYPE:
				v := string(value.(fulltextsearch.Type))
				query.Type = v
			case fulltextsearch.ANALYZERVERSION:
				v := string(value.(fulltextsearch.AnalyzerVersion))
				query.Version = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query has_child
câu query có dạng nested[boost=3,...]
  - query_search thay thế cho query
*/
func ParseHasChildQuery(m fulltextsearch.Querier) *types.NestedQuery {
	query := types.NewNestedQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.IGNOREUNMAPPED:
				v := bool(value.(fulltextsearch.IgnoreUnmapped))
				query.IgnoreUnmapped = &v
			case fulltextsearch.QUERYSEARCH:
				field := value.(fulltextsearch.Querier)
				i := ParseQueryToSearch(field)
				query.Query = i
			// case fulltextsearch.INNERHITS:
			// 	field, ok := value.(fulltextsearch.Querier)
			// 	if !ok {
			// 		break
			// 	}
			// 	i := ParseInnerHits(field)
			// 	query.InnerHits = i
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			case fulltextsearch.SCOREMODE:
				scm := childscoremode.ChildScoreMode{}
				v := string(value.(fulltextsearch.ScoreMode))
				scm.Name = v
				query.ScoreMode = &scm
			}
		}
	}
	return query
}
