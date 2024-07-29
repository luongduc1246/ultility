package reqparams

import (
	"errors"
	"regexp"
	"strings"

	"github.com/luongduc1246/ultility/structure"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FieldOld struct {
	Model     interface{} `json:"name"`
	Columns   []string    `json:"fields"`
	Relatives []*FieldOld `json:"relatives"`
}

func ParseQueryToFieldOld(model interface{}, s string) (*FieldOld, error) {
	fields := FieldOld{}
	fields.Model = model
	err := queryToFieldOld(s, &fields)
	if err != nil {
		return nil, err
	}
	return &fields, nil
}

func queryToFieldOld(s string, f *FieldOld) (err error) {
	stack := structure.NewStack[*FieldOld]()
	stack.Push(f)
	defer stack.Clear()
	var indexStart int
	for i, v := range s {
		switch v {
		case '[':
			field := FieldOld{
				Model: cases.Title(language.Und, cases.NoLower).String(s[indexStart:i]),
			}
			indexStart = i + 1
			stack.Push(&field)
		case ']':
			if s[i-1] != ']' {
				field := s[indexStart:i]
				indexStart = i + 1
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				peek.Columns = append(peek.Columns, field)
				pop, err := stack.Pop()
				if err != nil {
					return err
				}
				peek, err = stack.Peek()
				if err != nil {
					return err
				}
				peek.Relatives = append(peek.Relatives, pop)
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
			if s[i-1] != ']' {
				field := s[indexStart:i]
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				peek.Columns = append(peek.Columns, field)
			}
			indexStart = i + 1
		}
		if (i == len(s)-1) && s[i] != ']' {
			field := s[indexStart:]
			peek, err := stack.Peek()
			if err != nil {
				return err
			}
			peek.Columns = append(peek.Columns, field)
		}
	}
	return nil
}

func ParseQueryToFieldsTwo(model interface{}, s string) (*FieldOld, error) {
	sf := FieldOld{}
	sf.Model = model
	s = strings.Trim(s, "[]")
	regex := regexp.MustCompile(`[a-z ]+:[\[\]:, [a-z]+]+`)
	con := regex.FindAllString(s, -1)
	for _, v := range con {
		vs := strings.SplitN(v, ":", 2)
		if len(vs) != 2 {
			return nil, errors.New("asdfd")
		}
		m := cases.Title(language.Und, cases.NoLower).String("vs")
		csf, err := ParseQueryToFieldsTwo(m, vs[1])
		if err != nil {
			return nil, err
		}
		sf.Relatives = append(sf.Relatives, csf)
	}
	s = regex.ReplaceAllString(s, "")
	regex = regexp.MustCompile(`[a-z]+`)
	f := regex.FindAllString(s, -1)
	sf.Columns = f
	return &sf, nil
}
