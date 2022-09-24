package services

import (
	"github.com/rinonkia/go_api_tutorial/models"
	"github.com/rinonkia/go_api_tutorial/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}
