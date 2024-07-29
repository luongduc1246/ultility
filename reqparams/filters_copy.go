package reqparams

/*
	- query có dạng eq[adfa]=a,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]
	- Phân tích để lấy column cho câu query where và join trong sql
	- Dùng ngăn xếp để phân tích câu query
	- có lấy field của các quan hệ của các bảng
		- roles[...] là quan hệ trong bảng
*/

// import (
// 	"fmt"
// 	"net/url"
// 	"strings"
// 	"utility/arrays"
// 	"utility/structure"

// 	"golang.org/x/text/cases"
// 	"golang.org/x/text/language"
// )

// type FilterKey string

// const (
// 	EQ   FilterKey = "eq"
// 	NEQ  FilterKey = "neq"
// 	LT   FilterKey = "lt"
// 	LTE  FilterKey = "lte"
// 	GT   FilterKey = "gt"
// 	GTE  FilterKey = "gte"
// 	IN   FilterKey = "in"
// 	LIKE FilterKey = "like"
// 	AND  FilterKey = "and"
// 	NOT  FilterKey = "not"
// 	OR   FilterKey = "or"
// )

// type Exp interface{}

// type Eq struct {
// 	Column string
// 	Value  interface{}
// }

// type Neq Eq

// type Lt Eq

// type Lte Eq

// type Gt Eq

// type Gte Eq

// type Like Eq

// type In struct {
// 	Column string
// 	Values []interface{}
// }

// func expFromString(compare, column, value string) Exp {
// 	var exp Exp
// 	value, err := (url.QueryUnescape(value))
// 	if err != nil {
// 		return exp
// 	}
// 	switch FilterKey(compare) {
// 	case EQ:
// 		exp = Eq{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case NEQ:
// 		exp = Neq{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case LT:
// 		exp = Lt{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case LTE:
// 		exp = Lte{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case GT:
// 		exp = Gt{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case GTE:
// 		exp = Gte{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case LIKE:
// 		exp = Like{
// 			Column: column,
// 			Value:  value,
// 		}
// 	case IN:
// 		vals := strings.Split(value, ";")
// 		in := arrays.ConvertToSliceInterface(vals)
// 		exp = In{
// 			Column: column,
// 			Values: in,
// 		}
// 	}
// 	return exp
// }

// type Filter struct {
// 	Exps      []Exp // to parse clause condition
// 	Relatives map[string]IFilter
// }

// func (f *Filter) addExp(exp Exp) {
// 	f.Exps = append(f.Exps, exp)
// }
// func (f *Filter) GetExps() []Exp {
// 	return f.Exps
// }
// func (f *Filter) GetRelatives() map[string]IFilter {
// 	return f.Relatives
// }

// func (f Filter) addRelative(key string, value IFilter) {
// 	f.Relatives[key] = value
// }

// func NewFilter() *Filter {
// 	return &Filter{
// 		Exps:      make([]Exp, 0),
// 		Relatives: make(map[string]IFilter),
// 	}
// }

// type And struct{ *Filter }

// type Not struct{ *Filter }

// type Or struct{ *Filter }

// type IFilter interface {
// 	GetExps() []Exp
// 	GetRelatives() map[string]IFilter
// 	addExp(Exp)
// 	addRelative(string, IFilter)
// }

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
// 			case EQ, NEQ, LT, LTE, GT, GTE, LIKE, IN:
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
// 					indexValue = i + 2
// 				} else {
// 					if s[i-1] != ']' {
// 						if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
// 							compare := s[indexStart:indexBracketOpen]           // get name to compare
// 							value := s[indexValue:i]                            // get value for Exp
// 							column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 							exp := expFromString(compare, column, value)

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
// 						compare := s[indexStart:indexBracketOpen]           // get name to compare
// 						value := s[indexValue:i]                            // get value for Exp
// 						column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 						exp := expFromString(compare, column, value)
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
// 					compare := s[indexStart:indexBracketOpen]           // get name to compare
// 					value := s[indexValue:i]                            // get value for Exp
// 					column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 					indexStart = i + 1
// 					exp := expFromString(compare, column, value)
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
// 				compare := s[indexStart:indexBracketOpen]           // get name to compare
// 				value := s[indexValue:]                             // get value for Exp
// 				column := s[indexBracketOpen+1 : indexBracketClose] // get column
// 				exp := expFromString(compare, column, value)
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
