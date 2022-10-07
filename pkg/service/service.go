package service

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/repository"
)

type Advert interface {
	Add(advert AdvertAPI.AdvertInput)(int, error)
	GetAll(advertPerPage, offset int)([]AdvertAPI.AdvertInfo, error)
	GetById(id int)(AdvertAPI.AdvertInfo, error)
	CountAdverts()(int, error)
	GetImage(id int)([]AdvertAPI.AdvertImage, error)
}

type Service struct{
	Advert
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Advert: NewAdvertService(repos.Advert),
	}
}