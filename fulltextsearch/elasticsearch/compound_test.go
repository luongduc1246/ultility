package elasticsearch

import (
	"fmt"
	"testing"

	"github.com/luongduc1246/ultility/reqparams"
)

func TestParseBoolQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "bool[boost=3,filter[query_search[match[field[query=test]],intervals[fields[all_of[ordered=true]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Bool.Filter)
}
func TestParseBoostingQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "boosting[boost=3,negative[match[field[query=3],mess[query=test]],intervals[fields[all_of[ordered=true]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.Boosting.Negative.Intervals)
}
func TestParseConstantScoreQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "constant_score[boost=3,filter[match[field[query=test]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.ConstantScore.Filter.Match)
}
func TestParseDismaxQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "dis_max[boost=3,queries[query_search[match[field[query=test]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.DisMax)
}
func TestParseFunctionScoreQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[boost=3,query_search[match[fields[query=test]]],score_mode=score,boost_mode=boost]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore)
}
func TestParseDecayFunctionUntypedQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[exp[untyped_decay_function[decay_parameters[fields[decay=cay,offset=off,origin=origin,scale=2]]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].Exp)
}
func TestParseDecayFunctionDateQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[exp[date_decay_function[decay_parameters[fields[decay=cay,offset=off,origin=origin,scale=2]]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].Exp)
}
func TestParseDecayFunctionGeoQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[exp[geo_decay_function[decay_parameters[fields[decay=cay,offset=off,origin=origin,scale=2]]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].Exp)
}
func TestParseDecayFunctionNumericQuery(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[exp[numeric_decay_function[decay_parameters[fields[decay=cay,offset=off,origin=origin,scale=2]]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].Exp)
}
func TestParseRandomScore(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[random_score[field=3,seed=4]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].RandomScore)
}
func TestParseScriptScore(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[script_score[script[id=id123,lang=vn,options[op1=duc,op2=test],params[id=3]]]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].ScriptScore)
}
func TestParseFieldValueFactor(t *testing.T) {
	q := reqparams.NewQuery()
	s := "function_score[functions[function_score[field_value_factor[factor=4,field=id,missing=64,modifier=fd]]]]"
	q.Parse(s)
	query := ParseQueryToSearch(q)
	fmt.Println(query.FunctionScore.Functions[0].FieldValueFactor)
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := reqparams.NewQuery()
		s := "bool[boost=3]"
		q.Parse(s)
		ParseQueryToSearch(q)
	}
}
