package services

import (
	"context"

	"github.com/rinonkia/scaffold_golang/app/models"
	"github.com/rinonkia/scaffold_golang/apperrors"
	"github.com/rinonkia/scaffold_golang/repositories/interfaces"
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
