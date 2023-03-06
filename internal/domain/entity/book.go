package entity

import (
	"fmt"
)

type BookView struct {
	ID         string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	Year       int    `json:"year,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Busy       bool   `json:"busy,omitempty"`
}

type FullBook struct {
	Book
	Author Author
}

type Book struct {
	ID       string `json:"uuid,omitempty"`
	Name     string `json:"name,omitempty"`
	Year     int    `json:"year,omitempty"`
	AuthorID string `json:"author,omitempty"`
	Busy     bool   `json:"busy,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}

	b.Owner = owner
	b.Busy = true
	return nil
}
