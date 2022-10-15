package service

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/repository"
)

type AdvertService struct {
	repo repository.Advert
}

func NewAdvertService(repo repository.Advert) *AdvertService{
	return &AdvertService{repo: repo}
}

func(s *AdvertService) Add(advert AdvertAPI.AdvertInput)(int, error){
	return s.repo.Add(advert)
}

func(s *AdvertService) GetAll(advertPerPage, offset int)([]AdvertAPI.AdvertOutput, error){
	return s.repo.GetAll(advertPerPage, offset)
}

func(s *AdvertService) GetById(id int)(AdvertAPI.AdvertOutput, error){
	return s.repo.GetById(id)
}

func(s *AdvertService) CountAdverts()(int, error){
	return s.repo.CountAdverts()
}

func(s *AdvertService) GetImage(id int)([]AdvertAPI.AdvertImage, error){
	return s.repo.GetImage(id)
}

func(s *AdvertService) Delete(id int)error{
	return s.repo.Delete(id)
}

func(s *AdvertService) Update(id int, advert AdvertAPI.AdvertInput) error{
	return s.repo.Update(id, advert)
}

func(s *AdvertService) AddFav(userId, advertId int) error{
	return s.repo.AddFav(userId, advertId)
}

func(s *AdvertService) GetFav(userId int)([]AdvertAPI.AdvertOutput, error){
	return s.repo.GetFav(userId)
}

func(s *AdvertService) DeleteFav(userId, advertId int) error{
	return s.repo.DeleteFav(userId, advertId)
}

func(s *AdvertService) CheckFavList(userId, advertId int)(bool, error){
	return s.repo.CheckFavList(userId, advertId)
}

func(s *AdvertService) Search(search string)([]AdvertAPI.AdvertOutput, error){
	return s.repo.Search(search)
}
