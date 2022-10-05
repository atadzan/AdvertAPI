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
	AddDB(fname, ftype, filepath string, fsize int64)(string, error)
	GetImage(id int)([]AdvertAPI.AdvertImage, error)
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Advert: NewAdvertPostgres(db),
	}
}
