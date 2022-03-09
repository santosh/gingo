package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/santosh/gingo/models"
	"github.com/santosh/gingo/routes"
	"github.com/stretchr/testify/assert"
)


func TestBooksRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "9781612680194")
	assert.Contains(t, w.Body.String(), "9781781257654")
	assert.Contains(t, w.Body.String(), "9780593419052")
}

func TestBooksbyISBNRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/9781612680194", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Rich Dad Poor Dad")
}


func TestPostBookRoute(t *testing.T) {
	router := routes.SetupRouter()

	book := models.Book{
		ISBN: "1234567891012",
		Author: "Santosh Kumar",
		Title: "Hello World",
	}

	body, _ := json.Marshal(book)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "Hello World")
}
