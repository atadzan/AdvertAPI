package AdvertAPI

type AdvertInput struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Location    string        `json:"location"`
	PhoneNumber string        `json:"phone_number"`
	Price       int           `json:"price"`
	Images      []AdvertImage `json:"images"`
}

type AdvertInfo struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Location    string        `json:"location"`
	PhoneNumber string        `json:"phone_number"`
	Price       int           `json:"price"`
	PublishDate string        `json:"publish_date"`
	Views       string        `json:"views"`
	ImagesCount int           `json:"images_count"`
	Images      []ImageUrl `json:"images"`
}

type AdvertImage struct {
	Id       int
	Fname    string
	Fsize    int64
	Ftype    string
	Path     string
	AdvertId int
}

type ImageUrl struct {
	URL string
}

