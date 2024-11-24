package reqparams

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := "id,name,phone,{roles[id,name,babe]}"
	slice := NewSlice()
	err := slice.Parse(s)
	fmt.Println(err)
	showAll(slice)
}
func TestQuery(t *testing.T) {
	s := "eq{phone:0933539091},roles[id,name,babe]"
	query := NewQuery()
	err := query.Parse(s)
	fmt.Println(err)
	showAll(query)
}

func TestParseString(t *testing.T) {
	s := "{abc:ab}"
	quier, err := ParseToQuerier(s)
	fmt.Println(quier, err)

}

func showAll(q Querier) {
	params := q.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			queri, ok := value.(Querier)
			if ok {
				showAll(queri)
			}
			fmt.Println(key, value)
		}
	case []interface{}:
		for _, value := range t {
			queri, ok := value.(Querier)
			if ok {
				showAll(queri)
			}
			fmt.Println("mang", value)
		}
	}
}
