package fulltextsearch

const (
	KNN                QueryKey = "knn"
	NUMCANDIDATES      QueryKey = "num_candidates"
	QUERYVECTOR        QueryKey = "query_vector"
	QUERYVECTORBUILDER QueryKey = "query_vector_builder"
	SIMILARITY         QueryKey = "similarity"
	SPARSEVECTOR       QueryKey = "sparse_vector"
	INFERENCEID        QueryKey = "inference_id"
	PRUNE              QueryKey = "prune"
	PRUNINGCONFIG      QueryKey = "pruning_config"
	SEMANTIC           QueryKey = "semantic"
)

type Knn struct {
	Querier
}
type QueryVectorBuilder Knn

type QueryVector []float32

type NumCandidates int

type Similarity float32

type SparseVector struct {
	Querier
}
type PruningConfig SparseVector
type InferenceId string
type Prune bool

type Semantic struct {
	Querier
}
