package interfaces

import (
	"context"

	"github.com/rinonkia/scaffold_go_api/app/models"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment models.Comment) error
	SelectCommentList(ctx context.Context, articleID int) ([]*models.Comment, error)
}
