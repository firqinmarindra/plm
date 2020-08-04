package settingdb

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

type DatabaseConfig struct{}

const (
	host     = "localhost"
	dbname   = "postgres"
	user     = "firqin"
	password = "070496"
	schema   = "public"
)

func (DatabaseConfig DatabaseConfig) GetDatabaseConfig() *sql.DB {
	fmt.Print()
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", user, password, host, dbname, schema)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}


	return db
}
