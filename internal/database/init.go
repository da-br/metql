package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

func init() {
	sql.Register("sqlite3_with_extensions", &sqlite3.SQLiteDriver{
		Extensions: []string{
			"mod_spatialite",
		},
		// ConnectHook: func(conn *sqlite3.SQLiteConn) error {
		// 	if err := conn.LoadExtension("mod_spatialite", ""); err != nil {
		// 		slog.Error("could not load ext", slog.String("error", err.Error()))
		// 		return err
		// 	}
		// 	return nil
		// },
	})
}

func Init(location, name string, force bool) error {
	dbDir := filepath.Join(location, name)

	_, err := os.Stat(dbDir)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Info("creating db directory", slog.String("directory", dbDir))
			if err = os.MkdirAll(dbDir, os.ModePerm); err != nil {
				return fmt.Errorf("could not create directory %s; %w", dbDir, err)
			}
		} else {
			return fmt.Errorf("unknown error checking directory; %w", err)
		}
	}

	dbFile := path.Join(dbDir, "masterbase.db")
	fileStat, _ := os.Stat(dbFile)
	if fileStat != nil {
		if !force {
			slog.Warn("database already exists; to overwrite this pass the force flag, this will overwrite all contents of the database")
			return nil
		}

		slog.Warn("database already exists and force was passed. removing all entires after 5sec")
		time.Sleep(time.Second * 5)
		os.RemoveAll(dbDir)
		os.MkdirAll(dbDir, os.ModePerm)
	}

	connectionString := fmt.Sprintf("file:%s", dbFile)
	slog.Info("creating database", slog.String("connectionString", connectionString))

	db, err := sqlx.Open("sqlite3_with_extensions", connectionString)
	if err != nil {
		slog.Error("could not open connection to db", slog.String("error", err.Error()))
		return fmt.Errorf("could not open connection; %w", err)
	}

	defer func() {
		tmpErr := db.Close() // can  be null
		err = errors.Join(err, tmpErr)
	}()

	_, err = db.Exec(`SELECT InitSpatialMetaData()`)
	if err != nil {
		slog.Error("error loading spatiallite metadata", slog.String("error", err.Error()))
		return fmt.Errorf("could not load spatial lite metadata; %w", err)
	}

	return nil
}
