package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldvaluefactormodifier"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

type DecayParameters struct {
}

/*
Phân tích câu query bool
câu query có dạng bool[filter[query_search[...],query_search[...]],must[query_search[...]],...]
*/
func ParseBoolQuery(m fulltextsearch.Querier) *types.BoolQuery {
	query := types.NewBoolQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.MINIMUMSHOULDMATCH:
				v := value.(fulltextsearch.MinimumShouldMatch)
				query.MinimumShouldMatch = v
			case fulltextsearch.FILTER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					i := ParseQueryToSearch(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Filter = sliceQuery
			case fulltextsearch.MUST:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					i := ParseQueryToSearch(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Must = sliceQuery
			case fulltextsearch.MUSTNOT:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					i := ParseQueryToSearch(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.MustNot = sliceQuery
			case fulltextsearch.SHOULD:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					i := ParseQueryToSearch(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Should = sliceQuery
			}
		}
	}
	return query
}

/*
Phân tích câu query Boosting
Câu query có dạng boosting[negative[...],positive[...],...]
*/
func ParseBoostingQuery(m fulltextsearch.Querier) *types.BoostingQuery {
	query := types.NewBoostingQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.NEGATIVEBOOST:
				v := types.Float64(value.(fulltextsearch.NegativeBoost))
				query.NegativeBoost = v
			case fulltextsearch.NEGATIVE:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Negative = i
			case fulltextsearch.POSITIVE:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Positive = i
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query constantscore
Câu query có dạng constant_score[filter[...],...]
*/
func ParseConstantScoreQuery(m fulltextsearch.Querier) *types.ConstantScoreQuery {
	query := types.NewConstantScoreQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v

			case fulltextsearch.FILTER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Filter = i
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query dismax
Câu query có dạng dis_max[queries[query_search[...]],...]
*/
func ParseDisMaxQuery(m fulltextsearch.Querier) *types.DisMaxQuery {
	query := types.NewDisMaxQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v

			case fulltextsearch.QUERIES:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					i := ParseQueryToSearch(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Queries = sliceQuery
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
			case fulltextsearch.TIEBREAKER:
				v := types.Float64(value.(fulltextsearch.TieBreaker))
				query.TieBreaker = &v

			}
		}
	}
	return query
}

/*
Phân tích câu query FunctionScore
Câu query có dạng function_score[query_search[...],functions[function_score[exp[...],function_score[gauss[...]]]]],...]
  - function_score trong functions[function_score là trường hợp con
  - sử dụng "query_search" để thay thế "query"
*/
func ParseFunctionScoreQuery(m fulltextsearch.Querier) *types.FunctionScoreQuery {
	query := types.NewFunctionScoreQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.BOOST:
				v := float32(value.(fulltextsearch.Boost))
				query.Boost = &v
			case fulltextsearch.BOOSTMODE:
				boostMode := functionboostmode.FunctionBoostMode{}
				v := string(value.(fulltextsearch.BoostMode))
				boostMode.Name = v
				query.BoostMode = &boostMode
			case fulltextsearch.QUERYNAME:
				v := string(value.(fulltextsearch.QueryName))
				query.QueryName_ = &v
				/* sử dụng query_search để thay thế query */
			case fulltextsearch.QUERYSEARCH:
				field := value.(fulltextsearch.Querier)
				i := ParseQueryToSearch(field)
				query.Query = i
			case fulltextsearch.FUNCTIONS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]fulltextsearch.Querier)
				if !ok {
					break
				}
				sliceQuery := []types.FunctionScore{}
				for _, q := range options {
					i := ParseFunctionScore(q)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Functions = sliceQuery
			case fulltextsearch.MINSCORE:
				v := types.Float64(value.(fulltextsearch.MinScore))
				query.MinScore = &v
			case fulltextsearch.MAXBOOST:
				v := types.Float64(value.(fulltextsearch.MaxBoost))
				query.MaxBoost = &v
			case fulltextsearch.SCOREMODE:
				scoreMode := functionscoremode.FunctionScoreMode{}
				v := string(value.(fulltextsearch.ScoreMode))
				scoreMode.Name = v
				query.ScoreMode = &scoreMode
			}
		}
	}
	return query
}

/*
Phân tích trường function_score trong Functions trong FunctionScoreQuery

	functions[function_score[...],...]
*/
func ParseFunctionScore(m fulltextsearch.Querier) *types.FunctionScore {
	query := types.NewFunctionScore()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.EXP:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseDecayFunction(field)
				query.Exp = i
			case fulltextsearch.FIELDVALUEFACTOR:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseFieldValueFactor(field)
				query.FieldValueFactor = i
			case fulltextsearch.FILTER:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Filter = i
			case fulltextsearch.GAUSS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseDecayFunction(field)
				query.Gauss = i
			case fulltextsearch.LINEAR:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseDecayFunction(field)
				query.Linear = i
			case fulltextsearch.RANDOMSCORE:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseRandomScore(field)
				query.RandomScore = i
			case fulltextsearch.SCRIPTSCORE:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseScriptScore(field)
				query.ScriptScore = i
			case fulltextsearch.WEIGHT:
				v := types.Float64(value.(fulltextsearch.Weight))
				query.Weight = &v
			}
		}
	}
	return query
}

/*
Phân tích DecayFunction

	các dạng decayfunction
	- untyped_decay_function[decay_parameters[fiels[decay=cay,offset=off,origin=origin,scale=2]]]
	- date_decay_function[]
	- numeric_decay_function[]
	- geo_decay_function[]
*/
func ParseDecayFunction(m fulltextsearch.Querier) types.DecayFunction {
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.UNTYPEDDECAYFUNCTION:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseUntypedDecayFunction(field)
				return i
			case fulltextsearch.DATEDECAYFUNCTION:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseDateDecayFunction(field)
				return i
			case fulltextsearch.NUMERICDECAYFUNCTION:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseNumericDecayFunction(field)
				return i
			case fulltextsearch.GEODECAYFUNCTION:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				i := ParseGeoDecayFunction(field)
				return i
			}
		}
	}
	return nil
}

func ParseRandomScore(m fulltextsearch.Querier) *types.RandomScoreFunction {
	query := types.NewRandomScoreFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.FIELD:
				v := string(value.(fulltextsearch.Field))
				query.Field = &v
			case fulltextsearch.SEED:
				v := string(value.(fulltextsearch.Seed))
				query.Seed = v
			}
		}
	}
	return query
}
func ParseScriptScore(m fulltextsearch.Querier) *types.ScriptScoreFunction {
	query := types.NewScriptScoreFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.SCRIPT:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.Script = *ParseScript(field)
			}
		}
	}
	return query
}

func ParseFieldValueFactor(m fulltextsearch.Querier) *types.FieldValueFactorScoreFunction {
	query := types.NewFieldValueFactorScoreFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.FACTOR:
				v := types.Float64(value.(fulltextsearch.Factor))
				query.Factor = &v
			case fulltextsearch.FIELD:
				v := string(value.(fulltextsearch.Field))
				query.Field = v
			case fulltextsearch.MISSING:
				v := types.Float64(value.(fulltextsearch.Missing))
				query.Missing = &v
			case fulltextsearch.MODIFIER:
				v := string(value.(fulltextsearch.Modifier))
				mode := fieldvaluefactormodifier.FieldValueFactorModifier{}
				mode.Name = v
				query.Modifier = &mode
			}
		}
	}
	return query
}

func ParseUntypedDecayFunction(m fulltextsearch.Querier) *types.UntypedDecayFunction {
	query := types.NewUntypedDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.DECAYPARAMETERS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBase = ParseDecayPlacementUntyped(field)
			case fulltextsearch.MULTIVALUEMODE:
				mode := multivaluemode.MultiValueMode{}
				v := string(value.(fulltextsearch.MultiValueMode))
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

func ParseDecayPlacementUntyped(m fulltextsearch.Querier) map[string]types.DecayPlacement {
	base := make(map[string]types.DecayPlacement)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacement()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case fulltextsearch.OFFSET:
					f := json.RawMessage(value.(string))
					query.Offset = f
				case fulltextsearch.ORIGIN:
					f := json.RawMessage(value.(string))
					query.Origin = f
				case fulltextsearch.SCALE:
					f := json.RawMessage(value.(string))
					query.Scale = f
				}
			}
			base[string(key)] = *query
		}
	}
	return base
}
func ParseDateDecayFunction(m fulltextsearch.Querier) *types.DateDecayFunction {
	query := types.NewDateDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.DECAYPARAMETERS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBaseDateMathDuration = ParseDecayPlacementDate(field)
			case fulltextsearch.MULTIVALUEMODE:
				mode := multivaluemode.MultiValueMode{}
				v := string(value.(fulltextsearch.MultiValueMode))
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

func ParseDecayPlacementDate(m fulltextsearch.Querier) map[string]types.DecayPlacementDateMathDuration {
	base := make(map[string]types.DecayPlacementDateMathDuration)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacementDateMathDuration()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case fulltextsearch.OFFSET:
					f, ok := value.(types.Duration)
					if !ok {
						break
					}
					query.Offset = f
				case fulltextsearch.ORIGIN:
					f, ok := value.(string)
					if !ok {
						break
					}
					query.Origin = &f
				case fulltextsearch.SCALE:
					f, ok := value.(types.Duration)
					if !ok {
						break
					}
					query.Scale = f
				}
			}
			base[string(key)] = *query
		}
	}
	return base
}

func ParseNumericDecayFunction(m fulltextsearch.Querier) *types.NumericDecayFunction {
	query := types.NewNumericDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.DECAYPARAMETERS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBasedoubledouble = ParseDecayPlacementNumeric(field)
			case fulltextsearch.MULTIVALUEMODE:
				mode := multivaluemode.MultiValueMode{}
				v := string(value.(fulltextsearch.MultiValueMode))
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

func ParseDecayPlacementNumeric(m fulltextsearch.Querier) map[string]types.DecayPlacementdoubledouble {
	base := make(map[string]types.DecayPlacementdoubledouble)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacementdoubledouble()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case fulltextsearch.OFFSET:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Offset = &f
					}
				case fulltextsearch.ORIGIN:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Origin = &f
					}
				case fulltextsearch.SCALE:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Scale = &f
					}
				}
			}
			base[string(key)] = *query
		}
	}
	return base
}

func ParseGeoDecayFunction(m fulltextsearch.Querier) *types.GeoDecayFunction {
	query := types.NewGeoDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.DECAYPARAMETERS:
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBaseGeoLocationDistance = ParseDecayPlacementGeo(field)
			case fulltextsearch.MULTIVALUEMODE:
				mode := multivaluemode.MultiValueMode{}
				v := string(value.(fulltextsearch.MultiValueMode))
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

func ParseDecayPlacementGeo(m fulltextsearch.Querier) map[string]types.DecayPlacementGeoLocationDistance {
	base := make(map[string]types.DecayPlacementGeoLocationDistance)
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacementGeoLocationDistance()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					if s, err := strconv.ParseFloat(value.(string), 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case fulltextsearch.OFFSET:
					f, ok := value.(string)
					if !ok {
						break
					}
					query.Offset = &f
				case fulltextsearch.ORIGIN:
					f, ok := value.(types.GeoLocation)
					if !ok {
						break
					}
					query.Origin = &f
				case fulltextsearch.SCALE:
					f, ok := value.(string)
					if !ok {
						break
					}
					query.Scale = &f
				}
			}
			base[string(key)] = *query
		}
	}
	return base
}
