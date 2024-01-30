package entity

type Product struct {
	ID           string
	Code         string
	Name         string
	ImageURL     string
	Price        int
	Description  string
	Active       bool
	CategoryID   string
	CategoryName string
}
