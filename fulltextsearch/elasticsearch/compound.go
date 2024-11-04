package elasticsearch

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldvaluefactormodifier"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

type DecayParameters struct {
}

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
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				sliceQuery := []types.Query{}
				for k, q := range options {
					queryNew := fulltextsearch.NewQuerySearch()
					queryNew.AddParam(k, q)
					i := ParseQueryToSearch(queryNew)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Filter = sliceQuery
			case fulltextsearch.MUST:
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				sliceQuery := []types.Query{}
				for k, q := range options {
					queryNew := fulltextsearch.NewQuerySearch()
					queryNew.AddParam(k, q)
					i := ParseQueryToSearch(queryNew)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Must = sliceQuery
			case fulltextsearch.MUSTNOT:
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				sliceQuery := []types.Query{}
				for k, q := range options {
					queryNew := fulltextsearch.NewQuerySearch()
					queryNew.AddParam(k, q)
					i := ParseQueryToSearch(queryNew)
					sliceQuery = append(sliceQuery, *i)
				}
				query.MustNot = sliceQuery
			case fulltextsearch.SHOULD:
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				sliceQuery := []types.Query{}
				for k, q := range options {
					queryNew := fulltextsearch.NewQuerySearch()
					queryNew.AddParam(k, q)
					i := ParseQueryToSearch(queryNew)
					sliceQuery = append(sliceQuery, *i)
				}
				query.Should = sliceQuery
				/* trường hơp sử dụng cho []]fulltextsearch.Querier  */
				// case fulltextsearch.SHOULD:
				// 	field := value.(fulltextsearch.Querier)
				// 	options := field.GetParams().([]fulltextsearch.Querier)
				// 	sliceQuery := []types.Query{}
				// 	for _, q := range options {
				// 		i := ParseQueryToSearch(q)
				// 		sliceQuery = append(sliceQuery, *i)
				// 	}
				// 	query.Should = sliceQuery
			}
		}
	}
	return query
}

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
				field := value.(fulltextsearch.Querier)
				i := ParseQueryToSearch(field)
				query.Negative = i
			case fulltextsearch.POSITIVE:
				field := value.(fulltextsearch.Querier)
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
				field := value.(fulltextsearch.Querier)
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
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().(map[fulltextsearch.QueryKey]interface{})
				sliceQuery := []types.Query{}
				for k, q := range options {
					/* phải tạo một NewQuerySearch để làm việc với hàm ParseQueryToSearch */
					queryNew := fulltextsearch.NewQuerySearch()
					queryNew.AddParam(k, q)
					i := ParseQueryToSearch(queryNew)
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
			case fulltextsearch.QUERY:
				field := value.(fulltextsearch.Querier)
				i := ParseQueryToSearch(field)
				query.Query = i
			case fulltextsearch.FUNCTIONS:
				field := value.(fulltextsearch.Querier)
				options := field.GetParams().([]fulltextsearch.Querier)
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
				v := string(value.(fulltextsearch.BoostMode))
				scoreMode.Name = v
				query.ScoreMode = &scoreMode
			}
		}
	}
	return query
}

func ParseFunctionScore(m fulltextsearch.Querier) *types.FunctionScore {
	query := types.NewFunctionScore()
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.EXP:
				field := value.(fulltextsearch.Querier)
				i := ParseDecayFunction(field)
				query.Exp = i
			case fulltextsearch.FIELDVALUEFACTOR:
				field := value.(fulltextsearch.Querier)
				i := ParseFieldValueFactor(field)
				query.FieldValueFactor = i
			case fulltextsearch.FILTER:
				field := value.(fulltextsearch.Querier)
				i := ParseQueryToSearch(field)
				query.Filter = i
			case fulltextsearch.GAUSS:
				field := value.(fulltextsearch.Querier)
				i := ParseDecayFunction(field)
				query.Gauss = i
			case fulltextsearch.LINEAR:
				field := value.(fulltextsearch.Querier)
				i := ParseDecayFunction(field)
				query.Linear = i
			case fulltextsearch.RANDOMSCORE:
				field := value.(fulltextsearch.Querier)
				i := ParseRandomScore(field)
				query.RandomScore = i
			case fulltextsearch.SCRIPTSCORE:
				field := value.(fulltextsearch.Querier)
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

func ParseDecayFunction(m fulltextsearch.Querier) types.DecayFunction {
	params := m.GetParams()
	switch t := params.(type) {
	case map[fulltextsearch.QueryKey]interface{}:
		for key, value := range t {
			switch key {
			case fulltextsearch.UNTYPEDDECAYFUNCTION:
				field := value.(fulltextsearch.Querier)
				i := ParseUntypedDecayFunction(field)
				return i
			case fulltextsearch.DATEDECAYFUNCTION:
				field := value.(fulltextsearch.Querier)
				i := ParseDateDecayFunction(field)
				return i
			case fulltextsearch.NUMERICDECAYFUNCTION:
				field := value.(fulltextsearch.Querier)
				i := ParseNumericDecayFunction(field)
				return i
			case fulltextsearch.GEODECAYFUNCTION:
				field := value.(fulltextsearch.Querier)
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
			case fulltextsearch.FIELD:
				field := value.(fulltextsearch.Querier)
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
				field := value.(fulltextsearch.Querier)
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
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					f := value.(types.Float64)
					query.Decay = &f
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
				field := value.(fulltextsearch.Querier)
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
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					f := value.(types.Float64)
					query.Decay = &f
				case fulltextsearch.OFFSET:
					f := value.(types.Duration)
					query.Offset = f
				case fulltextsearch.ORIGIN:
					f := value.(string)
					query.Origin = &f
				case fulltextsearch.SCALE:
					f := value.(types.Duration)
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
				field := value.(fulltextsearch.Querier)
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
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					f := value.(types.Float64)
					query.Decay = &f
				case fulltextsearch.OFFSET:
					f := value.(types.Float64)
					query.Offset = &f
				case fulltextsearch.ORIGIN:
					f := value.(types.Float64)
					query.Origin = &f
				case fulltextsearch.SCALE:
					f := value.(types.Float64)
					query.Scale = &f
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
				field := value.(fulltextsearch.Querier)
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
			options := value.(fulltextsearch.Querier)
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[fulltextsearch.QueryKey]interface{})
			for field, value := range mapOptions {
				switch field {
				case fulltextsearch.DECAY:
					f := value.(types.Float64)
					query.Decay = &f
				case fulltextsearch.OFFSET:
					f := value.(string)
					query.Offset = &f
				case fulltextsearch.ORIGIN:
					f := value.(types.GeoLocation)
					query.Origin = &f
				case fulltextsearch.SCALE:
					f := value.(string)
					query.Scale = &f
				}
			}
			base[string(key)] = *query
		}
	}
	return base
}
