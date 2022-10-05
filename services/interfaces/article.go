package interfaces

import "github.com/rinonkia/go_api_tutorial/models"

type ArticleService interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}
