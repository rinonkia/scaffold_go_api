package services

import (
	"context"
	"github.com/rinonkia/go_api_tutorial/app/models"
	"github.com/rinonkia/go_api_tutorial/apperrors"
	"github.com/rinonkia/go_api_tutorial/repositories/interfaces"
)

type CommentService struct {
	comment interfaces.CommentRepository
}

func NewCommentService(comment interfaces.CommentRepository) *CommentService {
	return &CommentService{comment: comment}
}
func (s *CommentService) PostCommentService(ctx context.Context, comment models.Comment) error {
	err := s.comment.InsertComment(ctx, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return err
	}
	return nil
}
