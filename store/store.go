package store

import (
	"database/sql"
	"log"

	"github.com/abstract300/light/messages"
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

func (st *Store) StoreChannelMessage(msg messages.Message) error {
	_, err := st.DB.Exec("INSERT INTO messageauthor(id, username) values($1, $2)", msg.MessageAuthor.ID, msg.MessageAuthor.Username)

	if err != nil {
		log.Println(err)
	}

	return err
}
