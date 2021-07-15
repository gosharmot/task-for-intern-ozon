package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	databaseUrl   string
	db            *sql.DB
	urlRepository *URLRepository
}

func NewStore() *Store {
	return &Store{
		databaseUrl: "host=localhost dbname=task_for_intern sslmode=disable", //task_for_intern
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.databaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) URL() *URLRepository {
	if s.urlRepository != nil {
		return s.urlRepository
	}

	s.urlRepository = &URLRepository{
		store: s,
	}

	return s.urlRepository
}
