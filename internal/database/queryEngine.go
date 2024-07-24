package database

type FetchQuery struct {
	Models                       []string
	GeomtryWkt                   string
	TimeRangeStart, TimeRangeEnd float64
}

type GetQuery struct {
	SpatialQuery                 SpatialQueryType
	GeomtryWkt                   string
	TimeRangeStart, TimeRangeEnd float64
	Columns                      []string
}

type IngestQuery struct {
	FileName, TableName string
	Optimize            OptimizeType
}

func ParseGetQuery(query string) (GetQuery, error) {
	return GetQuery{}, nil
}

func ParseSetQuery(query string) (IngestQuery, error) {
	return IngestQuery{}, nil
}
