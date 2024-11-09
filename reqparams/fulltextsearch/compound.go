package fulltextsearch

const (
	BOOL               QueryKey = "bool"
	MUST               QueryKey = "must"
	MUSTNOT            QueryKey = "must_not"
	FILTER             QueryKey = "filter"
	SHOULD             QueryKey = "should"
	MINIMUMSHOULDMATCH QueryKey = "minimum_should_match"

	BOOSTING      QueryKey = "boosting"
	NEGATIVE      QueryKey = "negative"
	POSITIVE      QueryKey = "positive"
	NEGATIVEBOOST QueryKey = "negative_boost"
	CONSTANTSCORE QueryKey = "constant_score"

	DISMAX     QueryKey = "dis_max"
	QUERIES    QueryKey = "queries"
	TIEBREAKER QueryKey = "tie_breaker"

	FUNCTIONSCORE    QueryKey = "function_score"
	BOOSTMODE        QueryKey = "boost_mode"
	MAXBOOST         QueryKey = "max_boost"
	MINSCORE         QueryKey = "min_score"
	FUNCTIONS        QueryKey = "functions"
	SCOREMODE        QueryKey = "score_mode"
	WEIGHT           QueryKey = "weight"
	SCRIPTSCORE      QueryKey = "script_score"
	RANDOMSCORE      QueryKey = "random_score"
	EXP              QueryKey = "exp"
	GAUSS            QueryKey = "gauss"
	FIELDVALUEFACTOR QueryKey = "field_value_factor"

	FACTOR   QueryKey = "factor"
	MODIFIER QueryKey = "modifier"
	MISSING  QueryKey = "missing"
	SEED     QueryKey = "seed"
	/* Decay Function */
	UNTYPEDDECAYFUNCTION QueryKey = "untyped_decay_function"
	DATEDECAYFUNCTION    QueryKey = "date_decay_function"
	NUMERICDECAYFUNCTION QueryKey = "numeric_decay_function"
	GEODECAYFUNCTION     QueryKey = "geo_decay_function"
	DECAYPARAMETERS      QueryKey = "decay_parameters"
	DECAY                QueryKey = "decay"
	OFFSET               QueryKey = "offset"
	ORIGIN               QueryKey = "origin"
	SCALE                QueryKey = "scale"
	MULTIVALUEMODE       QueryKey = "multi_value_mode"
)

type Bool struct {
	Querier
}

type MinimumShouldMatch interface{}

type Boosting struct {
	Querier
}

type Negative Boosting

type Positive Boosting

type NegativeBoost float64

type ConstantScore struct {
	Querier
}

type DisMax struct {
	Querier
}

type TieBreaker float64

type FunctionScore struct {
	Querier
}

type ScriptScore FunctionScore
type RandomScore FunctionScore
type FieldValueFactor FunctionScore

type Exp struct {
	Querier
}
type Gauss Exp

type UntypedDecayFunction Exp
type DateDecayFunction Exp
type NumericDecayFunction Exp
type GeoDecayFunction Exp
type DecayParameters Exp

type Decay interface{}
type Offset interface{}
type Origin interface{}
type Scale interface{}

type BoostMode string
type MultiValueMode string
type MaxBoost float64
type MinScore float64
type Factor float64
type Missing float64
type Modifier string
type ScoreMode string
type Seed string
type Weight float64

type Functions struct {
	Params []Querier
}

func NewFunctions() *Functions {
	return &Functions{
		Params: make([]Querier, 0),
	}
}

func (f *Functions) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Functions) GetParams() interface{} {
	return f.Params
}

/* làm việc với filter[query_search[...],query_search[...]] */
type Filter struct {
	Params []Querier
}

func NewFilter() *Filter {
	return &Filter{
		Params: make([]Querier, 0),
	}
}

func (f *Filter) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Filter) GetParams() interface{} {
	return f.Params
}

/* làm việc với must[query_search[...],query_search[...]] */
type Must struct {
	Params []Querier
}

func NewMust() *Must {
	return &Must{
		Params: make([]Querier, 0),
	}
}

func (f *Must) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Must) GetParams() interface{} {
	return f.Params
}

/* làm việc với must_not[query_search[...],query_search[...]] */
type MustNot struct {
	Params []Querier
}

func NewMustNot() *MustNot {
	return &MustNot{
		Params: make([]Querier, 0),
	}
}

func (f *MustNot) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *MustNot) GetParams() interface{} {
	return f.Params
}

/* làm việc với should[query_search[...],query_search[...]] */
type Should struct {
	Params []Querier
}

func NewShould() *Should {
	return &Should{
		Params: make([]Querier, 0),
	}
}

func (f *Should) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Should) GetParams() interface{} {
	return f.Params
}

/* làm việc với should[query_search[...],query_search[...]] */
type Queries struct {
	Params []Querier
}

func NewQueries() *Queries {
	return &Queries{
		Params: make([]Querier, 0),
	}
}

func (f *Queries) AddParam(_ QueryKey, v interface{}) {
	q := v.(Querier)
	f.Params = append(f.Params, q)
}

func (f *Queries) GetParams() interface{} {
	return f.Params
}
