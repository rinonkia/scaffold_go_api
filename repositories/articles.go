package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/rinonkia/scaffold_golang/app/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const articleNumPerPage = 5

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) InsertArticle(ctx context.Context, article models.Article) error {
	return article.Insert(ctx, r.db, boil.Infer())
}

func (r *ArticleRepository) SelectArticleList(ctx context.Context, page int) ([]*models.Article, error) {
	articles, err := models.Articles().All(ctx, r.db)

	// TODO: pagination
	log.Printf("pagination: %d, page %d", articleNumPerPage, page)

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) SelectArticleDetail(ctx context.Context, id int) (*models.Article, error) {
	article, err := r.getArticle(ctx, id)
	if err != nil {
		return nil, err
	}
	// TODO: commentの取得
	return article, nil
}

func (r *ArticleRepository) UpdateNiceNum(ctx context.Context, articleID int) (*models.Article, error) {
	// TODO: transaction
	// TODO: models.Articleを返す
	article, err := r.getArticle(ctx, articleID)
	if err != nil {
		return nil, err
	}

	article.Nice++
	_, err = article.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *ArticleRepository) getArticle(ctx context.Context, articleID int) (*models.Article, error) {
	article, err := models.Articles(models.ArticleWhere.ArticleID.EQ(uint(articleID))).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return article, nil
}
