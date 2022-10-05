package interfaces

import "github.com/rinonkia/go_api_tutorial/models"

type CommentService interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
