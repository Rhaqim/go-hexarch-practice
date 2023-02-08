package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	// ping
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	return &Adapter{
		db: db,
	}, nil
}

func (adp Adapter) CloseDB() {
	adp.db.Close()
}

func (adp Adapter) AddtoHistory(result int32, operation string) error {

	querystring, args, err := sq.Insert("history").Columns("result", "operation", "created_at").Values(result, operation, time.Now()).ToSql()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = adp.db.Exec(querystring, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
