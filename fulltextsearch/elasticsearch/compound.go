package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldvaluefactormodifier"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
	"github.com/luongduc1246/ultility/reqparams"
)

type DecayParameters struct {
}

/*
Phân tích câu query bool
câu query có dạng bool{filter[{...},{...}}],must[{...}],...}
*/
func ParseBoolQuery(m reqparams.Querier) *types.BoolQuery {
	query := types.NewBoolQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				str, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "minimum_should_match":
				query.MinimumShouldMatch = value
			case "filter":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					if que, ok := q.(reqparams.Querier); ok {
						i := ParseQueryToSearch(que)
						sliceQuery = append(sliceQuery, *i)
					}
				}
				query.Filter = sliceQuery
			case "must":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					if que, ok := q.(reqparams.Querier); ok {
						i := ParseQueryToSearch(que)
						sliceQuery = append(sliceQuery, *i)
					}
				}
				query.Must = sliceQuery
			case "must_not":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					if que, ok := q.(reqparams.Querier); ok {
						i := ParseQueryToSearch(que)
						sliceQuery = append(sliceQuery, *i)
					}
				}
				query.MustNot = sliceQuery
			case "should":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					if que, ok := q.(reqparams.Querier); ok {
						i := ParseQueryToSearch(que)
						sliceQuery = append(sliceQuery, *i)
					}
				}
				query.Should = sliceQuery
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query Boosting
Câu query có dạng boosting{negative{...},positive{...},...}
*/
func ParseBoostingQuery(m reqparams.Querier) *types.BoostingQuery {
	query := types.NewBoostingQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				str, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "negative_boost":
				str, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(str, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.NegativeBoost = vFloat64
			case "negative":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Negative = i
			case "positive":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Positive = i
			case "_name":
				v := value.(string)
				query.QueryName_ = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query constantscore
Câu query có dạng constant_score{filter{...},...}
*/
func ParseConstantScoreQuery(m reqparams.Querier) *types.ConstantScoreQuery {
	query := types.NewConstantScoreQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				str, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "filter":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Filter = i
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query dismax
Câu query có dạng dis_max{queries[{...}],...}
*/
func ParseDisMaxQuery(m reqparams.Querier) *types.DisMaxQuery {
	query := types.NewDisMaxQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				str, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32

			case "queries":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.Query{}
				for _, q := range options {
					if que, ok := q.(reqparams.Querier); ok {
						i := ParseQueryToSearch(que)
						sliceQuery = append(sliceQuery, *i)
					}
				}
				query.Queries = sliceQuery
			case "_name":
				v := value.(string)
				query.QueryName_ = &v
			case "tie_breaker":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.TieBreaker = &vFloat64

			}
		}
	}
	return query
}

/*
Phân tích câu query fuction_score
Câu query có dạng function_score{query{...},functions[{exp{...}],...}
*/
func ParseFunctionScoreQuery(m reqparams.Querier) *types.FunctionScoreQuery {
	query := types.NewFunctionScoreQuery()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 32)
				if err != nil {
					break
				}
				vFloat32 := float32(v)
				query.Boost = &vFloat32
			case "boost_mode":
				boostMode := functionboostmode.FunctionBoostMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				boostMode.Name = v
				query.BoostMode = &boostMode
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "query":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Query = i
			case "functions":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.FunctionScore{}
				for _, q := range options {
					if que, ok := q.(reqparams.Querier); ok {
						i := ParseFunctionScore(que)
						sliceQuery = append(sliceQuery, *i)
					}
				}
				query.Functions = sliceQuery
			case "min_score":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.MinScore = &vFloat64
			case "max_boost":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.MaxBoost = &vFloat64
			case "score_mode":
				scoreMode := functionscoremode.FunctionScoreMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				scoreMode.Name = v
				query.ScoreMode = &scoreMode
			}
		}
	}
	return query
}

/*
Phân tích trường function_score trong Functions trong FunctionScoreQuery

	functions[{...},{...}]
*/
func ParseFunctionScore(m reqparams.Querier) *types.FunctionScore {
	query := types.NewFunctionScore()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "exp":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseDecayFunction(field)
				query.Exp = i
			case "field_value_factor":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseFieldValueFactor(field)
				query.FieldValueFactor = i
			case "filter":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Filter = i
			case "gauss":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseDecayFunction(field)
				query.Gauss = i
			case "linear":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseDecayFunction(field)
				query.Linear = i
			case "random_score":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseRandomScore(field)
				query.RandomScore = i
			case "script_score":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseScriptScore(field)
				query.ScriptScore = i
			case "weight":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Weight = &vFloat64
			}
		}
	}
	return query
}

func ParseRandomScore(m reqparams.Querier) *types.RandomScoreFunction {
	query := types.NewRandomScoreFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = &v
			case "seed":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Seed = v
			}
		}
	}
	return query
}

/*
phân tích script_score
query có dạng script_score{script{...}}
*/
func ParseScriptScore(m reqparams.Querier) *types.ScriptScoreFunction {
	query := types.NewScriptScoreFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "script":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Script = *ParseScript(field)
			}
		}
	}
	return query
}

/*
phân tích field_value_factor
query có dạng field_value_factor{script{...}}
*/

func ParseFieldValueFactor(m reqparams.Querier) *types.FieldValueFactorScoreFunction {
	query := types.NewFieldValueFactorScoreFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "factor":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Factor = &vFloat64
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			case "missing":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Missing = &vFloat64
			case "modifier":
				v, ok := value.(string)
				if !ok {
					break
				}
				mode := fieldvaluefactormodifier.FieldValueFactorModifier{}
				mode.Name = v
				query.Modifier = &mode
			}
		}
	}
	return query
}

/*
Phân tích exp,gauss,linear(decayfunction)

	các dạng decayfunction
	- untyped{}
	- date{}
	- numeric{}
	- geo{}
*/
func ParseDecayFunction(m reqparams.Querier) types.DecayFunction {
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "untyped":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseUntypedDecayFunction(field)
				return i
			case "date":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseDateDecayFunction(field)
				return i
			case "numeric":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseNumericDecayFunction(field)
				return i
			case "geo":
				field, ok := value.(reqparams.Querier)
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

/*
	phân tích untyped
	query : untyped{decay_function_base{},multi_value_mode:testmode}
*/

func ParseUntypedDecayFunction(m reqparams.Querier) *types.UntypedDecayFunction {
	query := types.NewUntypedDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "decay_function_base":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBase = ParseDecayPlacementUntyped(field)
			case "multi_value_mode":
				mode := multivaluemode.MultiValueMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

/*
	phân tích decay_function_base
	decay_function_base{fields{decay:...}}
*/

func ParseDecayPlacementUntyped(m reqparams.Querier) map[string]types.DecayPlacement {
	base := make(map[string]types.DecayPlacement)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacement()
			options, ok := value.(reqparams.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[string]interface{})
			for field, value := range mapOptions {
				switch field {
				case "decay":
					v, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(v, 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case "offset":
					s, ok := value.(string)
					if !ok {
						break
					}
					f := json.RawMessage(s)
					query.Offset = f
				case "origin":
					s, ok := value.(string)
					if !ok {
						break
					}
					f := json.RawMessage(s)
					query.Origin = f
				case "scale":
					s, ok := value.(string)
					if !ok {
						break
					}
					f := json.RawMessage(s)
					query.Scale = f
				}
			}
			base[string(key)] = *query
		}
	}
	return base
}

/*
phân tích date
query : date{decay_function_base{},multi_value_mode:testmode}
*/
func ParseDateDecayFunction(m reqparams.Querier) *types.DateDecayFunction {
	query := types.NewDateDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "decay_function_base":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBaseDateMathDuration = ParseDecayPlacementDate(field)
			case "multi_value_mode":
				mode := multivaluemode.MultiValueMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

/*
phân tích decay_function_base
decay_function_base{fields{decay:...}}
*/
func ParseDecayPlacementDate(m reqparams.Querier) map[string]types.DecayPlacementDateMathDuration {
	base := make(map[string]types.DecayPlacementDateMathDuration)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacementDateMathDuration()
			options, ok := value.(reqparams.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[string]interface{})
			for field, value := range mapOptions {
				switch field {
				case "decay":
					s, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(s, 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case "offset":
					f, ok := value.(types.Duration)
					if !ok {
						break
					}
					query.Offset = f
				case "origin":
					f, ok := value.(string)
					if !ok {
						break
					}
					query.Origin = &f
				case "scale":
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

func ParseNumericDecayFunction(m reqparams.Querier) *types.NumericDecayFunction {
	query := types.NewNumericDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "decay_function_base":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBasedoubledouble = ParseDecayPlacementNumeric(field)
			case "multi_value_mode":
				mode := multivaluemode.MultiValueMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

func ParseDecayPlacementNumeric(m reqparams.Querier) map[string]types.DecayPlacementdoubledouble {
	base := make(map[string]types.DecayPlacementdoubledouble)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacementdoubledouble()
			options, ok := value.(reqparams.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[string]interface{})
			for field, value := range mapOptions {
				switch field {
				case "decay":
					s, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(s, 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case "offset":
					s, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(s, 64); err == nil {
						f := types.Float64(s)
						query.Offset = &f
					}
				case "origin":
					s, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(s, 64); err == nil {
						f := types.Float64(s)
						query.Origin = &f
					}
				case "scale":
					s, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(s, 64); err == nil {
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

func ParseGeoDecayFunction(m reqparams.Querier) *types.GeoDecayFunction {
	query := types.NewGeoDecayFunction()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "decay_function_base":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.DecayFunctionBaseGeoLocationDistance = ParseDecayPlacementGeo(field)
			case "multi_value_mode":
				mode := multivaluemode.MultiValueMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				mode.Name = v
				query.MultiValueMode = &mode
			}
		}
	}
	return query
}

func ParseDecayPlacementGeo(m reqparams.Querier) map[string]types.DecayPlacementGeoLocationDistance {
	base := make(map[string]types.DecayPlacementGeoLocationDistance)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewDecayPlacementGeoLocationDistance()
			options, ok := value.(reqparams.Querier)
			if !ok {
				return nil
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[string]interface{})
			for field, value := range mapOptions {
				switch field {
				case "decay":
					s, ok := value.(string)
					if !ok {
						break
					}
					if s, err := strconv.ParseFloat(s, 64); err == nil {
						f := types.Float64(s)
						query.Decay = &f
					}
				case "offset":
					f, ok := value.(string)
					if !ok {
						break
					}
					query.Offset = &f
				case "origin":
					f, ok := value.(types.GeoLocation)
					if !ok {
						break
					}
					query.Origin = &f
				case "scale":
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
