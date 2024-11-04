package fulltextsearch

const (
	SPANCONTAINING   QueryKey = "span_containing"
	BIG              QueryKey = "big"
	LITTLE           QueryKey = "little"
	SPANFIELDMASKING QueryKey = "span_field_masking"
	SPANFIRST        QueryKey = "span_first"
	END              QueryKey = "end"
	SPANMULTITERM    QueryKey = "span_multi"
	SPANNEAR         QueryKey = "span_near"
	CLAUSES          QueryKey = "clauses"
	INORDER          QueryKey = "in_order"
	SPANNOT          QueryKey = "span_not"
	EXCLUDE          QueryKey = "exclude"
	INCLUDE          QueryKey = "include"
	POST             QueryKey = "post"
	DIST             QueryKey = "dist"
	PRE              QueryKey = "pre"
	SPANOR           QueryKey = "span_or"
	SPANTERM         QueryKey = "span_term"
	SPANWITHIN       QueryKey = "span_within"
)

type SpanContaining struct {
	Querier
}
type Big SpanContaining
type Little SpanContaining

type SpanFieldMasking struct {
	Querier
}

type SpanFirst struct {
	Querier
}
type End int

type SpanMultiTerm struct {
	Querier
}

type SpanNear struct {
	Querier
}

type Clauses struct {
	Params []Querier
}

func NewClauses() *Clauses {
	return &Clauses{
		Params: make([]Querier, 0),
	}
}

func (f *Clauses) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Clauses) GetParams() interface{} {
	return f.Params
}

type InOrder bool

type SpanNot struct {
	Querier
}

type Exclude SpanNot
type Include SpanNot

type Dist int

type Post int

type Pre int

type SpanOr struct {
	Querier
}
type SpanTerm struct {
	Querier
}
type SpanWithin struct {
	Querier
}
