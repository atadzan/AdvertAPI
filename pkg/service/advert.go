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

func(s *AdvertService) GetAll()([]AdvertAPI.AdvertInfo, error){
	return s.repo.GetAll()
}

func(s *AdvertService) GetById(id int)(AdvertAPI.AdvertInfo, error){
	return s.repo.GetById(id)
}
