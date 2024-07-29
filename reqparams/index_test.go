package reqparams

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/luongduc1246/ultility/structure"
)

type Ext interface{}

type C struct {
	Comp  string
	Key   string
	Value interface{}
}

var a = "(eq|neq|lt|lte|gt|gte|in|like)"

func TestExp(t *testing.T) {
	s := `eq[phone]=test,like[name]=babaa,in[id]=1;2;3,not[eq[name]=baba,roles[eq[name]=user],or[eq[name]=3]],roles[eq[name]=admin]`
	// regex := regexp.MustCompile(`(^eq|neq|lt|lte|gt|gte|in|like).+?(?:(],))`)
	regex := regexp.MustCompile(`(?P<first>[a-z]+\[)([a-z]+\[[a-z]+\]=.+)*([a-z]+\[.+\])*?(?P<last>(],))`)
	con := regex.FindAllString(s, -1)
	fmt.Println(con)
	fmt.Println(regex.FindAllStringSubmatch(s, 2))
	// for _, e := range con {
	// 	e = strings.Trim(e, ",")
	// 	switch e {
	// 	case condition:
	// a := []Exp{}

	// regex := regexp.MustCompile(`(^eq|neq|lt|lte|gt|gte|in|like).+?(?:,)`)
	// regex := regexp.MustCompile(`not[.+?(?:,)`)
	// con := regex.FindAllString(s, -1)
	// fmt.Println(con)
	// f := `eq:[phone=test],like:[name=babaa],in:[id=1,68],not:[or:[eq:[name=baba]]],roles::[eq:[name=admin],in:[id=1,2,3]]`
	// regex := regexp.MustCompile(`[^(eq)]+:`)
	// con := regex.FindAllString(f, -1)
	// fmt.Println(con)
}

type filter struct {
	name string
	ext  []Ext
	rela []filter
}
type not *filter

func TestMame(t *testing.T) {
	s := `eq[phone]=test,like[name]=babaa,in[id]=1;2;3,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]`
	fil := filter{}
	Mame(s, &fil)

}

func BenchmarkMame(b *testing.B) {
	s := `eq[phone]=test,like[name]=babaa,in[id]=1;2;3,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]`
	for i := 0; i < b.N; i++ {
		fil := filter{}
		Mame(s, &fil)
	}
}

func Mame(s string, f *filter) error {
	stack := structure.NewStack[*filter]()
	stack.Push(f)
	batdau := 0
	mongoac := 0
	dongngoac := 0
	index := 0
	var compare, value, key string
	for i, e := range s {
		switch e {
		case '[':
			mongoac = i
			switch s[batdau:i] {
			case "eq", "like", "in":
			case "not":
				fields := filter{
					name: s[batdau:i],
				}
				var n not
				n = &fields
				batdau = i + 1
				stack.Push(n)
			default:
				fil := filter{
					name: s[batdau:i],
				}
				batdau = i + 1
				stack.Push(&fil)
			}
		case ']':
			if i+1 < len(s) {
				if s[i+1] == '=' {
					dongngoac = i
					index = i + 2
				} else {
					if s[i-1] != ']' {
						compare = s[batdau:mongoac]
						value = s[index:i]
						key = s[mongoac+1 : dongngoac]
						fil, _ := stack.Peek()
						c := C{Comp: compare, Value: value, Key: key}
						fil.ext = append(fil.ext, c)
						f, _ := stack.Pop()
						fil, _ = stack.Peek()
						fil.rela = append(fil.rela, *f)
					} else {
						batdau = i + 2
						f, _ := stack.Pop()
						fil, _ := stack.Peek()
						fil.rela = append(fil.rela, *f)
					}
				}
			} else {
				if s[i-1] != ']' {
					compare = s[batdau:mongoac]
					value = s[index:i]
					key = s[mongoac+1 : dongngoac]
					fil, _ := stack.Peek()
					c := C{Comp: compare, Value: value, Key: key}
					fil.ext = append(fil.ext, c)
					f, _ := stack.Pop()
					fil, _ = stack.Peek()
					fil.rela = append(fil.rela, *f)
				} else {
					f, _ := stack.Pop()
					fil, _ := stack.Peek()
					fil.rela = append(fil.rela, *f)
				}

			}
		case ',':
			if s[i-1] != ']' {
				compare = s[batdau:mongoac]
				value = s[index:i]
				key = s[mongoac+1 : dongngoac]
				batdau = i + 1
				fil, _ := stack.Peek()
				c := C{Comp: compare, Value: value, Key: key}
				fil.ext = append(fil.ext, c)
			}
		}
	}
	return nil
}
