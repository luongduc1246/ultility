package reqparams

/*
	- query co dang asc[name],desc[id],roles[asc[name],desc[id]]
*/

import (
	"github.com/luongduc1246/ultility/arrays"
	"github.com/luongduc1246/ultility/structure"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SortKey string

const (
	ASC  SortKey = "asc"
	DESC SortKey = "desc"
)

type Order struct {
	Column string
	Desc   bool
}

func sortFromString(column, by string) Order {
	var order Order
	switch SortKey(by) {
	case ASC:
		order.Column = column
		order.Desc = false
	case DESC:
		order.Column = column
		order.Desc = true
	}
	return order
}

type Sort struct {
	Orders    []Order
	Relatives map[string]*Sort
}

func NewSort() *Sort {
	return &Sort{
		Orders:    make([]Order, 0),
		Relatives: make(map[string]*Sort),
	}
}

func (s *Sort) addOrder(or Order) {
	if !arrays.Contain(s.Orders, or) {
		s.Orders = append(s.Orders, or)
	}
}

func (sort *Sort) Parse(s string) error {
	err := queryToSortMap(s, sort)
	if err != nil {
		return err
	}
	return nil
}

func (sort *Sort) ParseQuerierToSort(q Querier) error {
	params := q.GetParams()
	switch t := params.(type) {
	case []interface{}:
		for _, value := range t {
			switch v := value.(type) {
			case string:
				order := Order{}
				order.Column = v
				sort.addOrder(order)
			case *Query:
				pars := v.Params
				for key, valPars := range pars {
					switch tPars := valPars.(type) {
					case string:
						order := Order{}
						order.Column = key
						switch SortKey(tPars) {
						case ASC:
							order.Desc = false
						case DESC:
							order.Desc = true
						default:
							order.Desc = false
						}
						sort.addOrder(order)
					case *Query:
						if key == "options" {
							for keyOpts, valOpts := range tPars.Params {
								switch tValOpts := valOpts.(type) {
								case *Query:
									fieldOrder, ok := tValOpts.Params["order"]
									if ok {
										orderString, ok := fieldOrder.(string)
										if ok {
											order := Order{}
											order.Column = keyOpts
											switch SortKey(orderString) {
											case ASC:
												order.Desc = false
											case DESC:
												order.Desc = true
											default:
												order.Desc = false
											}
											sort.addOrder(order)
										}

									}
								}
							}
						}
					case *Slice:
						sortNew := NewSort()
						err := sortNew.ParseQuerierToSort(tPars)
						if err == nil {
							model := cases.Title(language.Und, cases.NoLower).String(key)
							sort.Relatives[model] = sortNew
						}
					}
				}
			}
		}
	}
	return nil
}

func queryToSortMap(s string, sort *Sort) (err error) {
	stack := structure.NewStack[*Sort]()
	stack.Push(sort)
	defer stack.Clear()
	var indexStart, indexBracketOpen int
	for i, v := range s {
		switch v {
		case '{':
			indexBracketOpen = i
			switch SortKey(s[indexStart:i]) {
			case ASC, DESC:

			default:
				model := cases.Title(language.Und, cases.NoLower).String(s[indexStart:i])
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				if sortChild, ok := peek.Relatives[model]; ok {
					stack.Push(sortChild)
				} else {
					sortChild = NewSort()
					peek.Relatives[model] = sortChild
					stack.Push(sortChild)
				}
				indexStart = i + 1
			}
		case '}':
			if s[i-1] != '}' {
				if indexStart < indexBracketOpen {
					by := s[indexStart:indexBracketOpen]
					column := s[indexBracketOpen+1 : i]
					order := sortFromString(column, by)
					peek, err := stack.Peek()
					if err != nil {
						return err
					}
					peek.addOrder(order)
				} else {
					var txtError string
					if i < indexStart {
						txtError = s[i:indexStart]
					} else {
						txtError = s[indexStart:i]
					}
					return ErrorSort{
						At: txtError,
					}
				}
			} else {
				_, err := stack.Pop()
				if err != nil {
					return err
				}
			}
		case ',':
			indexStart = i + 1
		}
	}
	return nil
}
