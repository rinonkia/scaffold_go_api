package interfaces

import (
	"context"
	"github.com/rinonkia/go_api_tutorial/app/models"
)

type CommentService interface {
	PostCommentService(ctx context.Context, comment models.Comment) error
}
