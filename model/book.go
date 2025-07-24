package model

import (
	"errors"
	"strings"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Author      string
	ReleaseDate time.Time
	Score       float64
	Summary     string
}

func (b *Book) Validate() error {
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("título não pode estar vazio")
	}
	if strings.TrimSpace(b.Author) == "" {
		return errors.New("autor não pode estar vazio")
	}
	if b.ReleaseDate.IsZero() || b.ReleaseDate.Year() < 1000 {
		return errors.New("data de lançamento inválida")
	}
	if b.Score < 0.0 || b.Score > 10.0 {
		return errors.New("nota deve estar entre 0.0 e 10.0")
	}
	return nil
}
