package models

type Image struct {
	Caption string
	Url     string
}
type Article struct {
	Title   string
	Content string
	Credit  string
	Writer  string
	Url     string
	Date    string
	Images  []Image
}
