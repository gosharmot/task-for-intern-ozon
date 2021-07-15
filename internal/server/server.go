package server

import (
	"encoding/json"
	"log"
	"net/http"
	"task-for-intern/internal/store"
)

type Server struct {
	adress string
	store  *store.Store
}

func NewServer(pattern string) *Server {
	return &Server{
		adress: pattern,
	}
}

func (s *Server) Listen() error {
	log.Println("Listening server...")

	http.HandleFunc("/short", s.getShortUrl())
	http.HandleFunc("/long", s.getLongUrl())

	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.adress, nil)
}

func (s *Server) getShortUrl() http.HandlerFunc {
	type request struct {
		Url string `json:"url"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			req := &request{}

			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				log.Println(err)
				return
			}

			shortUrl, err := s.store.URL().GetShortUrl(req.Url)
			if err != nil {
				log.Println(err)
				return
			}

			req.Url = shortUrl

			json.NewEncoder(rw).Encode(req)
		}
	}
}

func (s *Server) getLongUrl() http.HandlerFunc {
	type request struct {
		Url string `json:"url"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			req := &request{}

			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				return
			}

			longUrl, err := s.store.URL().GetLongUrl(req.Url)

			if err != nil {
				log.Println(err)
			}

			req.Url = longUrl

			json.NewEncoder(rw).Encode(req)
		}
	}
}

func (s *Server) configureStore() error {
	store := store.NewStore()

	if err := store.Open(); err != nil {
		return err
	}

	s.store = store
	return nil
}
