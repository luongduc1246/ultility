package reqparams

type Params struct {
	Fields *string `json:"fields" form:"fields" query:"fields"`
	Sort   *string `json:"sort" form:"sort" query:"sort"`       /* sort cho sql */
	Filter *string `json:"filter" form:"filter" query:"filter"` /* filter cho sql */
	Query  *string `json:"query" form:"query" query:"query"`    /* fulltextsearch "elasticsearch" */
	Page   int     `json:"page" form:"page" query:"page"`
	Limit  int     `json:"limit" form:"limit" query:"limit"`
}

type Search struct {
	Fields *Slice
	Sort   *Slice
	Filter *Query
	Query  *Query
	Page   int
	Limit  int
}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) Parse(p Params) error {
	if p.Fields != nil {
		slice := NewSlice()
		err := slice.Parse(*p.Fields)
		if err != nil {
			return err
		}
		if len(slice.Params) > 0 {
			s.Fields = slice
		}
	}
	if p.Filter != nil {
		filter := NewQuery()
		err := filter.Parse(*p.Filter)
		if err != nil {
			return err
		}
		if len(filter.Params) > 0 {
			s.Filter = filter
		}
	}
	if p.Sort != nil {
		sort := NewSlice()
		err := sort.Parse(*p.Sort)
		if err != nil {
			return err
		}
		if len(sort.Params) > 0 {
			s.Sort = sort
		}
	}
	if p.Query != nil {
		query := NewQuery()
		err := query.Parse(*p.Query)
		if err != nil {
			return err
		}
		if len(query.Params) > 0 {
			s.Query = query
		}
	}
	s.Page = p.Page
	s.Limit = p.Limit
	return nil
}
