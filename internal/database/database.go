package database

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DataBase interface {
	Start() error
	Exit() error
	Fetch(query FetchQuery) (*[]FetchResult, error)
	Get(query GetQuery) (*[]GetResult, error)
	Ingest(query IngestQuery) (QueryResult, error)
}

type MetqlDataBase struct {
	masterbase string
	conn       *sqlx.DB
}

func NewDatabase(masterbase string) DataBase {
	db := &MetqlDataBase{}
	db.masterbase = masterbase
	return db
}

func (db *MetqlDataBase) Start() error {
	conn, err := sqlx.Open("sqlite3_with_extensions", db.masterbase)
	if err != nil {
		return fmt.Errorf("could not open connection; %w", err)
	}
	err = conn.Ping()
	if err != nil {
		return fmt.Errorf("could not open connection; %w", err)
	}

	db.conn = conn
	return nil
}

func (db *MetqlDataBase) Exit() error {
	err := db.conn.Close()
	if err != nil {
		return fmt.Errorf("could not close connections %w", err)
	}
	return nil
}

func (db *MetqlDataBase) Exec(query string) (QueryResult, error) {
	if strings.HasPrefix(query, "GET ") {
		getQuery, err := ParseGetQuery(query)
		if err != nil {
			return nil, fmt.Errorf("could not parse query %w", err)
		}
		return db.Get(getQuery)
	}

	if strings.HasPrefix(query, "EAT ") {
		setQuery, err := ParseSetQuery(query)
		if err != nil {
			return nil, fmt.Errorf("could not parse query %w", err)
		}
		return db.Ingest(setQuery)
	}

	return nil, fmt.Errorf("could not parse query %s", query)
}

func (db *MetqlDataBase) Fetch(query FetchQuery) (*[]FetchResult, error) {
	var masterQueryBuilder strings.Builder
	masterQueryBuilder.WriteString("SELECT * FROM keeper")

	if query.GeomtryWkt == "" && len(query.Models) == 0 {
		masterQueryBuilder.WriteRune(';')
		masterQuery := masterQueryBuilder.String()

		results := []FetchResult{}
		err := db.conn.Select(&results, masterQuery)
		if err != nil {
			e := make([]FetchResult, 0)
			return &e, fmt.Errorf("could not execute fetch command; %w", err)
		}

		return &results, nil
	}

	masterQueryBuilder.WriteString(" WHERE ")
	addAnd := false

	params := make(map[string]string)
	if query.GeomtryWkt != "" {
		masterQueryBuilder.WriteString("ST_Intersects(geom, :geo)")
		params["geo"] = fmt.Sprintf("ST_GeomFromText(%s)", query.GeomtryWkt)
		addAnd = true
	}

	useIn := false
	if len(query.Models) > 0 {
		if addAnd {
			masterQueryBuilder.WriteString(" AND ")
		}

		masterQueryBuilder.WriteString("model_name IN (:models)")
		params["models"] = fmt.Sprintf("ST_GeomFromText(%s)", query.GeomtryWkt)

		useIn = true
	}

	masterQuery := masterQueryBuilder.String()
	namedQuery, args, err := sqlx.Named(masterQuery, params)
	if err != nil {
		e := make([]FetchResult, 0)
		return &e, fmt.Errorf("could not create named query; %w", err)
	}

	if useIn {
		namedQuery, args, err = sqlx.In(namedQuery)
		if err != nil {
			e := make([]FetchResult, 0)
			return &e, fmt.Errorf("could not create IN query; %w", err)
		}
	}

	results := []FetchResult{}
	namedQuery = db.conn.Rebind(namedQuery)
	err = db.conn.Select(&results, namedQuery, args...)
	if err != nil {
		e := make([]FetchResult, 0)
		return &e, fmt.Errorf("could not execute query; %w", err)
	}

	return &results, nil
}

func (db *MetqlDataBase) Get(query GetQuery) (*[]GetResult, error) {
	return nil, nil
}

func (db *MetqlDataBase) Ingest(query IngestQuery) (QueryResult, error) {
	return nil, nil
}
