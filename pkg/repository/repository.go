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
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Advert: NewAdvertPostgres(db),
	}
}
