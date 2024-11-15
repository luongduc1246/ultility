package reqparams

/*
	- các giá trị value nên url encode
	- query có dạng eq[adfa]=value1,not[eq[name]=value2,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]
	- Phân tích để lấy column cho câu query where và join trong sql
	- Dùng ngăn xếp để phân tích câu query
	- có lấy field của các quan hệ của các bảng
		- roles[...] là quan hệ trong bảng

	- làm việc với json (extract,haskey làm việc giống như bình thường có dạng extract[name]=blc,haskey)
	likes[field]=(a;b;c:)

*/

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/luongduc1246/ultility/arrays"
	"github.com/luongduc1246/ultility/structure"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FilterKey string

const (
	EQ   FilterKey = "eq"
	NEQ  FilterKey = "neq"
	LT   FilterKey = "lt"
	LTE  FilterKey = "lte"
	GT   FilterKey = "gt"
	GTE  FilterKey = "gte"
	IN   FilterKey = "in"
	LIKE FilterKey = "like"
	AND  FilterKey = "and"
	NOT  FilterKey = "not"
	OR   FilterKey = "or"
	/* làm việc với JSON */
	EXTRACT FilterKey = "extract"
	HASKEY  FilterKey = "haskey"
	EQUALS  FilterKey = "equals"
	LIKES   FilterKey = "likes"
	/* làm việc với JSONARRAY */
	CONTAINS FilterKey = "contains"
)

type Exp interface{}

type Eq struct {
	Column string
	Value  interface{}
}

type Neq Eq

type Lt Eq

type Lte Eq

type Gt Eq

type Gte Eq

type Like Eq

type In struct {
	Column string
	Values []interface{}
}

/* làm việc với Json */
type Extract struct {
	Column string
	Value  string
}
type Contains Eq
type Haskey struct {
	Column string
	Values []string
}

type Likes struct {
	Column string
	Keys   []string
	Value  interface{}
}
type Equals Likes

func expFromString(filterKey, column, value string) Exp {
	var exp Exp
	switch FilterKey(filterKey) {
	case EQ:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Eq{
			Column: column,
			Value:  queryValue,
		}
	case NEQ:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Neq{
			Column: column,
			Value:  queryValue,
		}
	case LT:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Lt{
			Column: column,
			Value:  queryValue,
		}
	case LTE:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Lte{
			Column: column,
			Value:  queryValue,
		}
	case GT:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Gt{
			Column: column,
			Value:  queryValue,
		}
	case GTE:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Gte{
			Column: column,
			Value:  queryValue,
		}
	case LIKE:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Like{
			Column: column,
			Value:  queryValue,
		}
	case IN:
		in := []interface{}{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return exp
		}
		for _, v := range vals {
			queryVal, err := (url.QueryUnescape(v))
			if err == nil {
				in = append(in, queryVal)
			}
		}
		exp = In{
			Column: column,
			Values: in,
		}
	// Làm việc với Json
	case EXTRACT:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Extract{
			Column: column,
			Value:  queryValue,
		}
	case CONTAINS:
		queryValue, err := (url.QueryUnescape(value))
		if err != nil {
			return exp
		}
		exp = Contains{
			Column: column,
			Value:  queryValue,
		}
	case HASKEY:
		has := []string{}
		vals := strings.Split(value, ";")
		if len(vals) == 0 {
			return exp
		}
		for _, v := range vals {
			queryVal, err := (url.QueryUnescape(v))
			if err == nil {
				has = append(has, queryVal)
			}
		}
		exp = Haskey{
			Column: column,
			Values: has,
		}
	case LIKES: /* query có dạng likes[name]=abc;abc:test */
		keyValue := strings.Split(value, ":")
		if len(keyValue) != 2 {
			return exp
		}
		keyLikes := []string{}
		keys := strings.Split(keyValue[0], ";")
		if len(keys) == 0 {
			return exp
		}
		for _, v := range keys {
			queryVal, err := (url.QueryUnescape(v))
			if err == nil {
				keyLikes = append(keyLikes, queryVal)
			}
		}
		queryValue, err := (url.QueryUnescape(keyValue[1]))
		if err != nil {
			return exp
		}
		exp = Likes{
			Column: column,
			Keys:   keyLikes,
			Value:  queryValue,
		}
	case EQUALS:
		keyValue := strings.Split(value, ":")
		if len(keyValue) != 2 {
			return exp
		}
		keyEquals := []string{}
		keys := strings.Split(keyValue[0], ";")
		if len(keys) == 0 {
			return exp
		}
		for _, v := range keys {
			queryVal, err := (url.QueryUnescape(v))
			if err == nil {
				keyEquals = append(keyEquals, queryVal)
			}
		}
		queryValue, err := (url.QueryUnescape(keyValue[1]))
		if err != nil {
			return exp
		}
		exp = Equals{
			Column: column,
			Keys:   keyEquals,
			Value:  queryValue,
		}
	}
	return exp
}

type Filter struct {
	Exps      []Exp // mảng các compare
	Relatives map[string]IFilter
}

func (f *Filter) addExp(exp Exp) {
	f.Exps = append(f.Exps, exp)
}
func (f *Filter) GetExps() []Exp {
	return f.Exps
}
func (f *Filter) GetRelatives() map[string]IFilter {
	return f.Relatives
}

func (f Filter) addRelative(key string, value IFilter) {
	f.Relatives[key] = value
}

func NewFilter() *Filter {
	return &Filter{
		Exps:      make([]Exp, 0),
		Relatives: make(map[string]IFilter),
	}
}

type And struct{ *Filter }

type Not struct{ *Filter }

type Or struct{ *Filter }

type IFilter interface {
	GetExps() []Exp
	GetRelatives() map[string]IFilter
	addExp(Exp)
	addRelative(string, IFilter)
}

/*
#Phân tích từ Exprer sang Filter
*/
func (f *Filter) ParseFromExpr(er Exprer) error {
	params := er.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch FilterKey(key) {
			case EQ:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Eq{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case NEQ:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Neq{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case LT:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Lt{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case LTE:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Lte{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case GT:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Gt{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case GTE:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Gte{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case LIKE:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Like{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case IN:
				/* in{column[value,value]} */
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							sli, ok := v.(*Slice)
							if ok {
								sliI, ok := sli.GetParams().([]interface{})
								if ok {
									exp := In{}
									exp.Column = k
									exp.Values = sliI
									f.Exps = append(f.Exps, exp)
								}

							}
						}
					}
				}
			case EXTRACT:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							s, ok := v.(string)
							if ok {
								exp := Extract{}
								exp.Value = s
								exp.Column = k

								f.Exps = append(f.Exps, exp)
							}
						}
					}
				}
			case CONTAINS:
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							s, ok := v.(string)
							if ok {
								exp := Contains{}
								exp.Value = s
								exp.Column = k

								f.Exps = append(f.Exps, exp)
							}
						}
					}
				}
			case HASKEY:
				/* haskey{column[value,value]} */
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							sli, ok := v.(*Slice)
							if ok {
								sliI, ok := sli.GetParams().([]interface{})
								sliString := arrays.ConvertSliceInterfaceToSliceType[string](sliI)
								if ok {
									exp := Haskey{}
									exp.Column = k
									exp.Values = sliString
									f.Exps = append(f.Exps, exp)
								}

							}
						}
					}
				}
			case LIKES:
				/* likes{column{value[path,path]}} */
				exp, ok := value.(*Expr)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							pare, ok := v.(*Expr)
							if ok {
								parv, ok := pare.GetParams().(map[string]interface{})
								if ok {
									for km, vm := range parv {
										sli, ok := vm.(*Slice)
										if ok {
											sliI, ok := sli.GetParams().([]interface{})
											sliString := arrays.ConvertSliceInterfaceToSliceType[string](sliI)
											if ok {
												exp := Likes{}
												exp.Column = k
												exp.Keys = sliString
												exp.Value = km
												f.Exps = append(f.Exps, exp)
											}

										}
									}
								}
							}
						}
					}
				}
			default:
				/* relative{...expr} */
				fmt.Println(value)
				expr, ok := value.(*Expr)
				if ok {
					filter := NewFilter()
					err := filter.ParseFromExpr(expr)
					if err == nil {
						keyRelative := cases.Title(language.Und, cases.NoLower).String(key)
						f.addRelative(keyRelative, filter)
					}
				}
			}
		}
	}
	return nil
}

// func (f *Filter) Parse(s string) error {
// 	err := queryToFilter(s, f)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func queryToFilter(s string, fields *Filter) (err error) {
// 	stack := structure.NewStack[IFilter]()
// 	stack.Push(fields)
// 	defer stack.Clear()
// 	var indexStart, indexBracketOpen, indexBracketClose, indexValue int
// 	for i, v := range s {
// 		switch v {
// 		case '[':
// 			indexBracketOpen = i
// 			switch FilterKey(s[indexStart:i]) {
// 			case EQ, NEQ, LT, LTE, GT, GTE, LIKE, IN, EXTRACT, HASKEY, CONTAINS, LIKES, EQUALS:
// 			case NOT:
// 				indexStart = i + 1
// 				var not = Not{NewFilter()}
// 				stack.Push(not)
// 			case OR:
// 				indexStart = i + 1
// 				var or = Or{NewFilter()}
// 				stack.Push(or)
// 			case AND:
// 				indexStart = i + 1
// 				var and = And{NewFilter()}
// 				stack.Push(and)
// 			default:
// 				/* làm việc với các quan hệ */
// 				model := cases.Title(language.Und, cases.NoLower).String(s[indexStart:i])
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				if filter, ok := peek.GetRelatives()[model]; ok {
// 					stack.Push(filter)
// 				} else {
// 					filter = NewFilter()
// 					peek.addRelative(model, filter)
// 					stack.Push(filter)
// 				}
// 				indexStart = i + 1
// 			}
// 		case ']':
// 			if i+1 < len(s) {
// 				if s[i+1] == '=' {
// 					indexBracketClose = i
// 					indexValue = i + 2 /* vị trí bắt đầu lấy giá trị value */
// 				} else {
// 					/* kiểm tra xem sau ] */
// 					if s[i-1] != ']' {
// 						if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 							filterKey := s[indexStart:indexBracketOpen]         // get name to filterKey
// 							value := s[indexValue:i]                            // get value for Exp
// 							column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 							exp := expFromString(filterKey, column, value)

// 							indexStart = i + 1
// 							peek, err := stack.Peek()
// 							if err != nil {
// 								return err
// 							}
// 							peek.addExp(exp)
// 							pop, err := stack.Pop()
// 							if err != nil {
// 								return err
// 							}
// 							peek, err = stack.Peek()
// 							if err != nil {
// 								return err
// 							}
// 							peek.addExp(pop)
// 						} else {
// 							return ErrParseFilterQuery{
// 								Index: fmt.Sprint(i),
// 								Char:  string(v),
// 							}
// 						}
// 					} else {
// 						/* trường hợp sau ] là ] */
// 						indexStart = i + 2
// 						pop, err := stack.Pop()
// 						if err != nil {
// 							return err
// 						}
// 						peek, err := stack.Peek()
// 						if err != nil {
// 							return err
// 						}
// 						peek.addExp(pop)
// 					}
// 				}
// 			} else {
// 				if s[i-1] != ']' {
// 					if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 						filterKey := s[indexStart:indexBracketOpen]         // get name to filterKey
// 						value := s[indexValue:i]                            // get value for Exp
// 						column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 						exp := expFromString(filterKey, column, value)
// 						indexStart = i + 1
// 						peek, err := stack.Peek()
// 						if err != nil {
// 							return err
// 						}
// 						peek.addExp(exp)
// 						pop, err := stack.Pop()
// 						if err != nil {
// 							return err
// 						}
// 						peek, err = stack.Peek()
// 						if err != nil {
// 							return err
// 						}
// 						peek.addExp(pop)
// 					} else {
// 						return ErrParseFilterQuery{
// 							Index: fmt.Sprint(i),
// 							Char:  string(v),
// 						}
// 					}
// 				} else {
// 					pop, err := stack.Pop()
// 					if err != nil {
// 						return err
// 					}
// 					peek, err := stack.Peek()
// 					if err != nil {
// 						return err
// 					}
// 					peek.addExp(pop)
// 				}
// 			}
// 		case ',':
// 			if s[i-1] != ']' {
// 				if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 					filterKey := s[indexStart:indexBracketOpen]         // get name to filterKey
// 					value := s[indexValue:i]                            // get value for Exp
// 					column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 					indexStart = i + 1
// 					exp := expFromString(filterKey, column, value)
// 					peek, err := stack.Peek()
// 					if err != nil {
// 						return err
// 					}
// 					peek.addExp(exp)
// 				} else {
// 					return ErrParseFilterQuery{
// 						Index: fmt.Sprint(i),
// 						Char:  string(v),
// 					}
// 				}
// 			}
// 		}
// 		if (i == len(s)-1) && s[i] != ']' {
// 			if indexStart < indexBracketOpen && indexBracketOpen < indexBracketClose {
// 				filterKey := s[indexStart:indexBracketOpen]         // get name to filterKey
// 				value := s[indexValue:]                             // get value for Exp
// 				column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 				exp := expFromString(filterKey, column, value)
// 				peek, err := stack.Peek()
// 				if err != nil {
// 					return err
// 				}
// 				peek.addExp(exp)
// 			} else {
// 				return ErrParseFilterQuery{
// 					Index: fmt.Sprint(i),
// 					Char:  string(v),
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

type Exprer interface {
	AddParam(key string, value interface{})
	GetParams() interface{}
}

type Expr struct {
	Params map[string]interface{}
}

type Slice struct {
	Params []interface{}
}

func NewSlice() *Slice {
	return &Slice{
		Params: make([]interface{}, 0),
	}
}

func (q *Slice) AddParam(i string, v interface{}) {
	q.Params = append(q.Params, v)
}
func (q *Slice) GetParams() interface{} {
	return q.Params
}

func NewExpr() *Expr {
	return &Expr{
		Params: make(map[string]interface{}),
	}
}

func (q Expr) AddParam(i string, v interface{}) {
	q.Params[i] = v
}
func (q *Expr) GetParams() interface{} {
	return q.Params
}

func (q *Expr) Parse(s string) error {
	stack := structure.NewStack[Exprer]()
	stack.Push(q)
	defer stack.Clear()
	var indexStart, indexValue int
	for i, v := range s {
		switch v {
		case '{':
			peek, err := stack.Peek()
			if err != nil {
				return err
			}
			switch peek.(type) {
			case *Slice:
				result := NewExpr()
				key, err := url.QueryUnescape(s[indexStart:i])
				if err != nil {
					return err
				}
				peek.AddParam(key, result)
				stack.Push(result)
				indexStart = i + 1
			default:
				result := NewExpr()
				key, err := url.QueryUnescape(s[indexStart:i])
				if err != nil {
					return err
				}
				peek.AddParam(key, result)
				stack.Push(result)
				indexStart = i + 1
			}

		case '[':
			result := NewSlice()
			peek, err := stack.Peek()
			if err != nil {
				return err
			}
			key, err := url.QueryUnescape(s[indexStart:i])
			if err != nil {
				return err
			}
			peek.AddParam(key, result)
			stack.Push(result)
			indexStart = i + 1
		case ':':
			indexValue = i + 1
		case '}':
			switch s[i-1] {
			case ']', '}':
				stack.Pop()
				indexStart = i + 2
			default:
				if (indexStart < indexValue-1) && (indexValue < i) {
					key, err := url.QueryUnescape(s[indexStart : indexValue-1])
					if err != nil {
						return err
					}
					value, err := url.QueryUnescape(s[indexValue:i])
					if err != nil {
						return err
					}
					peek, err := stack.Peek()
					if err != nil {
						return err
					}
					indexStart = i + 1
					peek.AddParam(key, value)
					stack.Pop()
				} else {
					var txtError string
					if i < indexStart {
						txtError = s[i:indexStart]
					} else {
						txtError = s[indexStart:i]
					}
					return ErrorFilter{
						Index: i,
						At:    txtError,
					}
				}
			}
		case ']':
			switch s[i-1] {
			case '}', ']':
			default:
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				switch peek.(type) {
				case *Slice:
					if indexStart < i {
						value, err := url.QueryUnescape(s[indexStart:i])
						if err != nil {
							return err
						}
						peek.AddParam("", value)
						indexStart = i + 1
					} else {
						var txtError string
						if i < indexStart {
							txtError = s[i:indexStart]
						} else {
							txtError = s[indexStart:i]
						}
						return ErrorFilter{
							Index: i,
							At:    txtError,
						}
					}
				}
			}
			stack.Pop()
		case ',':
			switch s[i-1] {
			case '}', ']':
				indexStart = i + 1
			default:
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				switch peek.(type) {
				case *Slice:
					value := s[indexStart:i]
					peek.AddParam("", value)
					indexStart = i + 1
				default:
					if (indexStart < indexValue-1) && (indexValue < i) {
						key, err := url.QueryUnescape(s[indexStart : indexValue-1])
						if err != nil {
							return err
						}
						value, err := url.QueryUnescape(s[indexValue:i])
						if err != nil {
							return err
						}
						peek.AddParam(key, value)
						indexStart = i + 1
					} else {
						var txtError string
						if i < indexStart {
							txtError = s[i:indexStart]
						} else {
							txtError = s[indexStart:i]
						}
						return ErrorFilter{
							Index: i,
							At:    txtError,
						}
					}
				}
			}
		}
		/*  */
		if i == len(s)-1 {
			switch s[i] {
			case '}', ']':
			default:
				if (indexStart < indexValue-1) && (indexValue < i+1) {
					key, err := url.QueryUnescape(s[indexStart : indexValue-1])
					if err != nil {
						return err
					}
					value, err := url.QueryUnescape(s[indexValue:])
					if err != nil {
						return err
					}
					peek, err := stack.Peek()
					if err != nil {
						return err
					}
					peek.AddParam(key, value)
				} else {
					var txtError string
					if i < indexStart {
						txtError = s[i:indexStart]
					} else {
						txtError = s[indexStart:i]
					}
					return ErrorFilter{
						Index: i,
						At:    txtError,
					}
				}
			}
		}
	}
	return nil
}
