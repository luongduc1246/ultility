package reqparams

// import (
// 	"fmt"
// 	"net/url"
// 	"strings"

// 	"github.com/luongduc1246/ultility/structure"
// )

// type QueryerKey string

// const (
// 	MATCH             QueryerKey = "match"
// 	MULTIMATCH        QueryerKey = "multi-match"
// 	MATCHPHRASE       QueryerKey = "match_phrase"
// 	MATCHPHRASEPREFIX QueryerKey = "match_phrase_prefix"

// 	TERMSSET QueryerKey = "terms_set"
// 	TERM     QueryerKey = "term"

// 	EXISTS   QueryerKey = "exists"
// 	IDS      QueryerKey = "ids"
// 	PREFIX   QueryerKey = "prefix"
// 	WILDCARD QueryerKey = "wildcard"
// 	REGEXP   QueryerKey = "regexp"
// 	FUZZY    QueryerKey = "fuzzy"
// 	BOOL     QueryerKey = "bool"

// 	QUERYSTRING       QueryerKey = "query_string"
// 	SIMPLEQUERYSTRING QueryerKey = "simple_query_string"

// 	RANGE    QueryerKey = "range"
// 	RANGELT  QueryerKey = "lt"
// 	RANGELTE QueryerKey = "lte"
// 	RANGEGT  QueryerKey = "gt"
// 	RANGEGTE QueryerKey = "gte"

// 	MUST    QueryerKey = "must"
// 	FILTER  QueryerKey = "filter"
// 	MUSTNOT QueryerKey = "must_not"
// 	SHOULD  QueryerKey = "should"

// 	NESTED    QueryerKey = "nested"
// 	HASCHILD  QueryerKey = "has_child"
// 	HASPARENT QueryerKey = "has_parent"

// 	BOOSTING QueryerKey = "boosting"
// 	POSITIVE QueryerKey = "positive"
// 	NEGATIVE QueryerKey = "negative"

// 	CONSTANTSCORE QueryerKey = "constant_score"
// 	DISMAX        QueryerKey = "dis_max"
// )

// type Match struct {
// 	Field string
// 	Value interface{}
// }

// type Term Match

// type QueryString Match

// type Fuzzy Match

// type Prefix Match

// type Wildcard Match

// type Regex Match

// type Exits struct {
// 	Field string
// }

// /* ranger */

// type Range struct {
// 	Queryer
// }

// type RangeGt struct {
// 	Value interface{}
// }

// type RangeGte struct {
// 	Value interface{}
// }
// type RangeLt struct {
// 	Value interface{}
// }
// type RangeLte struct {
// 	Value interface{}
// }

// type Ids struct {
// 	Values []interface{}
// }

// /* bool */

// type Bool struct {
// 	Queryer
// }

// type TermsSet struct {
// 	Queryer
// }

// type BoolMust Bool

// type BoolFilter Bool

// type BoolMustNot Bool

// type BoolShould Bool

// /* joining */

// type Nested struct {
// 	Queryer
// }

// type HasChild Nested

// type HasParent Nested

// /* boosting */

// type Boosting struct {
// 	Queryer
// }

// type Negative Boosting
// type Positive Boosting

// type ConstantScore struct {
// 	Queryer
// }
// type DisMax struct {
// 	Queryer
// }

// /* interface cho stack  */
// type Queryer interface {
// 	SetExtra(interface{})
// 	AddQuery(interface{})
// 	GetQuery() []interface{}
// 	GetExtra() interface{}
// }

// /* query để phân tích cho fulltextsearch */
// type Query struct {
// 	Extra interface{}
// 	Q     []interface{}
// }

// func NewQuery() *Query {
// 	return &Query{
// 		Q: make([]interface{}, 0),
// 	}
// }

// func (q *Query) AddQuery(i interface{}) {
// 	q.Q = append(q.Q, i)
// }
// func (q *Query) SetExtra(i interface{}) {
// 	q.Extra = i
// }
// func (q *Query) GetExtra() interface{} {
// 	return q.Extra
// }
// func (q *Query) GetQuery() []interface{} {
// 	return q.Q
// }

// func (q *Query) Parse(s string) error {
// 	s, err := url.QueryUnescape(s)
// 	if err != nil {
// 		return err
// 	}
// 	stack := structure.NewStack[Queryer]()
// 	stack.Push(q)
// 	defer stack.Clear()
// 	var indexStart, indexBracketOpen, indexBracketClose, indexValue int
// 	for i, v := range s {
// 		switch v {
// 		case '[':
// 			indexBracketOpen = i
// 			if s[i-1] == ']' {
// 				indexStart = i + 1
// 			} else {
// 				switch QueryerKey(s[indexStart:i]) {
// 				case BOOL:
// 					indexStart = i + 1
// 					var bool = Bool{NewQuery()}
// 					stack.Push(bool)
// 				case MUST:
// 					indexStart = i + 1
// 					var must = BoolMust{NewQuery()}
// 					stack.Push(must)
// 				case FILTER:
// 					indexStart = i + 1
// 					var filter = BoolFilter{NewQuery()}
// 					stack.Push(filter)
// 				case MUSTNOT:
// 					indexStart = i + 1
// 					var not = BoolMustNot{NewQuery()}
// 					stack.Push(not)
// 				case SHOULD:
// 					indexStart = i + 1
// 					var not = BoolShould{NewQuery()}
// 					stack.Push(not)
// 				case NESTED:
// 					indexStart = i + 1
// 					var nested = Nested{NewQuery()}
// 					stack.Push(nested)
// 				case RANGE:
// 					indexStart = i + 1
// 					var ran = Range{NewQuery()}
// 					stack.Push(ran)
// 				case TERMSSET:
// 					indexStart = i + 1
// 					var ts = TermsSet{NewQuery()}
// 					stack.Push(ts)
// 				case HASCHILD:
// 					indexStart = i + 1
// 					var hasChild = HasChild{NewQuery()}
// 					stack.Push(hasChild)
// 				case BOOSTING:
// 					indexStart = i + 1
// 					var boosting = Boosting{NewQuery()}
// 					stack.Push(boosting)
// 				case NEGATIVE:
// 					indexStart = i + 1
// 					var negative = Negative{NewQuery()}
// 					stack.Push(negative)
// 				case POSITIVE:
// 					indexStart = i + 1
// 					var positive = Positive{NewQuery()}
// 					stack.Push(positive)
// 				case CONSTANTSCORE:
// 					indexStart = i + 1
// 					var conScore = ConstantScore{NewQuery()}
// 					stack.Push(conScore)
// 				case DISMAX:
// 					indexStart = i + 1
// 					var disMax = DisMax{NewQuery()}
// 					stack.Push(disMax)
// 				}
// 			}
// 		case ']':
// 			indexBracketClose = i
// 			/* trường hợp ] ở cuối */
// 			if i+1 == len(s) {
// 				if s[i-1] == ']' {
// 					pop, err := stack.Pop()
// 					if err != nil {
// 						return err
// 					}
// 					peek, err := stack.Peek()
// 					if err != nil {
// 						return err
// 					}
// 					peek.AddQuery(pop)
// 				} else {
// 					if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 						key := s[indexStart:indexBracketOpen]              // get name Query
// 						value := s[indexValue:i]                           // get value Query
// 						field := s[indexBracketOpen+1 : indexBracketClose] // get filed
// 						fmt.Println(key, value, field)
// 						query := stringToQuery(key, field, value)
// 						fmt.Println(query)
// 						indexStart = i + 1
// 						/* lấy stack đầu */
// 						peek, err := stack.Peek()
// 						if err != nil {
// 							return err
// 						}
// 						peek.AddQuery(query)
// 						/* xóa stack đầu sau khi add */
// 						pop, err := stack.Pop()
// 						if err != nil {
// 							return err
// 						}
// 						/* lấy stack sau đó nữa */
// 						peek, err = stack.Peek()
// 						if err != nil {
// 							return err
// 						}
// 						peek.AddQuery(pop)
// 					}
// 				}
// 			} else {
// 				if s[i+1] == '=' {
// 					indexValue = i + 2 /* vị trí bắt đầu lấy giá trị value */
// 				}
// 				// if s[i+1] == '[' {
// 				// 	/* làm việc với Extra ví dụ path trong nested và filed trong range */
// 				// 	peek, err := stack.Peek()
// 				// 	if err != nil {
// 				// 		return err
// 				// 	}
// 				// 	extra := s[indexBracketOpen+1 : indexBracketClose]
// 				// 	peek.SetExtra(extra)
// 				// 	// fmt.Println(extra)
// 				// 	// pop, err := stack.Pop()
// 				// 	// if err != nil {
// 				// 	// 	return err
// 				// 	// }
// 				// 	// /* lấy stack sau đó nữa */
// 				// 	// peek, err = stack.Peek()
// 				// 	// if err != nil {
// 				// 	// 	return err
// 				// 	// }
// 				// 	// peek.AddQuery(pop)
// 				// 	// stack.Push(pop)
// 				// } else {
// 				// 	if s[i-1] != ']' {
// 				// 		if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 				// 			key := s[indexStart:indexBracketOpen]              // get name Query
// 				// 			value := s[indexValue:i]                           // get value Query
// 				// 			field := s[indexBracketOpen+1 : indexBracketClose] // get filed
// 				// 			query := stringToQuery(key, field, value)
// 				// 			indexStart = i + 1
// 				// 			/* lấy stack đầu */
// 				// 			peek, err := stack.Peek()
// 				// 			if err != nil {
// 				// 				return err
// 				// 			}
// 				// 			peek.AddQuery(query)
// 				// 			/* xóa stack đầu sau khi add */
// 				// 			pop, err := stack.Pop()
// 				// 			if err != nil {
// 				// 				return err
// 				// 			}
// 				// 			/* lấy stack sau đó nữa */
// 				// 			peek, err = stack.Peek()
// 				// 			if err != nil {
// 				// 				return err
// 				// 			}
// 				// 			peek.AddQuery(pop)
// 				// 		}
// 				// 	} else {
// 				// 		/* trường hợp trước là ] */
// 				// 		indexStart = i + 2
// 				// 		pop, err := stack.Pop()
// 				// 		if err != nil {
// 				// 			return err
// 				// 		}
// 				// 		peek, err := stack.Peek()
// 				// 		if err != nil {
// 				// 			return err
// 				// 		}
// 				// 		peek.AddQuery(pop)
// 				// 	}
// 				// }
// 			}

// 		case ',':
// 			if s[i-1] != ']' {
// 				if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 					key := s[indexStart:indexBracketOpen]              // get name Query
// 					value := s[indexValue:i]                           // get value Query
// 					field := s[indexBracketOpen+1 : indexBracketClose] // get filed
// 					query := stringToQuery(key, field, value)
// 					indexStart = i + 1
// 					peek, err := stack.Peek()
// 					if err != nil {
// 						return err
// 					}
// 					peek.AddQuery(query)
// 				} else {
// 					return ErrParseFullTextSearchQuery{
// 						Index: fmt.Sprint(i),
// 						Char:  string(v),
// 					}
// 				}
// 			}
// 		}
// 		if (i == len(s)-1) && s[i] != ']' {
// 			if indexStart < indexBracketOpen && indexBracketOpen < indexBracketClose {
// 				key := s[indexStart:indexBracketOpen]              // get name Query
// 				value := s[indexValue:i]                           // get value Query
// 				field := s[indexBracketOpen+1 : indexBracketClose] // get filed
// 				query := stringToQuery(key, field, value)
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.AddQuery(query)
// 			} else {
// 				return ErrParseFullTextSearchQuery{
// 					Index: fmt.Sprint(i),
// 					Char:  string(v),
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func stringToQuery(key string, field string, value string) interface{} {
// 	value, err := url.QueryUnescape(value)
// 	if err != nil {
// 		return nil
// 	}
// 	switch QueryerKey(key) {
// 	case MATCH:
// 		return Match{
// 			Field: field,
// 			Value: value,
// 		}
// 	case TERM:
// 		return Term{
// 			Field: field,
// 			Value: value,
// 		}
// 	case QUERYSTRING:
// 		return Term{
// 			Field: field,
// 			Value: value,
// 		}
// 	case FUZZY:
// 		return Fuzzy{
// 			Field: field,
// 			Value: value,
// 		}
// 	case PREFIX:
// 		return Prefix{
// 			Field: field,
// 			Value: value,
// 		}
// 	case WILDCARD:
// 		return Wildcard{
// 			Field: field,
// 			Value: value,
// 		}
// 	case REGEXP:
// 		return Wildcard{
// 			Field: field,
// 			Value: value,
// 		}
// 	case EXISTS:
// 		return Exits{
// 			Field: value, /* string Query có dạng exits[]=field */
// 		}
// 	case RANGELT:
// 		return RangeLt{
// 			Value: value, /* string Query có dạng lt[]=value */
// 		}
// 	case RANGELTE:
// 		return RangeLte{
// 			Value: value, /* string Query có dạng lte[]=value */
// 		}
// 	case RANGEGT:
// 		return RangeGt{
// 			Value: value, /* string Query có dạng gt[]=value */
// 		}
// 	case RANGEGTE:
// 		return RangeGte{
// 			Value: value, /* string Query có dạng gte[]=value */
// 		}
// 	case IDS:
// 		ids := []interface{}{}
// 		vals := strings.Split(value, ";")
// 		if len(vals) == 0 {
// 			return nil
// 		}
// 		for _, v := range vals {
// 			queryVal, err := (url.QueryUnescape(v))
// 			if err == nil {
// 				ids = append(ids, queryVal)
// 			}
// 		}

// 		return Ids{
// 			Values: ids, /* string Query có dạng ids[]=id1;id2 */
// 		}
// 	}
// 	return nil
// }
