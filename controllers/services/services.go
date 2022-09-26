package services

import "github.com/rinonkia/go_api_tutorial/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

	PostCommentService(comment models.Comment) (models.Comment, error)
}
