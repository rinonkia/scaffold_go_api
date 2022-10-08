package interfaces

import (
	"context"

	"github.com/rinonkia/scaffold_golang/app/models"
)

type ArticleRepository interface {
	InsertArticle(ctx context.Context, article models.Article) error
	SelectArticleList(ctx context.Context, page int) ([]*models.Article, error)
	SelectArticleDetail(ctx context.Context, id int) (*models.Article, error)
	UpdateNiceNum(ctx context.Context, articleID int) (*models.Article, error)
}
