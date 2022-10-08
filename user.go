package AdvertAPI

import "time"

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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FavList []int `json:"fav_list"`
}

