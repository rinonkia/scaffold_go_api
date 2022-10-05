package services

import (
	"github.com/rinonkia/go_api_tutorial/apperrors"
	"github.com/rinonkia/go_api_tutorial/models"
	"github.com/rinonkia/go_api_tutorial/repositories"
)

type CommentService struct {
	repository repositories.CommentRepository
}

func NewCommentService(repository *repositories.CommentRepository) *CommentService {
	return &CommentService{repository: *repository}
}
func (s *CommentService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := s.repository.InsertComment(comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return newComment, nil
}
