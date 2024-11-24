package reqparams

/*
	- các giá trị value nên url encode
	- query có dạng eq[adfa]=value1,not[eq[name]=value2,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]
	- Phân tích để lấy column cho câu query where và join trong sql
	- Dùng ngăn xếp để phân tích câu query
	- có lấy field của các quan hệ của các bảng
		- roles[...] là quan hệ trong bảng

	- làm việc với json (extract,haskey làm việc giống như bình thường có dạng extract[name]=blc,haskey)
	likes[field]=(a;b;c:)

*/

import (
	"fmt"

	"github.com/luongduc1246/ultility/arrays"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FilterKey string

const (
	EQ   FilterKey = "eq"
	NEQ  FilterKey = "neq"
	LT   FilterKey = "lt"
	LTE  FilterKey = "lte"
	GT   FilterKey = "gt"
	GTE  FilterKey = "gte"
	IN   FilterKey = "in"
	LIKE FilterKey = "like"
	AND  FilterKey = "and"
	NOT  FilterKey = "not"
	OR   FilterKey = "or"
	/* làm việc với JSON */
	EXTRACT FilterKey = "extract"
	HASKEY  FilterKey = "haskey"
	EQUALS  FilterKey = "equals"
	LIKES   FilterKey = "likes"
	/* làm việc với JSONARRAY */
	CONTAINS FilterKey = "contains"
)

type Exp interface{}

type Eq struct {
	Column string
	Value  interface{}
}

type Neq Eq

type Lt Eq

type Lte Eq

type Gt Eq

type Gte Eq

type Like Eq

type In struct {
	Column string
	Values []interface{}
}

/* làm việc với Json */
type Extract struct {
	Column string
	Value  string
}
type Contains Eq
type Haskey struct {
	Column string
	Values []string
}

type Likes struct {
	Column string
	Keys   []string
	Value  interface{}
}
type Equals Likes

type Filter struct {
	Exps      []Exp // mảng các compare
	Relatives map[string]IFilter
}

func (f *Filter) addExp(exp Exp) {
	f.Exps = append(f.Exps, exp)
}
func (f *Filter) GetExps() []Exp {
	return f.Exps
}
func (f *Filter) GetRelatives() map[string]IFilter {
	return f.Relatives
}

func (f Filter) addRelative(key string, value IFilter) {
	f.Relatives[key] = value
}

func NewFilter() *Filter {
	return &Filter{
		Exps:      make([]Exp, 0),
		Relatives: make(map[string]IFilter),
	}
}

type And struct{ *Filter }

type Not struct{ *Filter }

type Or struct{ *Filter }

type IFilter interface {
	GetExps() []Exp
	GetRelatives() map[string]IFilter
	addExp(Exp)
	addRelative(string, IFilter)
}

/*
#Phân tích từ Querier sang Filter
*/
func (f *Filter) ParseFromQuerier(er Querier) error {
	params := er.GetParams()
	switch t := params.(type) {
	case map[string]interface{}:
		for key, value := range t {
			switch FilterKey(key) {
			case EQ:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Eq{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case NEQ:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Neq{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case LT:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Lt{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case LTE:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Lte{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case GT:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Gt{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case GTE:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Gte{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case LIKE:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							exp := Like{}
							exp.Column = k
							exp.Value = v
							f.Exps = append(f.Exps, exp)
						}
					}
				}
			case IN:
				/* in{column[value,value]} */
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							sli, ok := v.(*Slice)
							if ok {
								sliI, ok := sli.GetParams().([]interface{})
								if ok {
									exp := In{}
									exp.Column = k
									exp.Values = sliI
									f.Exps = append(f.Exps, exp)
								}

							}
						}
					}
				}
			case EXTRACT:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							s, ok := v.(string)
							if ok {
								exp := Extract{}
								exp.Value = s
								exp.Column = k

								f.Exps = append(f.Exps, exp)
							}
						}
					}
				}
			case CONTAINS:
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							s, ok := v.(string)
							if ok {
								exp := Contains{}
								exp.Value = s
								exp.Column = k

								f.Exps = append(f.Exps, exp)
							}
						}
					}
				}
			case HASKEY:
				/* haskey{column[value,value]} */
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							sli, ok := v.(*Slice)
							if ok {
								sliI, ok := sli.GetParams().([]interface{})
								sliString := arrays.ConvertSliceInterfaceToSliceType[string](sliI)
								if ok {
									exp := Haskey{}
									exp.Column = k
									exp.Values = sliString
									f.Exps = append(f.Exps, exp)
								}

							}
						}
					}
				}
			case LIKES:
				/* likes{column{value[path,path]}} */
				exp, ok := value.(*Query)
				if ok {
					pars, ok := exp.GetParams().(map[string]interface{})
					if ok {
						for k, v := range pars {
							pare, ok := v.(*Query)
							if ok {
								parv, ok := pare.GetParams().(map[string]interface{})
								if ok {
									for km, vm := range parv {
										sli, ok := vm.(*Slice)
										if ok {
											sliI, ok := sli.GetParams().([]interface{})
											sliString := arrays.ConvertSliceInterfaceToSliceType[string](sliI)
											if ok {
												exp := Likes{}
												exp.Column = k
												exp.Keys = sliString
												exp.Value = km
												f.Exps = append(f.Exps, exp)
											}

										}
									}
								}
							}
						}
					}
				}
			default:
				/* relative{...expr} */
				fmt.Println(value)
				expr, ok := value.(*Query)
				if ok {
					filter := NewFilter()
					err := filter.ParseFromQuerier(expr)
					if err == nil {
						keyRelative := cases.Title(language.Und, cases.NoLower).String(key)
						f.addRelative(keyRelative, filter)
					}
				}
			}
		}
	}
	return nil
}
