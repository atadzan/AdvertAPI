package repository

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
)

type Advert interface {
	Add(advert AdvertAPI.AdvertInput)(int, error)
	GetAll(advertPerPage, offset int)([]AdvertAPI.AdvertInfo, error)
	GetById(id int)(AdvertAPI.AdvertInfo, error)
	CountAdverts()(int, error)
	AddDB(file AdvertAPI.AdvertImage)(string, error)
	GetImage(id int)([]AdvertAPI.AdvertImage, error)
	Delete(id int)error
	Update(id int, advert AdvertAPI.AdvertInput)error
}

type Authorization interface {
	CreateUser(user AdvertAPI.SignUpInput)(int, error)
	GetUser(username, password string)(AdvertAPI.User, error)
}

type Repository struct {
	Advert
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Advert: NewAdvertPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
