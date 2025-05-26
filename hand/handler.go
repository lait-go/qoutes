package hand

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}

type Store struct {
	sync.Mutex
	list   []Quote
	nextID int
}

func NewStore() *Store {
	return &Store{
		list:   make([]Quote, 0),
		nextID: 1,
	}
}

func (s *Store) AddQuote(w http.ResponseWriter, r *http.Request) {
	var q Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.Lock()
	q.ID = s.nextID
	s.nextID++
	s.list = append(s.list, q)
	s.Unlock()

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(q)
}

func (s *Store) GetAll(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	defer s.Unlock()
	json.NewEncoder(w).Encode(s.list)
}

func (s *Store) GetRandom(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	defer s.Unlock()
	if len(s.list) == 0 {
		http.Error(w, "No quotes", 404)
		return
	}
	rnd := s.list[rand.Intn(len(s.list))]
	json.NewEncoder(w).Encode(rnd)
}

func (s *Store) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	a := strings.ToLower(r.URL.Query().Get("author"))
	s.Lock()
	defer s.Unlock()
	var res []Quote
	for _, q := range s.list {
		if strings.ToLower(q.Author) == a {
			res = append(res, q)
		}
	}
	json.NewEncoder(w).Encode(res)
}

func (s *Store) QuotesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("QuotesHandler %s %s\n", r.Method, r.URL.String())

	switch r.Method {
	case "POST":
		s.AddQuote(w, r)
	case "GET":
		author := r.URL.Query().Get("author")
		if author != "" {
			s.GetByAuthor(w, r)
		} else {
			s.GetAll(w, r)
		}
	default:
		http.Error(w, "Not allowed", 405)
	}
}

func (s *Store) RandomHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("RandomHandler %s %s\n", r.Method, r.URL.String())
	if r.Method != "GET" {
		http.Error(w, "Not allowed", 405)
		return
	}
	s.GetRandom(w, r)
}

func (s *Store) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteHandler %s %s\n", r.Method, r.URL.String())
	if r.Method != "DELETE" {
		http.Error(w, "Not allowed", 405)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad ID", 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	for i, q := range s.list {
		if q.ID == id {
			s.list = append(s.list[:i], s.list[i+1:]...)
			w.WriteHeader(204)
			return
		}
	}

	http.Error(w, "Not found", 404)
}
