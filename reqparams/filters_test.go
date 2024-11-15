package reqparams

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	s := "eq{phone:test},haskey{name[name,test,pape]},likes{column{value[abc,bcd]}},relative{neq{testneq:valueneq}}"
	expr := NewExpr()
	err := expr.Parse(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(expr)
	f := NewFilter()
	err = f.ParseFromExpr(expr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}
