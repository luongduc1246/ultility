package gormdb

const MaxLimit = 50
const MinLimit = 1

type pageLimit struct {
	Limit  int
	Offset int
}

func NewPageLimit() *pageLimit {
	return &pageLimit{}
}

func (p *pageLimit) Parse(page, limit int) {
	switch true {
	case limit > MaxLimit:
		p.Limit = MaxLimit
	case limit < MinLimit:
		p.Limit = MinLimit
	default:
		p.Limit = limit
	}
	if page > 0 {
		p.Offset = (page - 1) * p.Limit
	}
}
