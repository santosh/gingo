package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/santosh/gingo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

	DB = database
}

// Books slice to seed book data.

func Seed() {
	var books = []models.Book{
		{ISBN: "9781612680194", Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki"},
		{ISBN: "9781781257654", Title: "The Daily Stotic", Author: "Ryan Holiday"},
		{ISBN: "9780593419052", Title: "A Mind for Numbers", Author: "Barbara Oklay"},
	}

	for _, book := range books {
		fmt.Println(book)

		url := "http://localhost:8090/api/v1/books"
		bookBytes, _ := json.Marshal(book)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bookBytes))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
}
