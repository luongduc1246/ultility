package fulltextsearch

const (
	EXISTS                   QueryKey = "exists"
	FIELD                    QueryKey = "field"
	FUZZY                    QueryKey = "fuzzy"
	TRANSPOSITIONS           QueryKey = "transpositions"
	VALUE                    QueryKey = "value"
	IDS                      QueryKey = "ids"
	VALUES                   QueryKey = "values"
	PREFIX                   QueryKey = "prefix"
	PREFIXSTRING             QueryKey = "prefix_string"
	CASEINSENSITIVE          QueryKey = "case_insensitive"
	RANGE                    QueryKey = "range"
	GT                       QueryKey = "gt"
	GTE                      QueryKey = "gte"
	LT                       QueryKey = "lt"
	LTE                      QueryKey = "lte"
	FORMAT                   QueryKey = "format"
	RELATION                 QueryKey = "relation"
	REGEXP                   QueryKey = "regexp"
	FLAGS                    QueryKey = "flags"
	TERM                     QueryKey = "term"
	TERMSTRING               QueryKey = "term_string"
	TERMS                    QueryKey = "terms"
	TERMSSET                 QueryKey = "terms_set"
	SET                      QueryKey = "set" // thay thế cho terms trong terms_set
	MINIMUMSHOULDMATCHFIELD  QueryKey = "minimum_should_match_field"
	MINIMUMSHOULDMATCHSCRIPT QueryKey = "minimum_should_match_script"
	WILDCARD                 QueryKey = "wildcard"
	CARD                     QueryKey = "card" // thay thế cho wildcard trong wild
)

type Exists struct {
	Querier
}
type Field string

type Fuzzy struct {
	Querier
}

type Transpositions bool

type Value string

type Ids struct {
	Querier
}
type Values []string

type Prefix struct {
	Querier
}
type CaseInsensitive bool

type Range struct {
	Querier
}
type Gt interface{}
type Gte interface{}
type Lt interface{}
type Lte interface{}
type Format string
type Relation string

type Regexp struct {
	Querier
}
type Flags string

type Term struct {
	Querier
}

type Terms Term

type TermsSet Term
type MinimumShouldMatchScript Term

type Set []string
type MinimumShouldMatchField string

type Wildcard struct {
	Querier
}

type Card string
type TermString string
type PrefixString string
