package fulltextsearch

const (
	GEOBOUNDINGBOX   QueryKey = "geo_bounding_box"
	VALIDATIONMETHOD QueryKey = "validation_method"
	GEODISTANCE      QueryKey = "geo_distance"
	DISTANCE         QueryKey = "distance"
	DISTANCETYPE     QueryKey = "distance_type"
	GEOPOLYGON       QueryKey = "geo_polygon"
	GEOSHAPE         QueryKey = "geo_shape"
	SHAPE            QueryKey = "shape"
)

type GeoBoundingBox struct {
	Querier
}

type ValidationMethod string

type GeoDistance struct {
	Querier
}
type Distance string
type DistanceType string

type GeoPolygon struct {
	Querier
}
type GeoShape struct {
	Querier
}
type Shape struct {
	Querier
}
