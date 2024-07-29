package reqparams

type Params struct {
	Field  []string `json:"fields" form:"fields" query:"fields"`
	Sort   []string `json:"sort" form:"sort" query:"sort"`
	Filter []string `json:"filter" form:"filter" query:"filter"`
	Page   int      `json:"page" form:"page" query:"page"`
	Limit  int      `json:"limit" form:"limit" query:"limit"`
}

type Search struct {
	Field  *Field
	Sort   *Sort
	Filter *Filter
	Page   int
	Limit  int
}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) Parse(p Params) error {
	var err error
	field := NewField()
	for _, value := range p.Field {
		err = field.Parse(value)
		if err != nil {
			return err
		}
	}
	if len(field.Columns) > 0 || len(field.Relatives) > 0 {
		s.Field = field
	}
	filter := NewFilter()
	for _, value := range p.Filter {
		err = filter.Parse(value)
		if err != nil {
			return err
		}
	}
	if len(filter.Exps) > 0 || len(filter.Relatives) > 0 {
		s.Filter = filter
	}
	sort := NewSort()
	for _, value := range p.Sort {
		err = sort.Parse(value)
		if err != nil {
			return err
		}
	}
	if len(sort.Orders) > 0 || len(sort.Relatives) > 0 {
		s.Sort = sort
	}
	s.Page = p.Page
	s.Limit = p.Limit
	return nil
}
