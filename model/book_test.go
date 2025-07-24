package model

import (
	"testing"
	"time"
)

func TestBookValidation(t *testing.T) {
	validBook := Book{
		Title:       "Example",
		Author:      "Author",
		ReleaseDate: time.Now(),
		Score:       8.5,
	}

	if err := validBook.Validate(); err != nil {
		t.Errorf("valid book failed validation: %v", err)
	}

	invalidBooks := []Book{
		{Title: "", Author: "Author", ReleaseDate: time.Now(), Score: 8},
		{Title: "Title", Author: "", ReleaseDate: time.Now(), Score: 8},
		{Title: "Title", Author: "Author", ReleaseDate: time.Time{}, Score: 8},
		{Title: "Title", Author: "Author", ReleaseDate: time.Now(), Score: -1},
		{Title: "Title", Author: "Author", ReleaseDate: time.Now(), Score: 11},
	}

	for _, b := range invalidBooks {
		if err := b.Validate(); err == nil {
			t.Errorf("expected validation error for book: %+v", b)
		}
	}
}
