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

func(s *AdvertService) GetAll(advertPerPage, offset int)([]AdvertAPI.AdvertInfo, error){
	return s.repo.GetAll(advertPerPage, offset)
}

func(s *AdvertService) GetById(id int)(AdvertAPI.AdvertInfo, error){
	return s.repo.GetById(id)
}

func(s *AdvertService) CountAdverts()(int, error){
	return s.repo.CountAdverts()
}

func(s *AdvertService) AddDB(fname, ftype, filepath string, fsize int64)(string, error){
	return s.repo.AddDB(fname, ftype, filepath, fsize)
}

func(s *AdvertService) GetImage(id int)([]AdvertAPI.AdvertImage, error){
	return s.repo.GetImage(id)
}

