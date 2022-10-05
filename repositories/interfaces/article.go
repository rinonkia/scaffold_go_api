package interfaces

import (
	"github.com/rinonkia/go_api_tutorial/models"
)

type ArticleRepository interface {
	InsertArticle(article models.Article) (models.Article, error)
	SelectArticleList(page int) ([]models.Article, error)
	SelectArticleDetail(id int) (models.Article, error)
	UpdateNiceNum(articleID int) error
}
