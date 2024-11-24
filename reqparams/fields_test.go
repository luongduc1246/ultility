package reqparams

import (
	"fmt"
	"testing"
)

func TestFields(t *testing.T) {
	s := "id,name,phone,{roles[id,name,babe]}"
	slice := NewSlice()
	err := slice.Parse(s)
	if err != nil {
		return
	}
	f := NewFields()
	f.ParseFromQuerier(slice)
	fmt.Print(f)
}
