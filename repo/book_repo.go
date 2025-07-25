package repo

import (
	"bookstore/db"
	"bookstore/model"
	"time"
)

// AddBook adds a new book to the database.
func AddBook(book model.Book) error {
	if err := book.Validate(); err != nil {
		return err
	}
	stmt, err := db.DB.Prepare("INSERT INTO books(title, author, release_date, score, summary) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.Title, book.Author, book.ReleaseDate.Format("2006-01-02"), book.Score, book.Summary)
	return err
}

// ListBooks retrieves all books from the database.
func ListBooks() ([]model.Book, error) {
	rows, err := db.DB.Query("SELECT id, title, author, release_date, score, summary FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		var dateStr string
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &dateStr, &b.Score, &b.Summary); err != nil {
			return nil, err
		}
		b.ReleaseDate, _ = time.Parse("2006-01-02", dateStr)
		books = append(books, b)
	}
	return books, nil
}

// GetBookByID retrieves a book by its ID from the database.
// If the book is not found, it returns an error.
func GetBookByID(id int) (model.Book, error) {
	var b model.Book
	var dateStr string
	row := db.DB.QueryRow("SELECT id, title, author, release_date, score, summary FROM books WHERE id = ?", id)
	err := row.Scan(&b.ID, &b.Title, &b.Author, &dateStr, &b.Score, &b.Summary)
	if err != nil {
		return b, err
	}
	b.ReleaseDate, _ = time.Parse("2006-01-02", dateStr)
	return b, nil
}

// UpdateBook updates an existing book in the database.
// If the book does not exist, it returns an error.
func UpdateBook(book model.Book) error {
	stmt, err := db.DB.Prepare("UPDATE books SET title = ?, author = ?, release_date = ?, score = ?, summary = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.Title, book.Author, book.ReleaseDate.Format("2006-01-02"), book.Score, book.Summary, book.ID)
	return err
}

// DeleteBook deletes a book from the database by its ID.
// It returns true if the book was deleted, or false if it was not found.
func DeleteBook(id int) (bool, error) {
	stmt, err := db.DB.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}
	affected, _ := res.RowsAffected()
	return affected > 0, nil
}
