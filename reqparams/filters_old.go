package reqparams

import (
	"fmt"

	"github.com/luongduc1246/ultility/structure"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FilterOld struct {
	Model     interface{}
	Exps      []Exp // to parse clause condition
	Relatives []*FilterOld
}

type AndOld *FilterOld

type NotOld *FilterOld

type OrOld *FilterOld

func ParseQueryToFilterOld(model interface{}, s string) (*FilterOld, error) {
	filters := FilterOld{}
	filters.Model = model
	err := queryToFilterOld(s, &filters)
	if err != nil {
		return nil, err
	}
	return &filters, nil
}
func queryToFilterOld(s string, fields *FilterOld) (err error) {
	stack := structure.NewStack[*FilterOld]()
	stack.Push(fields)
	defer stack.Clear()
	var indexStart, indexBracketOpen, indexBracketClose, indexValue int
	for i, v := range s {
		switch v {
		case '[':
			indexBracketOpen = i
			switch FilterKey(s[indexStart:i]) {
			case EQ, NEQ, LT, LTE, GT, GTE, LIKE, IN:
			case NOT:
				f := FilterOld{}
				indexStart = i + 1
				var not NotOld
				f.Model = not
				not = &f
				stack.Push(not)
			case OR:
				f := FilterOld{}
				indexStart = i + 1
				var or OrOld
				f.Model = or
				or = &f
				stack.Push(or)
			case AND:
				f := FilterOld{}
				indexStart = i + 1
				var and AndOld
				f.Model = and
				and = &f
				stack.Push(and)
			default:
				fil := FilterOld{
					Model: cases.Title(language.Und, cases.NoLower).String(s[indexStart:i]),
				}
				indexStart = i + 1
				stack.Push(&fil)
			}
		case ']':
			if i+1 < len(s) {
				if s[i+1] == '=' {
					indexBracketClose = i
					indexValue = i + 2
				} else {
					if s[i-1] != ']' {
						if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
							compare := s[indexStart:indexBracketOpen]           // get name to compare
							value := s[indexValue:i]                            // get value for Exp
							column := s[indexBracketOpen+1 : indexBracketClose] // get column
							exp := expFromString(compare, column, value)
							indexStart = i + 1
							peek, err := stack.Peek()
							if err != nil {
								return err
							}
							peek.Exps = append(peek.Exps, exp)
							pop, err := stack.Pop()
							if err != nil {
								return err
							}
							peek, err = stack.Peek()
							if err != nil {
								return err
							}
							addRelaOrExpToFilters(pop, peek)
						} else {
							return ErrParseFilterQuery{
								Index: fmt.Sprint(i),
								Char:  string(v),
							}
						}
					} else {
						indexStart = i + 2
						pop, err := stack.Pop()
						if err != nil {
							return err
						}
						peek, err := stack.Peek()
						if err != nil {
							return err
						}
						addRelaOrExpToFilters(pop, peek)
					}
				}
			} else {
				if s[i-1] != ']' {
					if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
						compare := s[indexStart:indexBracketOpen]           // get name to compare
						value := s[indexValue:i]                            // get value for Exp
						column := s[indexBracketOpen+1 : indexBracketClose] // get column
						exp := expFromString(compare, column, value)
						indexStart = i + 1
						peek, err := stack.Peek()
						if err != nil {
							return err
						}
						peek.Exps = append(peek.Exps, exp)
						pop, err := stack.Pop()
						if err != nil {
							return err
						}
						peek, err = stack.Peek()
						if err != nil {
							return err
						}
						addRelaOrExpToFilters(pop, peek)
					} else {
						return ErrParseFilterQuery{
							Index: fmt.Sprint(i),
							Char:  string(v),
						}
					}
				} else {
					pop, err := stack.Pop()
					if err != nil {
						return err
					}
					peek, err := stack.Peek()
					if err != nil {
						return err
					}
					addRelaOrExpToFilters(pop, peek)
				}
			}
		case ',':
			if s[i-1] != ']' {
				if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
					compare := s[indexStart:indexBracketOpen]           // get name to compare
					value := s[indexValue:i]                            // get value for Exp
					column := s[indexBracketOpen+1 : indexBracketClose] // get column
					indexStart = i + 1
					exp := expFromString(compare, column, value)
					peek, err := stack.Peek()
					if err != nil {
						return err
					}
					peek.Exps = append(peek.Exps, exp)
				} else {
					return ErrParseFilterQuery{
						Index: fmt.Sprint(i),
						Char:  string(v),
					}
				}
			}
		}
		if (i == len(s)-1) && s[i] != ']' {
			if indexStart < indexBracketOpen && indexValue < i && indexBracketOpen < indexBracketClose {
				compare := s[indexStart:indexBracketOpen]           // get name to compare
				value := s[indexValue:]                             // get value for Exp
				column := s[indexBracketOpen+1 : indexBracketClose] // get column
				exp := expFromString(compare, column, value)
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				peek.Exps = append(peek.Exps, exp)
			} else {
				return ErrParseFilterQuery{
					Index: fmt.Sprint(i),
					Char:  string(v),
				}
			}
		}
	}
	return nil
}

func addRelaOrExpToFilters(compare interface{}, origin *FilterOld) {
	f, ok := compare.(*FilterOld)
	if ok {
		switch f.Model.(type) {
		case NotOld:
			not := NotOld(f)
			origin.Exps = append(origin.Exps, not)
		case OrOld:
			or := OrOld(f)
			origin.Exps = append(origin.Exps, or)
		case AndOld:
			and := AndOld(f)
			origin.Exps = append(origin.Exps, and)
		default:
			origin.Relatives = append(origin.Relatives, f)
		}
	}
}
