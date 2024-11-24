package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/kuromojitokenizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/language"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptsorttype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snowballlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/luongduc1246/ultility/reqparams"
)

/*
Phân tích câu query nested
câu query có dạng nested{boost=3,...}
*/
func ParseNestedQuery(m reqparams.Querier) *types.NestedQuery {
	query := types.NewNestedQuery()
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
			case "ignore_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IgnoreUnmapped = &v
			case "inner_hits":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseInnerHits(field)
				query.InnerHits = i
			case "query":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Query = i
			case "path":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Path = v
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "score_mode":
				scm := childscoremode.ChildScoreMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				scm.Name = v
				query.ScoreMode = &scm
			}
		}
	}
	return query
}

/*
Phân tích câu query inner_hits

	inner_hits{docvalue_fields[{...},{...}],script_fields{fields{...},message{...}}}
*/
func ParseInnerHits(m reqparams.Querier) *types.InnerHits {
	query := types.NewInnerHits()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "collapse":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseCollapse(field)
				query.Collapse = i
			case "docvalue_fields":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.DocvalueFields = ParseDocValueFields(field)
			case "name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Name = &v
			case "explain":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Explain = &v
			case "fields":
				field, ok := value.(reqparams.Querier)
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
				query.Fields = fields
			case "from":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.From = &v
			case "highlight":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Highlight = ParseHighLight(field)
			case "ignore_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IgnoreUnmapped = &v
			case "script_fields":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.ScriptFields = ParseScriptFields(field)
			case "seq_no_primary_term":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.SeqNoPrimaryTerm = &v
			case "size":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.Size = &v
			case "sort":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Sort = ParseSliceSort(field)
			case "_source":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Source_ = ParseSource(field)
			case "stored_fields":
				field, ok := value.(reqparams.Querier)
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
				query.StoredFields = fields
			case "track_scores":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.TrackScores = &v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Version = &v
			}
		}
	}
	return query
}

/*
Phân tích _source cua inner_hits

	_source[false,true] hoặc _source[{...},{...}]
*/
func ParseSource(v interface{}) types.SourceConfig {
	switch t := v.(type) {
	case *reqparams.Query:
		query := types.SourceFilter{}
		m, ok := t.GetParams().(map[string]interface{})
		if !ok {
			return nil
		}
		for key, value := range m {
			switch key {
			case "excludes":
				field, ok := value.(reqparams.Querier)
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
				query.Excludes = fields
			case "includes":
				field, ok := value.(reqparams.Querier)
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
				query.Includes = fields
			}
		}
		return query
	case string:
		return t
	}
	return nil
}

/*
Phân tích sort cua inner_hits

	sort[false,true] hoặc sort[{...},{...}]
*/
func ParseSliceSort(m reqparams.Querier) []types.SortCombinations {
	sliceQuery := []types.SortCombinations{}
	options, ok := m.GetParams().([]interface{})
	if !ok {
		return nil
	}
	for _, q := range options {
		switch q.(type) {
		case *reqparams.Query:
			v, ok := q.(reqparams.Querier)
			if !ok {
				break
			}
			sort := ParseSortCombinations(v)
			sliceQuery = append(sliceQuery, sort)
		case string:
			sliceQuery = append(sliceQuery, q)
		}
	}
	return sliceQuery
}

/*
		Phân tích sortOptions của sort
	 	 sort[{...},{...}]
*/
func ParseSortCombinations(m reqparams.Querier) *types.SortOptions {
	query := types.NewSortOptions()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "_doc":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Doc_ = ParseScoreSort(field)
			case "_geo_distance":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.GeoDistance_ = ParseGeoDistance(field)
			case "_score":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Score_ = ParseScoreSort(field)
			case "_script":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Script_ = ParseScriptSort(field)
			case "options":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.SortOptions = ParseOptionsSort(field)
			}
		}
	}
	return query
}

/*
		Phân tích options của sort
	 	 sort[{options{fields{...}}}]
*/

func ParseOptionsSort(m reqparams.Querier) map[string]types.FieldSort {
	scriptFields := make(map[string]types.FieldSort)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewFieldSort()
			options, ok := value.(reqparams.Querier)
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
				case "format":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Format = &v
				case "missing":
					v, ok := value.(string)
					if !ok {
						break
					}
					query.Missing = &v
				case "mode":
					model := sortmode.SortMode{}
					v, ok := value.(string)
					if !ok {
						break
					}
					model.Name = v
					query.Mode = &model
				case "nested":
					field, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.Nested = ParseNestedSort(field)
				case "numeric_type":
					model := fieldsortnumerictype.FieldSortNumericType{}
					v, ok := value.(string)
					if !ok {
						break
					}
					model.Name = v
					query.NumericType = &model
				case "order":
					order := sortorder.SortOrder{}
					v, ok := value.(string)
					if !ok {
						break
					}
					order.Name = v
					query.Order = &order
				case "unmapped_type":
					order := fieldtype.FieldType{}
					v, ok := value.(string)
					if !ok {
						break
					}
					order.Name = v
					query.UnmappedType = &order
				}
			}
			scriptFields[key] = *query
		}
	default:
		return nil
	}
	return scriptFields
}

/*
		Phân tích _script của sort
	 	 sort[{_script{...}}]
*/
func ParseScriptSort(m reqparams.Querier) *types.ScriptSort {
	query := types.NewScriptSort()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "mode":
				model := sortmode.SortMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Mode = &model
			case "nested":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Nested = ParseNestedSort(field)
			case "order":
				order := sortorder.SortOrder{}
				v, ok := value.(string)
				if !ok {
					break
				}
				order.Name = v
				query.Order = &order
			case "script":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Script = *ParseScript(field)
			case "type":
				model := scriptsorttype.ScriptSortType{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Type = &model
			}
		}
	}
	return query
}

/*
		Phân tích ScoreSort của sort
	 	 sort[{_doc},{_score}]
*/
func ParseScoreSort(m reqparams.Querier) *types.ScoreSort {
	query := types.NewScoreSort()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "order":
				order := sortorder.SortOrder{}
				v, ok := value.(string)
				if !ok {
					break
				}
				order.Name = v
				query.Order = &order
			}
		}
	}
	return query
}

/*
		Phân tích GeoDistanceSort của sort
	 	 sort[{_geo_distance}]
*/
func ParseGeoDistance(m reqparams.Querier) *types.GeoDistanceSort {
	query := types.NewGeoDistanceSort()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "distance_type":
				model := geodistancetype.GeoDistanceType{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.DistanceType = &model
			case "geo_distance_sort":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.GeoDistanceSort = ParseGeoDistanceSort(field)
			case "ignore_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IgnoreUnmapped = &v
			case "mode":
				model := sortmode.SortMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Mode = &model
			case "nested":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Nested = ParseNestedSort(field)
			case "order":
				order := sortorder.SortOrder{}
				v, ok := value.(string)
				if !ok {
					break
				}
				order.Name = v
				query.Order = &order
			case "unit":
				model := distanceunit.DistanceUnit{}
				v, ok := value.(string)
				if !ok {
					break
				}
				model.Name = v
				query.Unit = &model
			}
		}
	}
	return query
}

/*
phân tích câu query nested cua _geo_distance

	...{nested{...}}
*/
func ParseNestedSort(m reqparams.Querier) *types.NestedSortValue {
	query := types.NewNestedSortValue()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "filter":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Filter = i
			case "max_children":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxChildren = &v
			case "nested":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Nested = ParseNestedSort(field)
			case "path":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Path = v
			}
		}
	}
	return query
}

/*
Phân tích câu query geo_distance_sort

	geo_distance_sort{keys[...],fields[...]}
*/
func ParseGeoDistanceSort(m reqparams.Querier) map[string][]types.GeoLocation {
	mapSliceGeo := make(map[string][]types.GeoLocation)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			field, ok := value.(reqparams.Querier)
			if !ok {
				return nil
			}
			fields := make([]types.GeoLocation, 0)
			pars, ok := field.GetParams().([]interface{})
			if !ok {
				return nil
			}
			for _, v := range pars {
				geoLocation := ParseGeoLocale(v)
				if geoLocation != nil {
					fields = append(fields, geoLocation)
				}
			}
			mapSliceGeo[key] = fields
		}
	}
	return mapSliceGeo
}

/*
Phân tích câu query GeoLocation geo_distance_sort

	geo_distance_sort{keys[string,string]} hoặc ...{keys[{latlon},{geohash}]} hoặc ...{keys[[12,65,45],[45,54,54]]}
*/
func ParseGeoLocale(v interface{}) types.GeoLocation {
	switch t := v.(type) {
	case *reqparams.Query:
		m, ok := t.GetParams().(map[string]interface{})
		if !ok {
			return nil
		}
		if v, ok := m["geohash"]; ok {
			return types.GeoHashLocation{
				Geohash: v.(string),
			}
		}
		lat, latExists := m["lat"]
		lon, lonExists := m["lon"]
		if latExists && lonExists {
			lat64, err := strconv.ParseFloat(lat.(string), 64)
			if err != nil {
				return nil
			}
			lon64, err := strconv.ParseFloat(lon.(string), 64)
			if err != nil {
				return nil
			}
			return types.LatLonGeoLocation{
				Lat: types.Float64(lat64),
				Lon: types.Float64(lon64),
			}
		}
		return nil
	case *reqparams.Slice:
		params, ok := t.GetParams().([]interface{})
		if !ok {
			break
		}
		sliceFloat64 := []types.Float64{}
		for _, value := range params {
			s, ok := value.(string)
			if !ok {
				break
			}
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil
			}
			sliceFloat64 = append(sliceFloat64, types.Float64(v))
		}
		return sliceFloat64
	case string:
		return t
	}
	return nil
}

/*
Phân tích câu query script_field

	...{ignore_failure:false,script{id:3,...}}
*/
func ParseScriptFields(m reqparams.Querier) map[string]types.ScriptField {
	scriptFields := make(map[string]types.ScriptField)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewScriptField()
			options, ok := value.(reqparams.Querier)
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
				case "ignore_failure":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.IgnoreFailure = &v
				case "script":
					field, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.Script = *ParseScript(field)
				}
			}
			scriptFields[key] = *query
		}
	default:
		return nil
	}
	return scriptFields
}

/*
phân tích câu query collapse

	...{collapse[...],slice_inner_hits[inner_hits[...]]]
*/
func ParseCollapse(m reqparams.Querier) *types.FieldCollapse {
	query := types.NewFieldCollapse()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "collapse":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseCollapse(field)
				query.Collapse = i
			case "inner_hits":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options, ok := field.GetParams().([]interface{})
				if !ok {
					break
				}
				sliceQuery := []types.InnerHits{}
				for _, q := range options {
					v, ok := q.(reqparams.Querier)
					if !ok {
						break
					}
					i := ParseInnerHits(v)
					sliceQuery = append(sliceQuery, *i)
				}
				query.InnerHits = sliceQuery
			case "field":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Field = v
			case "max_concurrent_group_searches":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxConcurrentGroupSearches = &v
			}
		}
	}
	return query
}

/*
phân tích câu query doc_value_fields

	doc_value_fields[{field:abc,format:abc},{...}]
*/
func ParseDocValueFields(m reqparams.Querier) []types.FieldAndFormat {
	sliceQuery := []types.FieldAndFormat{}
	options, ok := m.GetParams().([]interface{})
	if !ok {
		return nil
	}
	for _, q := range options {
		v, ok := q.(reqparams.Querier)
		if !ok {
			return nil
		}
		i := ParseFieldAndFormat(v)
		sliceQuery = append(sliceQuery, *i)
	}
	return sliceQuery
}

/*
phân tích câu query field_and_format

	...{field=abc,format=abc}
*/
func ParseFieldAndFormat(m reqparams.Querier) *types.FieldAndFormat {
	query := types.NewFieldAndFormat()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "field":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Field = v
			case "format":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Field = v
			case "include_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IncludeUnmapped = &v
			}
		}
	}
	return query
}

/*
phân tích câu query highlight

	highlight{options{a:b,b:c}}
*/
func ParseHighLight(m reqparams.Querier) *types.Highlight {
	query := types.NewHighlight()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "boundary_chars":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.BoundaryChars = &v
			case "boundary_max_scan":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.BoundaryMaxScan = &v
			case "boundary_scanner":
				b := boundaryscanner.BoundaryScanner{}
				v, ok := value.(string)
				if !ok {
					break
				}
				b.Name = v
				query.BoundaryScanner = &b
			case "boundary_scanner_locale":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.BoundaryScannerLocale = &v
			case "encoder":
				b := highlighterencoder.HighlighterEncoder{}
				v, ok := value.(string)
				if !ok {
					break
				}
				b.Name = v
				query.Encoder = &b
			case "fields":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.Fields = ParseHighLightFields(field)
			case "force_source":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.ForceSource = &v
			case "fragment_size":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.FragmentSize = &v
			case "fragmenter":
				b := highlighterfragmenter.HighlighterFragmenter{}
				v, ok := value.(string)
				if !ok {
					break
				}
				b.Name = v
				query.Fragmenter = &b
			case "highlight_filter":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.HighlightFilter = &v
			case "highlight_query":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				query.HighlightQuery = ParseQueryToSearch(field)
			case "max_analyzed_offset":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxAnalyzedOffset = &v
			case "max_fragment_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxFragmentLength = &v
			case "no_match_size":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.NoMatchSize = &v
			case "number_of_fragments":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.NumberOfFragments = &v
			case "options":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				options := make(map[string]json.RawMessage)
				pars, ok := field.GetParams().(map[string]interface{})
				if !ok {
					break
				}
				for k, v := range pars {
					s, ok := v.(string)
					if ok {
						strValue := json.RawMessage(s)
						options[k] = strValue
					}
				}
				query.Options = options
			case "order":
				b := highlighterorder.HighlighterOrder{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				b.Name = v
				query.Order = &b
			case "phrase_limit":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.PhraseLimit = &v
			case "post_tags":
				field, ok := value.(reqparams.Querier)

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
				query.PostTags = fields
			case "pre_tags":
				field, ok := value.(reqparams.Querier)
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
				query.PreTags = fields
			case "require_field_match":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.RequireFieldMatch = &v
			case "tags_schema":
				b := highlightertagsschema.HighlighterTagsSchema{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				b.Name = v
				query.TagsSchema = &b
			case "type":
				b := highlightertype.HighlighterType{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				b.Name = v
				query.Type = &b
			}
		}
	}
	return query
}

/*
Phân tích câu query hightlight_fields *

	...{analyzer{}}
*/
func ParseHighLightFields(m reqparams.Querier) map[string]types.HighlightField {
	hight := make(map[string]types.HighlightField)
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			query := types.NewHighlightField()
			options, ok := value.(reqparams.Querier)

			if !ok {
				break
			}
			imapOptions := options.GetParams()
			mapOptions := imapOptions.(map[string]interface{})
			for field, value := range mapOptions {
				switch field {
				case "boundary_chars":
					s, ok := value.(string)
					if !ok {
						break
					}
					v := s
					query.BoundaryChars = &v
				case "boundary_max_scan":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.BoundaryMaxScan = &v
				case "boundary_scanner":
					s, ok := value.(string)
					if !ok {
						break
					}
					b := boundaryscanner.BoundaryScanner{}
					v := s
					b.Name = v
					query.BoundaryScanner = &b
				case "boundary_scanner_locale":
					s, ok := value.(string)
					if !ok {
						break
					}
					v := s
					query.BoundaryScannerLocale = &v
				case "force_source":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.ForceSource = &v
				case "fragment_offset":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.FragmentOffset = &v
				case "fragment_size":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.FragmentSize = &v
				case "fragmenter":
					b := highlighterfragmenter.HighlighterFragmenter{}
					s, ok := value.(string)
					if !ok {
						break
					}
					v := s
					b.Name = v
					query.Fragmenter = &b
				case "highlight_filter":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.HighlightFilter = &v
				case "highlight_query":
					field, ok := value.(reqparams.Querier)
					if !ok {
						break
					}
					query.HighlightQuery = ParseQueryToSearch(field)
				case "matched_fields":
					field, ok := value.(reqparams.Querier)

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
					query.MatchedFields = fields
				case "max_analyzed_offset":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxAnalyzedOffset = &v
				case "max_fragment_length":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.MaxFragmentLength = &v
				case "no_match_size":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.NoMatchSize = &v
				case "number_of_fragments":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.NumberOfFragments = &v
				case "options":
					field, ok := value.(reqparams.Querier)

					if !ok {
						break
					}
					options := make(map[string]json.RawMessage)
					pars, ok := field.GetParams().(map[string]interface{})
					if !ok {
						break
					}
					for k, v := range pars {
						s, ok := v.(string)
						if ok {
							strValue := json.RawMessage(s)
							options[k] = strValue
						}

					}
					query.Options = options
				case "order":
					b := highlighterorder.HighlighterOrder{}
					s, ok := value.(string)
					if !ok {
						break
					}
					v := s
					b.Name = v
					query.Order = &b
				case "phrase_limit":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.Atoi(s)
					if err != nil {
						break
					}
					query.PhraseLimit = &v
				case "post_tags":
					field, ok := value.(reqparams.Querier)

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
					query.PostTags = fields
				case "pre_tags":
					field, ok := value.(reqparams.Querier)
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
					query.PreTags = fields
				case "require_field_match":
					s, ok := value.(string)
					if !ok {
						break
					}
					v, err := strconv.ParseBool(s)
					if err != nil {
						break
					}
					query.RequireFieldMatch = &v
				case "tags_schema":
					b := highlightertagsschema.HighlighterTagsSchema{}
					s, ok := value.(string)
					if !ok {
						break
					}
					v := s
					b.Name = v
					query.TagsSchema = &b
				case "type":
					b := highlightertype.HighlighterType{}
					s, ok := value.(string)
					if !ok {
						break
					}
					v := s
					b.Name = v
					query.Type = &b

				}
			}
			hight[string(key)] = *query
		}
	default:
		return nil
	}
	return hight
}

/*
Phân tích analyzer của highlight

	 highlight_analyzer{custom{}...},...}
		các dạng highlight_analyzer
		- custom
		- finger_print
		- keyword
		- language
		- nori
		- pattern
		- simple
		- standar
		- stop
		- white_space
		- icu
		- kuromoji
		- snow_ball
		- dutch
*/
func ParseHighlightAnalyzer(m reqparams.Querier) types.Analyzer {
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "custom":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseCustomAnalyzer(field)
				return i
			case "finger_print":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseFingerPrintAnalyzer(field)
				return i
			case "keyword":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseKeywordAnalyzer(field)
				return i
			case "language":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseLanguageAnalyzer(field)
				return i
			case "nori":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseNoriAnalyzer(field)
				return i
			case "pattern":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParsePatternAnalyzer(field)
				return i
			case "simple":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseSimpleAnalyzer(field)
				return i
			case "standard":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseStandardAnalyzer(field)
				return i
			case "stop":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseStopAnalyzer(field)
				return i
			case "white_space":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseWhiteSpaceAnalyzer(field)
				return i
			case "icu":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseIcuAnalyzer(field)
				return i
			case "kuromoji":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseKuromojiAnalyzer(field)
				return i
			case "snow":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseSnowballAnalyzer(field)
				return i
			case "dutch":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseDutchAnalyzer(field)
				return i
			}
		}
	}
	return nil
}

/*
phân tích câu query custom của analyzer

	custom{char_filter[]...}
*/
func ParseCustomAnalyzer(m reqparams.Querier) *types.CustomAnalyzer {
	query := types.NewCustomAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "char_filter":
				field, ok := value.(reqparams.Querier)

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
				query.CharFilter = fields
			case "filter":
				field, ok := value.(reqparams.Querier)

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
				query.Filter = fields
			case "position_increment_gap":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.PositionIncrementGap = &v
			case "position_offset_gap":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.PositionOffsetGap = &v
			case "tokenizer":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Tokenizer = v
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query stop cua analyzer

	stop{...}
*/
func ParseStopAnalyzer(m reqparams.Querier) *types.StopAnalyzer {
	query := types.NewStopAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "stopwords":
				field, ok := value.(reqparams.Querier)
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
				query.Stopwords = fields
			case "stopwords_path":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.StopwordsPath = &v
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query finger_print cua analyzer

	finger_print{...}
*/
func ParseFingerPrintAnalyzer(m reqparams.Querier) *types.FingerprintAnalyzer {
	query := types.NewFingerprintAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "max_output_size":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxOutputSize = v
			case "preserve_original":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.PreserveOriginal = v
			case "separator":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Separator = v
			case "stopwords":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields
			case "stopwords_path":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.StopwordsPath = &v
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query language cua analyzer

	language{...}
*/
func ParseLanguageAnalyzer(m reqparams.Querier) *types.LanguageAnalyzer {
	query := types.NewLanguageAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "language":
				lang := language.Language{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				lang.Name = v
				query.Language = lang
			case "stem_exclusion":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields
			case "stopwords":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields
			case "stopwords_path":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.StopwordsPath = &v
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query snowball cua analyzer

	snowball{...}
*/
func ParseSnowballAnalyzer(m reqparams.Querier) *types.SnowballAnalyzer {
	query := types.NewSnowballAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "language":
				lang := snowballlanguage.SnowballLanguage{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				lang.Name = v
				query.Language = lang
			case "stem_exclusion":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields
			case "stopwords":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields

			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query dutch cua analyzer

	dutch{...}
*/
func ParseDutchAnalyzer(m reqparams.Querier) *types.DutchAnalyzer {
	query := types.NewDutchAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "stopwords":
				field, ok := value.(reqparams.Querier)
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
				query.Stopwords = fields

			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query pattern cua analyzer

	pattern{...}
*/
func ParsePatternAnalyzer(m reqparams.Querier) *types.PatternAnalyzer {
	query := types.NewPatternAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "flags":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Flags = &v
			case "lowercase":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Lowercase = &v
			case "stopwords":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "pattern":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Pattern = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v

			}
		}
	}
	return query
}

/*
phân tích câu query standar cua analyzer

	standar{...}
*/
func ParseStandardAnalyzer(m reqparams.Querier) *types.StandardAnalyzer {
	query := types.NewStandardAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "max_token_length":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxTokenLength = &v
			case "stopwords":
				field, ok := value.(reqparams.Querier)

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
				query.Stopwords = fields
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query nori cua analyzer

	nori{...}
*/
func ParseNoriAnalyzer(m reqparams.Querier) *types.NoriAnalyzer {
	query := types.NewNoriAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "decompound_mode":
				decom := noridecompoundmode.NoriDecompoundMode{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				decom.Name = v
				query.DecompoundMode = &decom
			case "stoptags":
				field, ok := value.(reqparams.Querier)

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
				query.Stoptags = fields
			case "user_dictionary":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.UserDictionary = &v
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query white_space cua analyzer

	white_space{...}
*/
func ParseWhiteSpaceAnalyzer(m reqparams.Querier) *types.WhitespaceAnalyzer {
	query := types.NewWhitespaceAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query icu cua analyzer

	icu{...}
*/
func ParseIcuAnalyzer(m reqparams.Querier) *types.IcuAnalyzer {
	query := types.NewIcuAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "method":
				method := icunormalizationtype.IcuNormalizationType{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				method.Name = v
				query.Method = method
			case "mode":
				mod := icunormalizationmode.IcuNormalizationMode{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				mod.Name = v
				query.Mode = mod
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			}
		}
	}
	return query
}

/*
phân tích câu query kuromoji analyzer
query có dạng icu_analyzer[...]
*/
func ParseKuromojiAnalyzer(m reqparams.Querier) *types.KuromojiAnalyzer {
	query := types.NewKuromojiAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "mode":
				mod := kuromojitokenizationmode.KuromojiTokenizationMode{}
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				mod.Name = v
				query.Mode = mod
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "user_dictionary":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.UserDictionary = &v
			}
		}
	}
	return query
}

/*
phân tích câu query keyword cua analyzer

	keyword{...}
*/
func ParseKeywordAnalyzer(m reqparams.Querier) *types.KeywordAnalyzer {
	query := types.NewKeywordAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
phân tích câu query simple cua analyzer

	simple{...}
*/
func ParseSimpleAnalyzer(m reqparams.Querier) *types.SimpleAnalyzer {
	query := types.NewSimpleAnalyzer()
	params := m.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch key {
			case "type":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Type = v
			case "version":
				s, ok := value.(string)
				if !ok {
					break
				}
				v := s
				query.Version = &v
			}
		}
	}
	return query
}

/*
Phân tích câu query has_child
câu query có dạng has_child{boost=3,...}
*/
func ParseHasChildQuery(m reqparams.Querier) *types.HasChildQuery {
	query := types.NewHasChildQuery()
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
			case "ignore_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}

				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IgnoreUnmapped = &v
			case "inner_hits":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseInnerHits(field)
				query.InnerHits = i
			case "max_children":
				s, ok := value.(string)
				if !ok {
					break
				}
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MaxChildren = &v
			case "min_children":
				s, ok := value.(string)
				if !ok {
					break
				}
				if !ok {
					break
				}
				v, err := strconv.Atoi(s)
				if err != nil {
					break
				}
				query.MinChildren = &v
			case "query":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Query = i
			case "_name":
				s, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &s
			case "score_mode":
				scm := childscoremode.ChildScoreMode{}
				v, ok := value.(string)
				if !ok {
					break
				}
				scm.Name = v
				query.ScoreMode = &scm
			case "type":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Type = v
			}
		}
	}
	return query
}

/*
Phân tích câu query has_child

	has_child{boost=3,...}
*/
func ParseHasParentQuery(m reqparams.Querier) *types.HasParentQuery {
	query := types.NewHasParentQuery()
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
			case "ignore_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IgnoreUnmapped = &v
			case "inner_hits":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseInnerHits(field)
				query.InnerHits = i

			case "query":
				field, ok := value.(reqparams.Querier)
				if !ok {
					break
				}
				i := ParseQueryToSearch(field)
				query.Query = i
			case "_name":
				s, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &s
			case "parent_type":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.ParentType = v
			case "score":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.Score = &v
			}
		}
	}
	return query
}
func ParseParentIdQuery(m reqparams.Querier) *types.ParentIdQuery {
	query := types.NewParentIdQuery()
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
			case "ignore_unmapped":
				s, ok := value.(string)
				if !ok {
					break
				}
				v, err := strconv.ParseBool(s)
				if err != nil {
					break
				}
				query.IgnoreUnmapped = &v
			case "_name":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.QueryName_ = &v
			case "id":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Id = &v
			case "type":
				v, ok := value.(string)
				if !ok {
					break
				}
				query.Type = &v

			}
		}
	}
	return query
}
