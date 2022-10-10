package AdvertAPI

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt string `json:"created_at"`
	//UpdatedAt string `json:"updated_at"`
	FavList string`json:"fav_list"`
}

