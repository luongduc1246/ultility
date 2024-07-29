package reqparams

import (
	"fmt"

	"github.com/luongduc1246/ultility/structure"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SortOld struct {
	Model     interface{}
	Orders    []Order
	Relatives []*SortOld
}

func ParseQueryToSortOld(model interface{}, s string) (*SortOld, error) {
	sort := SortOld{}
	sort.Model = model
	err := queryToSortOld(s, &sort)
	if err != nil {
		return nil, err
	}
	return &sort, nil
}

func queryToSortOld(s string, sort *SortOld) (err error) {
	stack := structure.NewStack[*SortOld]()
	stack.Push(sort)
	defer stack.Clear()
	var indexStart, indexBracketOpen int
	for i, v := range s {
		switch v {
		case '[':
			indexBracketOpen = i
			switch SortKey(s[indexStart:i]) {
			case ASC, DESC:

			default:
				sort := SortOld{
					Model: cases.Title(language.Und, cases.NoLower).String(s[indexStart:i]),
				}
				indexStart = i + 1
				stack.Push(&sort)
			}
		case ']':
			if s[i-1] != ']' {
				if indexStart < indexBracketOpen {
					by := s[indexStart:indexBracketOpen]
					column := s[indexBracketOpen+1 : i]
					order := sortFromString(column, by)
					peek, err := stack.Peek()
					if err != nil {
						return err
					}
					peek.Orders = append(peek.Orders, order)
				} else {
					return ErrParseSortQuery{
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
				peek.Relatives = append(peek.Relatives, pop)
			}

		case ',':
			indexStart = i + 1
		}
	}
	return nil
}
