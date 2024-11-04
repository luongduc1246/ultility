package fulltextsearch

const (
	NESTED         QueryKey = "nested"
	IGNOREUNMAPPED QueryKey = "ignore_unmapped"
	INNERHITS      QueryKey = "inner_hits"

	HASCHILD    QueryKey = "has_child"
	MAXCHILDREN QueryKey = "max_children"
	MINCHILDREN QueryKey = "min_children"

	HASPARENT  QueryKey = "has_parent"
	SCORE      QueryKey = "score"
	PARENTTYPE QueryKey = "parent_type"
	PARENTID   QueryKey = "parent_id"
	ID         QueryKey = "id"
)

type Nested struct {
	Querier
}
type InnerHits Nested

type IgnoreUnmapped bool

type HasChild struct {
	Querier
}

type MaxChildren int
type MinChildren int

type HasParent struct {
	Querier
}

type Score bool
type ParentType string

type ParentId struct {
	Querier
}

type Id string
