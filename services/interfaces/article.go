package interfaces

import (
	"context"
	"github.com/rinonkia/go_api_tutorial/app/models"
)

type ArticleService interface {
	PostArticleService(ctx context.Context, article models.Article) error
	GetArticleService(ctx context.Context, articleID int) (*models.Article, error)
	GetArticleListService(ctx context.Context, page int) ([]*models.Article, error)
	PostNiceService(ctx context.Context, articleID int) (*models.Article, error)
}
