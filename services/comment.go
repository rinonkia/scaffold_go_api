package services

import (
	"github.com/rinonkia/go_api_tutorial/apperrors"
	"github.com/rinonkia/go_api_tutorial/models"
	"github.com/rinonkia/go_api_tutorial/repositories/interfaces"
)

type CommentService struct {
	comment interfaces.CommentRepository
}

func NewCommentService(comment interfaces.CommentRepository) *CommentService {
	return &CommentService{comment: comment}
}
func (s *CommentService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := s.comment.InsertComment(comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return newComment, nil
}
