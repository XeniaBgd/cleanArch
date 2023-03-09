package entity

type Author struct {
	ID   string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Year int    `json:"year,omitempty"`
}
