package service

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/repository"
)

type Advert interface {
	Add(advert AdvertAPI.AdvertInput) (int, error)
	GetAll(advertPerPage, offset int) ([]AdvertAPI.AdvertInfo, error)
	GetById(id int) (AdvertAPI.AdvertInfo, error)
	GetImage(id int) ([]AdvertAPI.AdvertImage, error)
	CountAdverts() (int, error)
	Delete(id int) error
	Update(id int, advert AdvertAPI.AdvertInput) error
	AddFav(userId, advertId int) error
	GetFav(userId int) ([]AdvertAPI.AdvertInfo, error)
	DeleteFav(userId, advertId int) error
	Search(search string) ([]AdvertAPI.AdvertInfo, error)
}

type Comment interface {
	AddCom(comment AdvertAPI.InputComm, userId, advertId int) error
	DelCom(advertId, userId, commentId int) error
	UpdCom(comment AdvertAPI.InputComm, userId, advertId, commentId int) error
}

type Category interface {
	Add(category AdvertAPI.CategoryInput) (int, error)
	GetMain()([]AdvertAPI.CategoryOutput, error)
	GetNested(categoryId int)([]AdvertAPI.CategoryOutput, error)
}

type Authorization interface {
	CreateUser(user AdvertAPI.SignUpInput) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Advert
	Authorization
	Comment
	Category
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Advert:        NewAdvertService(repos.Advert),
		Authorization: NewAuthService(repos.Authorization),
		Comment:       NewCommentService(repos.Comment),
		Category:      NewCategoryService(repos.Category),
	}
}
