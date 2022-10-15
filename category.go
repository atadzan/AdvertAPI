package AdvertAPI

type CategoryInput struct {
	Title string `json:"title"`
	Main  bool   `json:"main"`
	ParentId int `json:"parent_id"`
}

type CategoryOutput struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Main     bool   `json:"main"`
	ParentId int    `json:"parent_id"`
}
