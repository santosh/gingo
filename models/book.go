package models

// Book represents data about a book.
type Book struct {
	ISBN   string `json:"isbn" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
