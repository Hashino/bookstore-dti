package repo

import (
	"bookstore/db"
	"bookstore/model"
	"os"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) {
	_ = os.Remove("test_books.db")
	dbFile := "test_books.db"
	os.Setenv("BOOKSTORE_DB", dbFile)
	err := db.InitTestDB(dbFile)
	if err != nil {
		t.Fatalf("failed to init test DB: %v", err)
	}
}

func teardownTestDB() {
	db.DB.Close()
	_ = os.Remove("test_books.db")
}

func TestAddAndGetBook(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB()

	book := model.Book{
		Title:       "Clean Code",
		Author:      "Robert C. Martin",
		ReleaseDate: time.Date(2008, 8, 1, 0, 0, 0, 0, time.UTC),
		Score:       9.0,
		Summary:     "A Handbook of Agile Software Craftsmanship",
	}

	err := AddBook(book)
	if err != nil {
		t.Fatalf("failed to add book: %v", err)
	}

	books, err := ListBooks()
	if err != nil {
		t.Fatalf("failed to list books: %v", err)
	}

	if len(books) != 1 {
		t.Fatalf("expected 1 book, got %d", len(books))
	}

	saved := books[0]
	if saved.Title != book.Title || saved.Author != book.Author {
		t.Errorf("book data mismatch: got %+v", saved)
	}
}

func TestUpdateBook(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB()

	book := model.Book{
		Title:       "Old Title",
		Author:      "Author",
		ReleaseDate: time.Now(),
		Score:       5.0,
		Summary:     "Old summary",
	}
	_ = AddBook(book)
	books, _ := ListBooks()
	book = books[0]

	book.Title = "New Title"
	book.Score = 7.5

	err := UpdateBook(book)
	if err != nil {
		t.Fatalf("failed to update book: %v", err)
	}

	updated, _ := GetBookByID(book.ID)
	if updated.Title != "New Title" || updated.Score != 7.5 {
		t.Errorf("book update failed, got %+v", updated)
	}
}

func TestDeleteBook(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB()

	book := model.Book{
		Title:       "Delete Me",
		Author:      "Author",
		ReleaseDate: time.Now(),
		Score:       3.0,
		Summary:     "",
	}
	_ = AddBook(book)
	books, _ := ListBooks()
	book = books[0]

	ok, err := DeleteBook(book.ID)
	if err != nil || !ok {
		t.Fatalf("failed to delete book: %v", err)
	}

	_, err = GetBookByID(book.ID)
	if err == nil {
		t.Error("expected error for non-existent book after deletion")
	}
}
