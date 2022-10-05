package interfaces

import "github.com/rinonkia/go_api_tutorial/models"

type CommentRepository interface {
	InsertComment(comment models.Comment) (models.Comment, error)
	SelectCommentList(articleID int) ([]models.Comment, error)
}
