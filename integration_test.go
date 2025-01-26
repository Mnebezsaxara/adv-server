package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to in-memory SQLite database")
	}
	if err := db.AutoMigrate(&User{}, &Booking{}); err != nil {
		panic("Failed to migrate database schema")
	}
	return db
}

func TestHandleAuth(t *testing.T) {
	// Настройка базы данных для теста
	db = setupTestDB()
	db.Create(&User{Email: "test@example.com", Password: "password123"})

	// Тест авторизации
	payload := map[string]string{"email": "test@example.com", "password": "password123"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handleAuth(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	var response ResponsePayload
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Status != "success" || response.Message != "Login successful" {
		t.Errorf("Unexpected response: %+v", response)
	}
}
