package database

import (
	"fmt"
	"strings"
)

type QueryType int

const (
	QUERY_ERROR QueryType = iota
	QUERY_GET
	QUERY_SET
)

func QueryTypeFromString(s string) (QueryType, error) {
	if strings.EqualFold(s, "get") {
		return QUERY_GET, nil
	}
	if strings.EqualFold(s, "set") {
		return QUERY_SET, nil
	}
	return QUERY_ERROR, fmt.Errorf("could not parse %s to %T", s, QUERY_ERROR)
}

type AggregrateType int

const (
	AGGREGRATE_NONE AggregrateType = iota
	AGGREGRATE_SUM
	AGGREGRATE_MIN
	AGGREGRATE_MAX
	AGGREGRATE_AVG
)

func AggregrateTypeFromString(s string) (AggregrateType, error) {
	if strings.EqualFold(s, "none") {
		return AGGREGRATE_NONE, nil
	}

	if strings.EqualFold(s, "sum") {
		return AGGREGRATE_SUM, nil
	}

	if strings.EqualFold(s, "min") {
		return AGGREGRATE_MIN, nil
	}

	if strings.EqualFold(s, "max") {
		return AGGREGRATE_MAX, nil
	}

	if strings.EqualFold(s, "avg") {
		return AGGREGRATE_AVG, nil
	}

	return AGGREGRATE_NONE, fmt.Errorf("could not parse %s to %T", s, AGGREGRATE_NONE)
}

type OptimizeType int

const (
	OPTIMIZED_NONE OptimizeType = iota
	OPTIMIZED_TIMESERIES
	OPTIMIZED_AREA
)

type SpatialQueryType int

const (
	SPATIALQUERYTYPE_NONE SpatialQueryType = iota
	SPATIALQUERYTYPE_IN
	SPATIALQUERYTYPE_NEAREST
)
