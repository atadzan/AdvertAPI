package service

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService{
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Add(category AdvertAPI.CategoryInput) (int, error){
	return s.repo.Add(category)
}

func(s *CategoryService) GetMain()([]AdvertAPI.CategoryOutput, error){
	return s.repo.GetMain()
}

func(s *CategoryService) GetNested(categoryId int)([]AdvertAPI.CategoryOutput, error){
	return s.repo.GetNested(categoryId)
}

func(s *CategoryService) GetCategoryAdverts(categoryId int)([]AdvertAPI.AdvertOutput, error){
	return s.repo.GetCategoryAdverts(categoryId)
}
