package models

// Member .
type Member struct {
	Model
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
