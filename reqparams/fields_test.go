package reqparams

import (
	"fmt"
	"testing"
)

func TestFieldMap(t *testing.T) {
	a := "phone,roles{uuid,name,users{uuid}}"
	b := "test,roleName{babe,uuid}"
	fm := NewField()
	fm.Parse(a)
	fm.Parse(b)
	scanMap(fm)
}

func scanMap(sf *Field) {
	fmt.Println(sf)
	if sf.Relatives != nil {
		for _, v := range sf.Relatives {
			scanMap(v)
		}
	}
}

func BenchmarkParseQueryToFieldMap(b *testing.B) {
	a := "a,b,c,maz{b,c,dia{a},maz{v,e}},m"
	for i := 0; i < b.N; i++ {
		NewField().Parse(a)
	}
}

func TestParseQueryToFieldOld(t *testing.T) {
	a := "a,b,c,maz{b,c,dia{a}},maz{v,e},m"
	f, _ := ParseQueryToFieldOld(nil, a)
	scan(f)
}

func scan(sf *FieldOld) {
	fmt.Println(sf)
	if sf.Relatives != nil {
		for _, v := range sf.Relatives {
			scan(v)
		}
	}
}

func TestParseQueryToFieldsTwo(t *testing.T) {
	a := "{a,b,c,maz:{b,c,dia:{a},e},m}"
	ParseQueryToFieldsTwo(nil, a)
}

func BenchmarkParseQueryToFieldsTwo(b *testing.B) {
	a := "{a,b,c,maz:{b,c,dia:{a,z,c},e},m}"
	for i := 0; i < b.N; i++ {
		ParseQueryToFieldsTwo(nil, a)
	}
}

func BenchmarkParseQueryToFieldOld(b *testing.B) {
	a := "a,b,c,maz{b,c,dia{a},e},m"
	for i := 0; i < b.N; i++ {
		ParseQueryToFieldOld(nil, a)
	}
}
