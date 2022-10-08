package repositories

import (
	"context"
	"database/sql"

	"github.com/rinonkia/go_api_tutorial/app/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) InsertComment(ctx context.Context, comment models.Comment) error {
	return comment.Insert(ctx, r.db, boil.Infer())
}

func (r *CommentRepository) SelectCommentList(ctx context.Context, articleID int) ([]*models.Comment, error) {
	comments, err := models.Comments(models.CommentWhere.ArticleID.EQ(uint(articleID))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
