package repository

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
)

type Advert interface {
	Add(advert AdvertAPI.AdvertInput) (int, error)
	GetAll(advertPerPage, offset int) ([]AdvertAPI.AdvertOutput, error)
	GetById(id int) (AdvertAPI.AdvertOutput, error)
	CountAdverts() (int, error)
	GetImage(id int) ([]AdvertAPI.AdvertImage, error)
	Delete(id int) error
	Update(id int, advert AdvertAPI.AdvertInput) error
	AddFav(userId, advertId int) error
	GetFav(userId int) ([]AdvertAPI.AdvertOutput, error)
	DeleteFav(userId, advertId int) error
	CheckFavList(userId, advertId int)(bool, error)
	Search(search string) ([]AdvertAPI.AdvertOutput, error)
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
	GetCategoryAdverts(categoryId int)([]AdvertAPI.AdvertOutput, error)
}

type Authorization interface {
	CreateUser(user AdvertAPI.SignUpInput) (int, error)
	GetUser(username, password string) (AdvertAPI.User, error)
}

type Repository struct {
	Advert
	Authorization
	Comment
	Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advert:        NewAdvertPostgres(db),
		Authorization: NewAuthPostgres(db),
		Comment:       NewCommentPostgres(db),
		Category:      NewCategoryPostgres(db),
	}
}
