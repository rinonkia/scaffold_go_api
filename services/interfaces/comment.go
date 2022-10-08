package interfaces

import (
	"context"

	"github.com/rinonkia/scaffold_golang/app/models"
)

type CommentService interface {
	PostCommentService(ctx context.Context, comment models.Comment) error
}
