package reqparams

import (
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

type ErrParseSortQuery struct {
	Index string
	Char  string
}

func (e ErrParseSortQuery) Error() string {
	var str strings.Builder
	str.WriteString(`sort query incorrect format at "{index:`)
	str.WriteString(e.Index)
	str.WriteString(`,value:`)
	str.WriteString(e.Char)
	str.WriteString(`}"`)
	return str.String()
}
