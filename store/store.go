package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func NewStore(dbinfo string) *Store {

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	store := &Store{
		DB: db,
	}

	return store
}
