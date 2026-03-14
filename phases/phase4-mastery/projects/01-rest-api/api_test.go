package restapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestListItems(t *testing.T) {
	store := NewStore()
	store.Create("a", "1")
	store.Create("b", "2")
	api := NewAPI(store)

	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	w := httptest.NewRecorder()
	api.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("GET /items status = %d, want 200", w.Code)
	}
	var items []Item
	json.NewDecoder(w.Body).Decode(&items)
	if len(items) != 2 {
		t.Errorf("GET /items returned %d items, want 2", len(items))
	}
}

func TestCreateItem(t *testing.T) {
	api := NewAPI(NewStore())
	body := bytes.NewBufferString(`{"name":"widget","value":"42"}`)
	req := httptest.NewRequest(http.MethodPost, "/items", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	api.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("POST /items status = %d, want 201", w.Code)
	}
	var created Item
	json.NewDecoder(w.Body).Decode(&created)
	if created.ID == 0 {
		t.Error("created item should have non-zero ID")
	}
}

func TestGetItem(t *testing.T) {
	store := NewStore()
	item := store.Create("test", "val")
	api := NewAPI(store)

	req := httptest.NewRequest(http.MethodGet, "/items/"+strconv.Itoa(item.ID), nil)
	w := httptest.NewRecorder()
	api.Handler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("GET /items/%d status = %d, want 200", item.ID, w.Code)
	}
}

func TestGetItemNotFound(t *testing.T) {
	api := NewAPI(NewStore())
	req := httptest.NewRequest(http.MethodGet, "/items/999", nil)
	w := httptest.NewRecorder()
	api.Handler().ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("GET /items/999 status = %d, want 404", w.Code)
	}
}

func TestDeleteItem(t *testing.T) {
	store := NewStore()
	item := store.Create("x", "y")
	api := NewAPI(store)

	req := httptest.NewRequest(http.MethodDelete, "/items/"+strconv.Itoa(item.ID), nil)
	w := httptest.NewRecorder()
	api.Handler().ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("DELETE /items/%d status = %d, want 204", item.ID, w.Code)
	}
}
