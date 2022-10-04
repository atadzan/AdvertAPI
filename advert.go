package AdvertAPI

type AdvertInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Location    string  `json:"location"`
	PhoneNumber string  `json:"phone_number"`
	Price       int     `json:"price"`
	Images      []Image `json:"images"`
}

type AdvertInfo struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Location    string  `json:"location"`
	PhoneNumber string  `json:"phone_number"`
	Price       int     `json:"price"`
	PublishDate string  `json:"publish_date"`
	Views       string  `json:"views"`
	Images      []Image `json:"images"`
}
type Image struct {
	ImagePath string `json:"image_path"`
}