package reqparams

import (
	"fmt"
	"strings"
)

type ErrParseFieldsQuery struct {
	Index string
	Char  string
}

func (e ErrParseFieldsQuery) Error() string {
	var str strings.Builder
	str.WriteString(`fields query incorrect format at "{index:`)
	str.WriteString(e.Index)
	str.WriteString(`,value:`)
	str.WriteString(e.Char)
	str.WriteString(`}"`)
	return str.String()
}

type ErrParseFilterQuery ErrParseFieldsQuery

func (e ErrParseFilterQuery) Error() string {
	var str strings.Builder
	str.WriteString(`filter query incorrect format at "{index:`)
	str.WriteString(e.Index)
	str.WriteString(`,value:`)
	str.WriteString(e.Char)
	str.WriteString(`}"`)
	return str.String()
}

type ErrorSort struct {
	At string
}

func (e ErrorSort) Error() string {
	return fmt.Sprintf("query incorrect format at '%v'", e.At)
}

type ErrorFilter struct {
	Index int
	At    string
}

func (e ErrorFilter) Error() string {
	return fmt.Sprintf("filter incorrect format {index:'%v',value:'%v'}", e.Index, e.At)
}

type ErrorQuery struct {
	Index int
	At    string
}

func (e ErrorQuery) Error() string {
	return fmt.Sprintf("query incorrect format {index:'%v',value:'%v'}", e.Index, e.At)
}
