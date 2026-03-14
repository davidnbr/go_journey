package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Item is a REST resource.
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Store holds items in memory.
type Store struct {
	mu     sync.RWMutex
	items  map[int]Item
	nextID int
}

func NewStore() *Store {
	return &Store{items: make(map[int]Item), nextID: 1}
}

func (s *Store) Create(name, value string) Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	item := Item{ID: s.nextID, Name: name, Value: value}
	s.items[s.nextID] = item
	s.nextID++
	return item
}

func (s *Store) Get(id int) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.items[id]
	return item, ok
}

func (s *Store) List() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return false
	}
	delete(s.items, id)
	return true
}

// API wires routes to handlers.
type API struct {
	store *Store
}

func NewAPI(store *Store) *API {
	return &API{store: store}
}

// Handler returns the http.Handler with all routes registered.
// Routes:
//
//	GET    /items        -> list all
//	POST   /items        -> create
//	GET    /items/{id}   -> get one
//	DELETE /items/{id}   -> delete
//
// TODO: use http.NewServeMux, register routes, implement handlers
func (a *API) Handler() http.Handler {
	mux := http.NewServeMux()
	// TODO: register routes
	return mux
}

// Middleware: Logger, Recovery, CORS header setter
// TODO: implement as in exercise 05

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

var _, _, _ = strconv.Atoi, strings.TrimPrefix, fmt.Sprintf
