package entity

type Author struct {
	ID   string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
