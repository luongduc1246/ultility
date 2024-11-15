package fulltextsearch

/*
	Các giá trị sau nên được url escape
*/

import (
	"net/url"

	"github.com/luongduc1246/ultility/structure"
)

type Querier interface {
	AddParam(i string, v interface{})
	GetParams() interface{}
}

type Query struct {
	Params map[string]interface{}
}

func NewQuery() *Query {
	return &Query{
		Params: make(map[string]interface{}),
	}
}

func (q Query) AddParam(i string, v interface{}) {
	q.Params[i] = v
}
func (q *Query) GetParams() interface{} {
	return q.Params
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

func (q *Query) Parse(s string) error {
	stack := structure.NewStack[Querier]()
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
				result := NewQuery()
				peek.AddParam(s[indexStart:i], result)
				stack.Push(result)
				indexStart = i + 1
			default:
				result := NewQuery()
				peek.AddParam(s[indexStart:i], result)
				stack.Push(result)
				indexStart = i + 1
			}

		case '[':
			result := NewSlice()
			peek, err := stack.Peek()
			if err != nil {
				return err
			}
			peek.AddParam(s[indexStart:i], result)
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
					key := s[indexStart : indexValue-1]
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
					return ErrorQuery{
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
						return ErrorQuery{
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
						key := s[indexStart : indexValue-1]
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
						return ErrorQuery{
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
					key := s[indexStart : indexValue-1]
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
					return ErrorQuery{
						Index: i,
						At:    txtError,
					}
				}
			}
		}
	}
	return nil
}
