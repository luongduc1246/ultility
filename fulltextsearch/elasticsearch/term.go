package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
	"github.com/luongduc1246/ultility/reqparams/fulltextsearch"
)

/*
Phân tích câu query exists

	exists{boost:3,...}
*/
func ParseExistsQuery(m fulltextsearch.Querier) *types.ExistsQuery {
	query := types.NewExistsQuery()
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
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "field":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Field = v
			}
		}
	}
	return query
}

/*
Phân tích câu query fuzzy

	fuzzy{anylyzer:true,...}
*/
func ParseFuzzyQuery(m fulltextsearch.Querier) map[string]types.FuzzyQuery {
	fuzzyQuery := make(map[string]types.FuzzyQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewFuzzyQuery()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {

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
				case "fuzziness":
					s, ok := value.(string)
					if !ok {
						break
					}
					v := types.Fuzziness(s)
					query.Fuzziness = &v
				case "max_expansions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxExpansions = &v

				case "prefix_length":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.PrefixLength = &v
				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v
				case "transpositions":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.Transpositions = &v
				case "value":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Value = v
				}
			}
			fuzzyQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return fuzzyQuery
}

/*
Phân tích câu query ids

	ids{boost:3,...}
*/
func ParseIdsQuery(m fulltextsearch.Querier) *types.IdsQuery {
	query := types.NewIdsQuery()
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
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "values":
				field, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				fields := make([]string, 0)
				pars, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					s, ok := v.(string)
					if ok {
						fields = append(fields, s)
					}

				}
				query.Values = fields
			}
		}
	}
	return query
}

/*
Phân tích câu query prefix

	prefix{field{anylyzer:true,...}}
*/
func ParsePrefixQuery(m fulltextsearch.Querier) map[string]types.PrefixQuery {
	mapQuery := make(map[string]types.PrefixQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewPrefixQuery()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {

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
				case "case_insensitive":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.CaseInsensitive = &v
				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v

				case "rewrite":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Rewrite = &v

				case "value":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Value = v
				}
			}
			mapQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return mapQuery
}

/*
Phân tích câu query range

	range{fields{type_of_range}}
*/
func ParseRangeQuery(m fulltextsearch.Querier) map[string]types.RangeQuery {
	mapQuery := make(map[string]types.RangeQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			var query types.RangeQuery
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {
				case "untyped":
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query = ParseUntypedRangeQuery(field)
				case "date":
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query = ParseDateRangeQuery(field)
				case "number":
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query = ParseNumberRangeQuery(field)

				case "term":
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query = ParseTermRangeQuery(field)

				}
			}
			mapQuery[string(key)] = query
		}
	default:
		return nil
	}
	return mapQuery
}

/*
Phân tích câu query untyped cua range

	{untyped{...}}
*/
func ParseUntypedRangeQuery(m fulltextsearch.Querier) types.UntypedRangeQuery {
	query := types.UntypedRangeQuery{}
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
			case "format":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Format = &v
			case "from":
				v, ok := value.(string)
				if !ok {
					break
				}
				j := json.RawMessage(v)
				query.From = &j
			case "gt":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Gt = json.RawMessage(v)
			case "gte":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Gte = json.RawMessage(v)
			case "lt":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Lt = json.RawMessage(v)
			case "lte":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Lte = json.RawMessage(v)
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "relation":
				model := rangerelation.RangeRelation{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Relation = &model
			case "time_zone":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.TimeZone = &v
			case "to":
				v, ok := value.(string)
				if !ok {
					break
				}
				j := json.RawMessage(v)
				query.To = &j
			}
		}
	}
	return query
}

/*
Phân tích câu query date cua range

	{date{...}}
*/
func ParseDateRangeQuery(m fulltextsearch.Querier) types.DateRangeQuery {
	query := types.DateRangeQuery{}
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
			case "format":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Format = &v
			case "from":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.From = &v
			case "gt":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Gt = &v
			case "gte":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Gte = &v
			case "lt":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Lt = &v
			case "lte":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Lte = &v
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "relation":
				model := rangerelation.RangeRelation{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Relation = &model
			case "time_zone":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.TimeZone = &v
			case "to":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.To = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query number cua range

	{number{...}}
*/
func ParseNumberRangeQuery(m fulltextsearch.Querier) types.NumberRangeQuery {
	query := types.NumberRangeQuery{}
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
			case "from":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.From = &vFloat64
			case "gt":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Gt = &vFloat64
			case "gte":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Gte = &vFloat64
			case "lt":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Lt = &vFloat64
			case "lte":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.Lte = &vFloat64
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "relation":
				model := rangerelation.RangeRelation{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Relation = &model
			case "to":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					break
				}
				vFloat64 := types.Float64(v)
				query.To = &vFloat64
			}
		}
	}
	return query
}

/*
Phân tích câu query term cua range

	{term{...}}
*/
func ParseTermRangeQuery(m fulltextsearch.Querier) types.TermRangeQuery {
	query := types.TermRangeQuery{}
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
			case "from":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.From = &v
			case "gt":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Gt = &v
			case "gte":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Gte = &v
			case "lt":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Lt = &v
			case "lte":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Lte = &v
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "relation":
				model := rangerelation.RangeRelation{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Relation = &model
			case "to":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.To = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query regexp

	regexp{field{anylyzer:true,...}}
*/
func ParseRegexpQuery(m fulltextsearch.Querier) map[string]types.RegexpQuery {
	mapQuery := make(map[string]types.RegexpQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewRegexpQuery()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {

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
				case "case_insensitive":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.CaseInsensitive = &v
				case "flags":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Flags = &v
				case "max_determinized_states":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxDeterminizedStates = &v
				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v

				case "rewrite":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Rewrite = &v

				case "value":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Value = v
				}
			}
			mapQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return mapQuery
}

/*
Phân tích câu query term

	term{field{boost:3,...}}
*/
func ParseTermQuery(m fulltextsearch.Querier) map[string]types.TermQuery {
	mapQuery := make(map[string]types.TermQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewTermQuery()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {

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
				case "case_insensitive":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.CaseInsensitive = &v

				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v

				case "value":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Value = v
				}
			}
			mapQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return mapQuery
}

/*
Phân tích câu query terms

	terms{boost:3,...}
*/
func ParseTermsQuery(m fulltextsearch.Querier) *types.TermsQuery {
	query := types.NewTermsQuery()
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
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "terms_query":
				quies, ok := value.(fulltextsearch.Querier)
				if !ok {
					break
				}
				query.TermsQuery = ParseMapTermsQuery(quies)
			}
		}
	}
	return query
}

/*
Phân tích câu query terms_query

	term{field{terms_query{id:string,index:string,path:string,routing:string}}} hoặc term{field{terms_query[string,int64,json.RawMessage...]}}
*/
func ParseMapTermsQuery(m fulltextsearch.Querier) map[string]types.TermsQueryField {
	mapQuery := make(map[string]types.TermsQueryField)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch quies := value.(type) {
			case *fulltextsearch.Query:
				mapQuery[key] = ParseFieldLookup(quies)
			case *fulltextsearch.Slice:
				fieldValues := []types.FieldValue{}
				pars, ok := quies.GetParams().([]interface{})
				if !ok {
					break
				}
				for _, v := range pars {
					fieldValues = append(fieldValues, v)
				}
				mapQuery[key] = fieldValues
			}

		}
	default:
		return nil
	}
	return mapQuery
}

/*
Phân tích câu query FieldLookup của TermsQueryField

	...{id:string,index:string,path:string,routing:string}
*/
func ParseFieldLookup(m fulltextsearch.Querier) types.FieldLookup {
	query := types.NewFieldLookup()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "id":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Id = v
			case "index":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Index = &v
			case "path":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Path = &v
			case "routing":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Routing = &v
			}
		}
	}
	return *query
}

/*
Phân tích câu query terms_set

	terms_set{field{boost:3,...}}
*/
func ParseTermsSetQuery(m fulltextsearch.Querier) map[string]types.TermsSetQuery {
	mapQuery := make(map[string]types.TermsSetQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewTermsSetQuery()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {

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
				case "minimum_should_match_field":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.MinimumShouldMatchField = &v
				case "minimum_should_match_script":
					quies, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					query.MinimumShouldMatchScript = ParseScript(quies)
				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v

				case "terms":
					field, ok := value.(fulltextsearch.Querier)
					if !ok {
						break
					}
					fields := make([]string, 0)
					pars, ok := field.GetParams().([]interface{})
					if !ok {
						break
					}
					for _, v := range pars {
						s, ok := v.(string)
						if ok {
							fields = append(fields, s)
						}

					}
					query.Terms = fields
				}
			}
			mapQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return mapQuery
}

/*
Phân tích câu query wildcard

	wildcard{field{boost:3,...}}
*/
func ParseWildcardQuery(m fulltextsearch.Querier) map[string]types.WildcardQuery {
	mapQuery := make(map[string]types.WildcardQuery)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewWildcardQuery()
			options, ok := value.(fulltextsearch.Querier)
			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions, ok := imapOptions.(map[string]interface{})
			if !ok {
				break
			}
			for field, value := range mapOptions {
				switch field {

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
				case "case_insensitive":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.CaseInsensitive = &v

				case "rewrite":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Rewrite = &v
				case "_name":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.QueryName_ = &v

				case "value":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Value = &v
				case "wildcard":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Wildcard = &v
				}
			}
			mapQuery[string(key)] = *query
		}
	default:
		return nil
	}
	return mapQuery
}
