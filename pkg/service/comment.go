package service

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/repository"
)

type CommentService struct{
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService{
	return &CommentService{repo: repo}
}

func(s *CommentService) AddCom(comment AdvertAPI.InputComm, userId, advertId int) error {
	return s.repo.AddCom(comment, userId, advertId)
}

func(s *CommentService) DelCom(advertId, userId, commentId int) error{
	return s.repo.DelCom(advertId, userId, commentId)
}

func(s *CommentService) UpdCom(comment AdvertAPI.InputComm, userId, advertId, commentId int) error{
	return s.repo.UpdCom(comment, userId, advertId, commentId)
}
