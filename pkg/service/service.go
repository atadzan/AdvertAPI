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
	Delete(id int)error
	Update(id int, advert AdvertAPI.AdvertInput) error
	AddFav(userId, advertId int) error
	GetFav(userId int) ([]AdvertAPI.AdvertInfo, error)
	DeleteFav(userId, advertId int) error
}

type Authorization interface {
	CreateUser(user AdvertAPI.SignUpInput)(int, error)
	GenerateToken(username, password string)(string, error)
	ParseToken(token string)(int, error)
}

type Service struct{
	Advert
	Authorization
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Advert: NewAdvertService(repos.Advert),
		Authorization: NewAuthService(repos.Authorization),
	}
}