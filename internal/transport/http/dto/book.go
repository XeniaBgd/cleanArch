package dto

type CreateBookDTO struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorID string `json:"author"`
}

type UpdateBookDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorID string `json:"author"`
	Busy     bool   `json:"busy"`
	Owner    string `json:"owner"`
}
