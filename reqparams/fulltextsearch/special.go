package fulltextsearch

import "encoding/json"

const (
	DISTANCEFEATURE        QueryKey = "distance_feature"
	MORELIKETHIS           QueryKey = "more_like_this"
	BOOSTTERMS             QueryKey = "boost_terms"
	MOREINCLUDE            QueryKey = "more_include" // thay thế cho include ở trong morethislike
	FAILONUNSUPPORTEDFIELD QueryKey = "fail_on_unsupported_field"
	LIKE                   QueryKey = "like"
	MAXDOCFRED             QueryKey = "max_doc_freq"
	MAXQUERYTERMS          QueryKey = "max_query_terms"
	MAXWORDLENGTH          QueryKey = "max_word_length"
	MINDOCFREQ             QueryKey = "min_doc_freq"
	MINTERMFREQ            QueryKey = "min_term_freq"
	MINWORDLENGTH          QueryKey = "min_word_length"
	ROUTING                QueryKey = "routing"
	STOPWORDS              QueryKey = "stop_words"
	UNLIKE                 QueryKey = "unlike"
	VERSION                QueryKey = "version"
	VERSIONTYPE            QueryKey = "version_type"
	PERCOLATE              QueryKey = "percolate"
	DOCUMENT               QueryKey = "document"
	DOCUMENTS              QueryKey = "documents"
	INDEX                  QueryKey = "index"
	NAME                   QueryKey = "name"
	PREFERENCE             QueryKey = "preference"

	RANKFEATURE   QueryKey = "rank_feature"
	LINEAR        QueryKey = "linear"
	LOG           QueryKey = "log"
	SATURATION    QueryKey = "saturation"
	SIGMOID       QueryKey = "sigmoid"
	SCRIPT        QueryKey = "script"
	LANG          QueryKey = "lang"
	OPTIONS       QueryKey = "options"
	PARAMS        QueryKey = "params"
	SOURCE        QueryKey = "source"
	WRAPPER       QueryKey = "wrapper"
	PINNED        QueryKey = "pinned"
	DOCS          QueryKey = "docs"
	ORGANIC       QueryKey = "organic"
	RULE          QueryKey = "rule"
	MATCHCRITERIA QueryKey = "match_criteria"
	RULESETIDS    QueryKey = "ruleset_ids"
)

type DistanceFeature struct {
	Querier
}

type MoreLikeThis struct {
	Querier
}

type BoostTerms float64

type MoreInclude bool

type FailOnUnsupportedField bool

type Like []string

type MaxDocFreq int
type MaxQueryTerms int
type MaxWordLength int

type MinDocFreq int
type MinWordLength int
type MinTermFreq int
type Routing string

type StopWords []string
type Unlike []string
type Version int64
type VersionType string

type Percolate struct {
	Querier
}

type Document json.RawMessage
type Documents []json.RawMessage
type Index string
type Name string
type Preference string

type RankFeature struct {
	Querier
}
type Log RankFeature
type Linear RankFeature
type Saturation RankFeature
type Sigmoid RankFeature

type Script struct {
	Querier
}

type Options Script
type Params Script
type Source string

type Lang string

type Wrapper struct {
	Querier
}

type Pinned struct {
	Querier
}

type Docs Pinned
type Organic Pinned

type Rule struct {
	Querier
}

type MatchCriteria json.RawMessage

type RulesetIds []string
