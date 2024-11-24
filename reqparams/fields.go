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

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Fields struct {
	Columns   []string           `json:"fields"`
	Relatives map[string]*Fields `json:"relatives"`
}

func NewFields() *Fields {
	return &Fields{
		Columns:   make([]string, 0),
		Relatives: make(map[string]*Fields),
	}
}

func (f *Fields) addColumn(s string) {
	if !arrays.Contain(f.Columns, s) {
		f.Columns = append(f.Columns, s)
	}
}

func (f *Fields) ParseFromQuerier(q Querier) error {
	params := q.GetParams()
	switch t := params.(type) {
	case []interface{}:
		for _, v := range t {
			switch c := v.(type) {
			case *Query:
				par := c.Params
				for key, value := range par {
					model := cases.Title(language.Und, cases.NoLower).String(key)
					slice, ok := value.(*Slice)
					if ok {
						newField := NewFields()
						err := newField.ParseFromQuerier(slice)
						if err == nil {
							f.Relatives[model] = newField
						}
					}
				}
			case string:
				f.addColumn(c)
			}
		}
	}
	return nil
}
