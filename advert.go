package AdvertAPI

type AdvertInput struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Location string `json:"location"`
	PhoneNumber string `json:"phone_number"`
	Price int `json:"price"`
	Images []Image `json:"images"`
}

type Image struct{
	ImagePath string `json:"image_path"`
}

type AdvertInfo struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Location string `json:"location"`
	PhoneNumber string `json:"phone_number"`
	Price int `json:"price"`
	//PublishDate  `json:"publish_date"`
	Views int `json:"views"`
}


