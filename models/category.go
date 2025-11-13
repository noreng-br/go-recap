package models

type Category struct {
	CategoryID   string `json:"category_id"`
	Name string `json:"name"`
}

type CategoryDTO struct {
	Name string `json:"name"`
}
