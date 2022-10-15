package AdvertAPI

type AdvertInput struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Category    int           `json:"category"`
	Location    string        `json:"location"`
	PhoneNumber string        `json:"phone_number"`
	Price       int           `json:"price"`
	UserId      int           `json:"user_id"`
	Images      []AdvertImage `json:"images"`
}

type AdvertOutput struct {
	Id           int        `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Category     int        `json:"category"`
	Location     string     `json:"location"`
	PhoneNumber  string     `json:"phone_number"`
	Price        int        `json:"price"`
	PublishDate  string     `json:"publish_date"`
	UserId       int        `json:"user_id"`
	Views        int        `json:"views"`
	ImagesCount  int        `json:"images_count"`
	CommentCount int        `json:"comment_count"`
	Images       []ImageUrl `json:"images"`
	Comments     []Comment  `json:"comments"`
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

type InputComm struct {
	Body string `json:"body"`
}

type Comment struct {
	Id        int    `json:"id"`
	AdvertId  int    `json:"advert_id"`
	Body      string `json:"body"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
