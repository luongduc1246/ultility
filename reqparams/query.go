package reqparams

/*
	Các giá trị sau nên được url escape
*/

import (
	"errors"
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
				key, err := url.QueryUnescape(s[indexStart:i])
				if err != nil {
					return err
				}
				peek.AddParam(key, result)
				stack.Push(result)
			default:
				result := NewQuery()
				key, err := url.QueryUnescape(s[indexStart:i])
				if err != nil {
					return err
				}
				peek.AddParam(key, result)
				stack.Push(result)

			}
			indexStart = i + 1
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
			indexStart = i + 1
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
			indexStart = i + 1
			stack.Pop()
		case ',':
			switch s[i-1] {
			case '}', ']':
			default:
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				switch peek.(type) {
				case *Slice:
					value, err := url.QueryUnescape(s[indexStart:i])
					if err != nil {
						return err
					}
					peek.AddParam("", value)
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
			indexStart = i + 1
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

func (q *Slice) Parse(s string) error {
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
				key, err := url.QueryUnescape(s[indexStart:i])
				if err != nil {
					return err
				}
				peek.AddParam(key, result)
				stack.Push(result)
			default:
				result := NewQuery()
				key, err := url.QueryUnescape(s[indexStart:i])
				if err != nil {
					return err
				}
				peek.AddParam(key, result)
				stack.Push(result)

			}
			indexStart = i + 1
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
			indexStart = i + 1
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
			indexStart = i + 1
			stack.Pop()
		case ',':
			switch s[i-1] {
			case '}', ']':
			default:
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				switch peek.(type) {
				case *Slice:
					value, err := url.QueryUnescape(s[indexStart:i])
					if err != nil {
						return err
					}
					peek.AddParam("", value)
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
			indexStart = i + 1
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

func ParseToQuerier(s string) (Querier, error) {
	var querier Querier
	stack := structure.NewStack[Querier]()
	switch s[0] {
	case '{', '[':
		defer stack.Clear()
		var indexStart, indexValue int
		for i, v := range s {
			switch v {
			case '{':
				if querier == nil {
					querier = NewQuery()
					stack.Push(querier)
				} else {
					peek, err := stack.Peek()
					if err != nil {
						return nil, err
					}
					switch peek.(type) {
					case *Slice:
						result := NewQuery()
						key, err := url.QueryUnescape(s[indexStart:i])
						if err != nil {
							return nil, err
						}
						peek.AddParam(key, result)
						stack.Push(result)
					default:
						result := NewQuery()
						key, err := url.QueryUnescape(s[indexStart:i])
						if err != nil {
							return nil, err
						}
						peek.AddParam(key, result)
						stack.Push(result)

					}
				}
				indexStart = i + 1
			case '[':
				if querier == nil {
					querier = NewSlice()
					stack.Push(querier)
				} else {
					result := NewSlice()
					peek, err := stack.Peek()
					if err != nil {
						return nil, err
					}
					key, err := url.QueryUnescape(s[indexStart:i])
					if err != nil {
						return nil, err
					}
					peek.AddParam(key, result)
					stack.Push(result)
				}
				indexStart = i + 1
			case ':':
				indexValue = i + 1
			case '}':
				switch s[i-1] {
				case ']', '}':
					stack.Pop()
				default:
					if (indexStart < indexValue-1) && (indexValue < i) {
						key, err := url.QueryUnescape(s[indexStart : indexValue-1])
						if err != nil {
							return nil, err
						}
						value, err := url.QueryUnescape(s[indexValue:i])
						if err != nil {
							return nil, err
						}
						peek, err := stack.Peek()
						if err != nil {
							return nil, err
						}
						peek.AddParam(key, value)
						stack.Pop()
					} else {
						var txtError string
						if i < indexStart {
							txtError = s[i:indexStart]
						} else {
							txtError = s[indexStart:i]
						}
						return nil, ErrorQuery{
							Index: i,
							At:    txtError,
						}
					}
				}
				indexStart = i + 1
			case ']':
				switch s[i-1] {
				case '}', ']':
				default:
					peek, err := stack.Peek()
					if err != nil {
						return nil, err
					}
					switch peek.(type) {
					case *Slice:
						if indexStart < i {
							value, err := url.QueryUnescape(s[indexStart:i])
							if err != nil {
								return nil, err
							}
							peek.AddParam("", value)
						} else {
							var txtError string
							if i < indexStart {
								txtError = s[i:indexStart]
							} else {
								txtError = s[indexStart:i]
							}
							return nil, ErrorQuery{
								Index: i,
								At:    txtError,
							}
						}
					}
				}
				indexStart = i + 1
				stack.Pop()
			case ',':
				switch s[i-1] {
				case '}', ']':
				default:
					peek, err := stack.Peek()
					if err != nil {
						return nil, err
					}
					switch peek.(type) {
					case *Slice:
						value, err := url.QueryUnescape(s[indexStart:i])
						if err != nil {
							return nil, err
						}
						peek.AddParam("", value)
					default:
						if (indexStart < indexValue-1) && (indexValue < i) {
							key, err := url.QueryUnescape(s[indexStart : indexValue-1])
							if err != nil {
								return nil, err
							}
							value, err := url.QueryUnescape(s[indexValue:i])
							if err != nil {
								return nil, err
							}
							peek.AddParam(key, value)
						} else {
							var txtError string
							if i < indexStart {
								txtError = s[i:indexStart]
							} else {
								txtError = s[indexStart:i]
							}
							return nil, ErrorQuery{
								Index: i,
								At:    txtError,
							}
						}
					}
				}
				indexStart = i + 1
			}
			/*  */
			if i == len(s)-1 {
				switch s[i] {
				case '}', ']':
				default:
					if (indexStart < indexValue-1) && (indexValue < i+1) {
						key, err := url.QueryUnescape(s[indexStart : indexValue-1])
						if err != nil {
							return nil, err
						}
						value, err := url.QueryUnescape(s[indexValue:])
						if err != nil {
							return nil, err
						}
						peek, err := stack.Peek()
						if err != nil {
							return nil, err
						}
						peek.AddParam(key, value)
					} else {
						var txtError string
						if i < indexStart {
							txtError = s[i:indexStart]
						} else {
							txtError = s[indexStart:i]
						}
						return nil, ErrorQuery{
							Index: i,
							At:    txtError,
						}
					}
				}
			}
		}
	default:
		return nil, errors.New("query must be start with '{' or '['")
	}
	return querier, nil
}
