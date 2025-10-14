package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/pz3-http/internal/storage"
)

func TestCreateTask_Success(t *testing.T) {
	store := storage.NewMemoryStore()
	handlers := NewHandlers(store)

	body := `{"title": "Купить молоко"}`

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handlers.CreateTask(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Ожидался статус 201, получен %d", rr.Code)
	}

	var task storage.Task
	json.Unmarshal(rr.Body.Bytes(), &task)

	if task.Title != "Купить молоко" {
		t.Errorf("Ожидался title 'Купить молоко', получен '%s'", task.Title)
	}

}
