package reqparams

/*
	- query có dạng [name,id,roles:[id,name,perm:[name],e],m]
	- Phân tích để lấy column cho câu query select và join trong sql
	- Dùng ngăn xếp để phân tích câu query
	- có lấy field của các quan hệ của các bảng
		- roles[...] là quan hệ trong bảng
*/
import (
	"github.com/luongduc1246/ultility/arrays"
	"github.com/luongduc1246/ultility/structure"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Field struct {
	Columns   []string          `json:"fields"`
	Relatives map[string]*Field `json:"relatives"`
}

func NewField() *Field {
	return &Field{
		Columns:   make([]string, 0),
		Relatives: make(map[string]*Field),
	}
}

func (f *Field) addColumn(s string) {
	if !arrays.Contain(f.Columns, s) {
		f.Columns = append(f.Columns, s)
	}
}

func (f *Field) Parse(s string) error {
	err := queryToField(s, f)
	if err != nil {
		return err
	}
	return nil
}

func queryToField(s string, f *Field) (err error) {
	stack := structure.NewStack[*Field]()
	stack.Push(f)
	defer stack.Clear()
	var indexStart int
	for i, v := range s {
		switch v {
		case '{':
			model := cases.Title(language.Und, cases.NoLower).String(s[indexStart:i])
			peek, err := stack.Peek()
			if err != nil {
				return err
			}
			if field, ok := peek.Relatives[model]; ok {
				stack.Push(field)
			} else {
				field = NewField()
				peek.Relatives[model] = field
				stack.Push(field)
			}
			indexStart = i + 1
		case '}':
			if s[i-1] != ']' {
				column := s[indexStart:i]
				indexStart = i + 1
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				peek.addColumn(column)
				_, err = stack.Pop()
				if err != nil {
					return err
				}
			} else {

				_, err := stack.Pop()
				if err != nil {
					return err
				}
			}
		case ',':
			if s[i-1] != ']' {
				column := s[indexStart:i]
				peek, err := stack.Peek()
				if err != nil {
					return err
				}
				peek.addColumn(column)
			}
			indexStart = i + 1
		}
		if (i == len(s)-1) && s[i] != ']' {
			column := s[indexStart:]
			peek, err := stack.Peek()
			if err != nil {
				return err
			}
			peek.addColumn(column)
		}
	}
	return nil
}
