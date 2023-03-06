package dto

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"year"`
}

type UpdateAuthorDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"year"`
}
