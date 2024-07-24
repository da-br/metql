package database

type QueryResult interface{}

type FetchResult struct {
	Id          int
	TimeLower   float64 `db:"time_lower"`
	TimeUpper   float64 `db:"time_upper"`
	Deleted     float64
	Created     float64
	ModelName   string `db:"model_name"`
	BoundingBox string `db:"bounding_box"`
	Geom        string
	Path        string
}

type GetResult struct{}
