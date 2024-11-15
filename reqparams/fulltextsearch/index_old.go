package fulltextsearch

// /*
// 	Các giá trị sau nên được url escape
// */

// import (
// 	"encoding/json"
// 	"net/url"
// 	"strconv"
// 	"strings"

// 	"github.com/luongduc1246/ultility/structure"
// )

// type QueryKey string

// const (
// 	BOOST       QueryKey = "boost"
// 	QUERY       QueryKey = "query"
// 	QUERYSEARCH QueryKey = "query_search"
// 	ANALYZER    QueryKey = "analyzer"
// 	QUERYNAME   QueryKey = "_name"
// )

// type Boost float32

// type Analyzer string
// type QueryName string

// type Query string

// type Paramer interface{}

// type Param struct {
// 	Field string
// 	Value interface{}
// }

// type Querier interface {
// 	AddParam(i QueryKey, v interface{})
// 	GetParams() interface{}
// }

// type QuerySearch struct {
// 	Params map[QueryKey]interface{}
// }

// func NewQuerySearch() *QuerySearch {
// 	return &QuerySearch{
// 		Params: make(map[QueryKey]interface{}),
// 	}
// }

// func (q *QuerySearch) AddParam(i QueryKey, v interface{}) {
// 	q.Params[i] = v
// }
// func (q *QuerySearch) GetParams() interface{} {
// 	return q.Params
// }

// func (q *QuerySearch) Parse(s string) error {
// 	stack := structure.NewStack[Querier]()
// 	stack.Push(q)
// 	defer stack.Clear()
// 	var indexStart, indexValue int
// 	for i, v := range s {
// 		switch v {
// 		case '[':
// 			switch QueryKey(s[indexStart:i]) {
// 			case QUERYSEARCH:
// 				query := NewQuerySearch()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MATCH, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			/* fulltext */
// 			case MATCH:
// 				match := Match{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MATCH, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case MATCHALL:
// 				match := MatchAll{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MATCHALL, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case MATCHBOOLPREFIX:
// 				match := MatchBoolPrefix{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MATCHBOOLPREFIX, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case MATCHPHRASE:
// 				match := MatchPhrase{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MATCHPHRASE, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case MATCHPHRASEPREFIX:
// 				match := MatchPhrasePrefix{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MATCHPHRASEPREFIX, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case COMBINEDFIELDS:
// 				match := CombinedFields{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(COMBINEDFIELDS, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case MULTIMATCH:
// 				match := MultiMatch{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MULTIMATCH, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case QUERYSTRING:
// 				match := QueryString{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(QUERYSTRING, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case SIMPLEQUERYSTRING:
// 				match := SimpleQueryString{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SIMPLEQUERYSTRING, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case AFTER:
// 				match := After{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(AFTER, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case BEFORE:
// 				match := Before{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(BEFORE, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case CONTAINEDBY:
// 				match := ContainedBy{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(CONTAINEDBY, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case CONTAINING:
// 				match := Containing{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(CONTAINING, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case NOTCONTAINEDBY:
// 				match := NotContainedBy{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(NOTCONTAINEDBY, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case NOTCONTAINING:
// 				match := NotContaining{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(NOTCONTAINING, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case NOTOVERLAPPING:
// 				match := NotOverlapping{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(NOTOVERLAPPING, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case OVERLAPPING:
// 				match := Overlapping{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(OVERLAPPING, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case SLICEINTERVALS:
// 				match := NewSliceIntervals()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SLICEINTERVALS, match)
// 				stack.Push(match)
// 				indexStart = i + 1
// 			case INTERVALS:
// 				query := Intervals{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(INTERVALS, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case ALLOF:
// 				query := AllOf{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(ALLOF, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case INTERVALSFILTER:
// 				query := IntervalsFilter{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(INTERVALSFILTER, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case ANYOF:
// 				query := AnyOf{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(ANYOF, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			/* compound */
// 			case BOOL:
// 				bool := Bool{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(BOOL, bool)
// 				stack.Push(bool)
// 				indexStart = i + 1

// 			case MUST:
// 				must := NewMust()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MUST, must)
// 				stack.Push(must)
// 				indexStart = i + 1
// 			case MUSTNOT:
// 				mustNot := NewMustNot()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MUSTNOT, mustNot)
// 				stack.Push(mustNot)
// 				indexStart = i + 1
// 			case FILTER:
// 				filter := NewFilter()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FILTER, filter)
// 				stack.Push(filter)
// 				indexStart = i + 1
// 			case SHOULD:
// 				should := NewShould()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FILTER, should)
// 				stack.Push(should)
// 				indexStart = i + 1

// 			case BOOSTING:
// 				boosting := Boosting{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(BOOSTING, boosting)
// 				stack.Push(boosting)
// 				indexStart = i + 1
// 			case POSITIVE:
// 				positive := Positive{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(POSITIVE, positive)
// 				stack.Push(positive)
// 				indexStart = i + 1
// 			case NEGATIVE:
// 				negative := Negative{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(NEGATIVE, negative)
// 				stack.Push(negative)
// 				indexStart = i + 1

// 			case CONSTANTSCORE:
// 				score := ConstantScore{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(CONSTANTSCORE, score)
// 				stack.Push(score)
// 				indexStart = i + 1
// 			case DISMAX:
// 				query := DisMax{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(DISMAX, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case QUERIES:
// 				query := NewQueries()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(QUERIES, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case FUNCTIONSCORE:
// 				query := FunctionScore{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FUNCTIONSCORE, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case SCRIPTSCORE:
// 				query := ScriptScore{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SCRIPTSCORE, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case RANDOMSCORE:
// 				query := RandomScore{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(RANDOMSCORE, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case FIELDVALUEFACTOR:
// 				query := FieldValueFactor{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FIELDVALUEFACTOR, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case FUNCTIONS:
// 				query := NewFunctions()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FUNCTIONS, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 				// decay function
// 			case EXP:
// 				query := Exp{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(EXP, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case GAUSS:
// 				query := Gauss{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(GAUSS, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case UNTYPEDDECAYFUNCTION:
// 				query := UntypedDecayFunction{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(UNTYPEDDECAYFUNCTION, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case DATEDECAYFUNCTION:
// 				query := DateDecayFunction{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(DATEDECAYFUNCTION, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case NUMERICDECAYFUNCTION:
// 				query := NumericDecayFunction{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(NUMERICDECAYFUNCTION, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case GEODECAYFUNCTION:
// 				query := GeoDecayFunction{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(GEODECAYFUNCTION, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			case DECAYPARAMETERS:
// 				query := DecayParameters{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(DECAYPARAMETERS, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 				/* joining */
// 			case NESTED:
// 				query := Nested{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(NESTED, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case INNERHITS:
// 				query := InnerHits{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(INNERHITS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case COLLAPSE:
// 				query := Collapse{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(COLLAPSE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SLICEINNERHITS:
// 				query := NewSliceInnerHits()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SLICEINNERHITS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case DOCVALUEFIELDS:
// 				query := NewDocvalueFields()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(DOCVALUEFIELDS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case FIELDANDFORMAT:
// 				query := FieldAndFormat{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FIELDANDFORMAT, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case HIGHLIGHT:
// 				query := FieldAndFormat{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FIELDANDFORMAT, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case HASCHILD:
// 				query := HasChild{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(HASCHILD, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case HASPARENT:
// 				query := HasParent{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(HASPARENT, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case PARENTID:
// 				query := ParentId{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(PARENTID, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 				/* term  */
// 			case EXISTS:
// 				query := Exists{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(EXISTS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case FUZZY:
// 				query := Fuzzy{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(FUZZY, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case IDS:
// 				query := Ids{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(IDS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case PREFIX:
// 				query := Prefix{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(PREFIX, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case RANGE:
// 				query := Range{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(RANGE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case REGEXP:
// 				query := Regexp{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(REGEXP, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case TERM:
// 				query := Term{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(TERM, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case TERMS:
// 				query := Terms{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(TERMS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case TERMSSET:
// 				query := TermsSet{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(TERMSSET, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case MINIMUMSHOULDMATCHSCRIPT:
// 				query := MinimumShouldMatchScript{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MINIMUMSHOULDMATCHSCRIPT, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case WILDCARD:
// 				query := Wildcard{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(WILDCARD, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 				/* geo */
// 			case GEOBOUNDINGBOX:
// 				query := GeoBoundingBox{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(GEOBOUNDINGBOX, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case GEODISTANCE:
// 				query := GeoDistance{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(GEODISTANCE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case GEOPOLYGON:
// 				query := GeoPolygon{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(GEOPOLYGON, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case GEOSHAPE:
// 				query := GeoShape{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(GEOSHAPE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SHAPE:
// 				query := Shape{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SHAPE, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 				/* span */

// 			case SPANCONTAINING:
// 				query := SpanContaining{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANCONTAINING, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case BIG:
// 				query := Big{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(BIG, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case LITTLE:
// 				query := Little{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(LITTLE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANFIELDMASKING:
// 				query := SpanFieldMasking{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANFIELDMASKING, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANFIRST:
// 				query := SpanFirst{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANFIRST, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANMULTITERM:
// 				query := SpanMultiTerm{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANMULTITERM, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANNEAR:
// 				query := SpanNear{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANNEAR, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case CLAUSES:
// 				query := NewClauses()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(CLAUSES, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANNOT:
// 				query := SpanNot{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANNOT, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case EXCLUDE:
// 				query := Exclude{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(EXCLUDE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case INCLUDE:
// 				query := Include{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(INCLUDE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANOR:
// 				query := SpanOr{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANOR, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANTERM:
// 				query := SpanTerm{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANTERM, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPANWITHIN:
// 				query := SpanWithin{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPANWITHIN, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			/* vector */
// 			case KNN:
// 				query := Knn{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(KNN, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case QUERYVECTORBUILDER:
// 				query := QueryVectorBuilder{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(QUERYVECTORBUILDER, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SPARSEVECTOR:
// 				query := SparseVector{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SPARSEVECTOR, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case PRUNINGCONFIG:
// 				query := PruningConfig{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(PRUNINGCONFIG, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SEMANTIC:
// 				query := Semantic{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SEMANTIC, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 				/* special */
// 			case DISTANCEFEATURE:
// 				query := DistanceFeature{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(DISTANCEFEATURE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case MORELIKETHIS:
// 				query := MoreLikeThis{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(MORELIKETHIS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case PERCOLATE:
// 				query := Percolate{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(PERCOLATE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case RANKFEATURE:
// 				query := RankFeature{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(RANKFEATURE, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case LINEAR:
// 				query := Linear{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(LINEAR, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case LOG:
// 				query := Log{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(LOG, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SATURATION:
// 				query := Saturation{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SATURATION, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SIGMOID:
// 				query := Sigmoid{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SIGMOID, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case SCRIPT:
// 				query := Script{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(SCRIPT, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case OPTIONS:
// 				query := Options{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(OPTIONS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case PARAMS:
// 				query := Params{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(PARAMS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case WRAPPER:
// 				query := Wrapper{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(WRAPPER, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case PINNED:
// 				query := Pinned{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(PINNED, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case DOCS:
// 				query := Docs{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(DOCS, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case ORGANIC:
// 				query := Organic{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(ORGANIC, query)
// 				stack.Push(query)
// 				indexStart = i + 1
// 			case RULE:
// 				query := Rule{NewQuerySearch()}
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(RULE, query)
// 				stack.Push(query)
// 				indexStart = i + 1

// 			default:
// 				err := stringToJoiningQuerier(stack, s[indexStart:i], i, &indexStart)
// 				if err != nil {
// 					return err
// 				}
// 				result := NewQuerySearch()
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddParam(QueryKey(s[indexStart:i]), result)
// 				stack.Push(result)
// 				indexStart = i + 1
// 			}
// 		case '=':
// 			indexValue = i + 1
// 		case ']':
// 			if s[i-1] != ']' {
// 				if (indexStart < indexValue-1) && (indexValue < i) {
// 					key := s[indexStart : indexValue-1]
// 					value := s[indexValue:i]
// 					peek, err := stack.Peek()
// 					if err != nil {
// 						return err
// 					}
// 					indexStart = i + 1
// 					v := stringToQuery(key, value)
// 					peek.AddParam(QueryKey(key), v)
// 					stack.Pop()
// 				} else {
// 					var txtError string
// 					if i < indexStart {
// 						txtError = s[i:indexStart]
// 					} else {
// 						txtError = s[indexStart:i]
// 					}
// 					return ErrorQuery{
// 						At: txtError,
// 					}
// 				}
// 			} else {
// 				indexStart = i + 2
// 				stack.Pop()
// 			}
// 		case ',':
// 			if s[i-1] != ']' {
// 				if (indexStart < indexValue-1) && (indexValue < i) {
// 					key := s[indexStart : indexValue-1]
// 					value := s[indexValue:i]
// 					peek, err := stack.Peek()
// 					if err != nil {
// 						return err
// 					}
// 					v := stringToQuery(key, value)
// 					peek.AddParam(QueryKey(key), v)
// 				} else {
// 					var txtError string
// 					if i < indexStart {
// 						txtError = s[i:indexStart]
// 					} else {
// 						txtError = s[indexStart:i]
// 					}
// 					return ErrorQuery{
// 						Index: i,
// 						At:    txtError,
// 					}
// 				}
// 			}
// 			indexStart = i + 1
// 		}
// 		/*  */
// 		if (i == len(s)-1) && s[i] != ']' {
// 			if (indexStart < indexValue-1) && (indexValue < i) {
// 				key := s[indexStart : indexValue-1]
// 				value := s[indexValue:]
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				v := stringToQuery(key, value)
// 				peek.AddParam(QueryKey(key), v)
// 			} else {
// 				var txtError string
// 				if i < indexStart {
// 					txtError = s[i:indexStart]
// 				} else {
// 					txtError = s[indexStart:i]
// 				}
// 				return ErrorQuery{
// 					Index: i,
// 					At:    txtError,
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func stringToQuery(key string, value string) interface{} {
// 	switch QueryKey(key) {
// 	/* interface */
// 	case MINIMUMSHOULDMATCH:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v MinimumShouldMatch
// 		v = value
// 		return v
// 	case FUZZINESS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Fuzziness
// 		v = value
// 		return v
// 	case DECAY:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Decay
// 		v = value
// 		return v
// 	case OFFSET:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Offset
// 		v = value
// 		return v
// 	case ORIGIN:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Origin
// 		v = value
// 		return v
// 	case SCALE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Scale
// 		v = value
// 		return v
// 	case GT:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Gt
// 		v = value
// 		return v
// 	case GTE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Gte
// 		v = value
// 		return v
// 	case LT:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Lt
// 		v = value
// 		return v
// 	case LTE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var v Lte
// 		v = value
// 		return v

// 		/* []byte */
// 	case MATCHCRITERIA:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		jval := json.RawMessage(value)
// 		return MatchCriteria(jval)
// 	case DOCUMENT:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		jval := json.RawMessage(value)
// 		return Document(jval)

// 	/* string */
// 	case QUERYNAME:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return QueryName(value)
// 	case REWRITE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Rewrite(value)
// 	case SEED:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Seed(value)
// 	case MODIFIER:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Modifier(value)
// 	case USEFIELD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return UseField(value)
// 	case PATTERN:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Pattern(value)
// 	case TERMSTRING:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return TermString(value)
// 	case PREFIXSTRING:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return PrefixString(value)
// 	case QUERY:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Query(value)
// 	case LANG:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Lang(value)
// 	case SOURCE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Source(value)
// 	case INDEX:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Index(value)
// 	case NAME:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Name(value)
// 	case MULTIVALUEMODE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MultiValueMode(value)
// 	case PREFERENCE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Preference(value)
// 	case ROUTING:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Routing(value)
// 	case VERSIONTYPE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return VersionType(value)
// 	case INFERENCEID:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return InferenceId(value)
// 	case DISTANCE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Distance(value)
// 	case DISTANCETYPE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return DistanceType(value)
// 	case VALIDATIONMETHOD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return ValidationMethod(value)
// 	case BOOSTMODE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return BoostMode(value)
// 	case CARD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Card(value)
// 	case MINIMUMSHOULDMATCHFIELD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MinimumShouldMatchField(value)
// 	case FLAGS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Flags(value)
// 	case FORMAT:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Format(value)
// 	case RELATION:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Relation(value)
// 	case VALUE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Value(value)
// 	case FIELD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Field(value)
// 	case ID:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Id(value)
// 	case PARENTTYPE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return ParentType(value)
// 	case SCOREMODE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return ScoreMode(value)
// 	case ANALYZER:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Analyzer(value)
// 	case FUZZYREWRITE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return FuzzyRewrite(value)
// 	case ZEROTERMSQUERY:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return ZeroTermsQuery(value)
// 	case TYPE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Type(value)
// 	case OPERATOR:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Operator(value)
// 	case DEFAULTFIELD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return DefaultField(value)
// 	case DEFAULTOPERATOR:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return DefaultOperator(value)
// 	case QUOTEANALYZER:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return QuoteAnalyzer(value)
// 	case QUOTEFIELDSUFFIX:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return QuoteFieldSuffix(value)
// 	case TIMEZONE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return TimeZone(value)

// 	/* slice */
// 	case FIELDS:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return Fields(fields)
// 	case RULESETIDS:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return RulesetIds(fields)
// 	case DOCUMENTS:
// 		fields := []json.RawMessage{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				jval := json.RawMessage(val)
// 				fields = append(fields, jval)
// 			}
// 		}
// 		return Documents(fields)
// 	case LIKE:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return Like(fields)
// 	case STOPWORDS:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return StopWords(fields)
// 	case UNLIKE:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return Unlike(fields)
// 	case QUERYVECTOR:
// 		fields := []float32{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				v, err := strconv.ParseFloat(val, 32)
// 				if err == nil {
// 					fields = append(fields, float32(v))
// 				}
// 			}
// 		}
// 		return QueryVector(fields)
// 	case SET:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return Set(fields)
// 	case VALUES:
// 		fields := []string{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			val, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				fields = append(fields, val)
// 			}
// 		}
// 		return Values(fields)

// 		/* int,float */
// 	case NEGATIVEBOOST:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return NegativeBoost(v)
// 	case BOOSTTERMS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return BoostTerms(v)
// 	case MAXBOOST:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxBoost(v)
// 	case MINSCORE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return MinScore(v)
// 	case WEIGHT:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return Weight(v)
// 	case TIEBREAKER:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return TieBreaker(v)
// 	case FACTOR:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return Factor(v)
// 	case MISSING:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return Missing(v)
// 	case CUTOFFFREQUENCY:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return CutoffFrequency(v)

// 	case PHRASESLOP:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return PhraseSlop(v)
// 	case VERSION:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.ParseInt(value, 10, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		return Version(v)

// 	case MAXEXPANSIONS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxExpansions(v)
// 	case MAXCONCURRENTGROUPSEARCHES:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxConcurrentGroupSearches(v)
// 	case END:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return End(v)
// 	case NUMCANDIDATES:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return NumCandidates(v)
// 	case MAXCHILDREN:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxChildren(v)
// 	case DIST:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Dist(v)
// 	case POST:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Post(v)
// 	case PRE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Pre(v)
// 	case MINCHILDREN:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MinChildren(v)
// 	case PREFIXLENGTH:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return PrefixLength(v)
// 	case SLOP:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Slop(v)
// 	case FUZZYMAXEXPANSIONS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return FuzzyMaxExpansions(v)
// 	case MAXGAPS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxGaps(v)
// 	case FUZZYPREFIXLENGTH:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return FuzzyPrefixLength(v)
// 	case MAXDETERMINIZEDSTATES:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxDeterminizedStates(v)
// 	case MAXDOCFRED:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxDocFreq(v)
// 	case MAXQUERYTERMS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxQueryTerms(v)
// 	case MAXWORDLENGTH:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MaxWordLength(v)
// 	case MINDOCFREQ:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MinDocFreq(v)
// 	case MINWORDLENGTH:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MinWordLength(v)
// 	case MINTERMFREQ:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		v, err := strconv.Atoi(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MinTermFreq(v)

// 	case BOOST:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		var b Boost
// 		v, err := strconv.ParseFloat(value, 32)
// 		if err != nil {
// 			return nil
// 		}
// 		b = Boost(v)
// 		return b
// 	case SIMILARITY:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}

// 		v, err := strconv.ParseFloat(value, 32)
// 		if err != nil {
// 			return nil
// 		}
// 		return Similarity(v)

// 		/* boolean */
// 	case AUTOGENERATESYNONYMSPHRASEQUERY:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return AutoGenerateSynonymsPhraseQuery(boolValue)
// 	case ALLOWLEADINGWILDCARD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return AllowLeadingWildcard(boolValue)
// 	case ANALYZEWILDCARD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return AnalyzeWildcard(boolValue)
// 	case ESCAPE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Escape(boolValue)

// 	case FUZZYTRANSPOSITIONS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return FuzzyTranspositions(boolValue)
// 	case SCORE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Score(boolValue)
// 	case LENIENT:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Lenient(boolValue)
// 	case TRANSPOSITIONS:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Transpositions(boolValue)
// 	case CASEINSENSITIVE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return CaseInsensitive(boolValue)
// 	case INORDER:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return InOrder(boolValue)
// 	case PRUNE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Prune(boolValue)
// 	case MOREINCLUDE:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return MoreInclude(boolValue)
// 	case FAILONUNSUPPORTEDFIELD:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return FailOnUnsupportedField(boolValue)
// 	case ORDERED:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return Ordered(boolValue)
// 	case INCLUDEUNMAPPED:
// 		value, err := url.QueryUnescape(value)
// 		if err != nil {
// 			return nil
// 		}
// 		boolValue, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return nil
// 		}
// 		return IncludeUnmapped(boolValue)

// 	default:
// 		joining := stringToJoining(key, value)
// 		if joining != nil {
// 			return joining
// 		}
// 		if strings.Contains(value, ";") {
// 			fields := []interface{}{}
// 			vals := strings.Split(value, ";")
// 			if len(vals) == 0 {
// 				return nil
// 			}
// 			for _, v := range vals {
// 				val, err := (url.QueryUnescape(v))
// 				if err == nil {
// 					fields = append(fields, val)
// 				}
// 			}
// 			return fields
// 		} else {
// 			value, err := url.QueryUnescape(value)
// 			if err != nil {
// 				return nil
// 			}
// 			return value
// 		}
// 	}
// }
