package services

import (
	"context"

	"github.com/rinonkia/scaffold_go_api/app/models"
	"github.com/rinonkia/scaffold_go_api/apperrors"
	"github.com/rinonkia/scaffold_go_api/repositories/interfaces"
)

type ArticleService struct {
	article interfaces.ArticleRepository
	comment interfaces.CommentRepository
}

func NewArticleService(
	article interfaces.ArticleRepository,
	comment interfaces.CommentRepository,
) *ArticleService {
	return &ArticleService{
		article: article,
		comment: comment,
	}
}

func (s *ArticleService) GetArticleService(ctx context.Context, articleID int) (*models.Article, error) {
	article, err := s.article.SelectArticleDetail(ctx, articleID)
	if err != nil {
		return nil, err
	}

	// commentList, err := s.comment.SelectCommentList(ctx, articleID)
	// if err != nil {
	// 	return models.Article{}, err
	// }
	// article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *ArticleService) PostArticleService(ctx context.Context, article models.Article) error {
	err := s.article.InsertArticle(ctx, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return err
	}

	return nil
}

func (s *ArticleService) GetArticleListService(ctx context.Context, page int) ([]*models.Article, error) {
	articleList, err := s.article.SelectArticleList(ctx, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *ArticleService) PostNiceService(ctx context.Context, articleID int) (*models.Article, error) {
	article, err := s.article.UpdateNiceNum(ctx, articleID)
	if err != nil {
		return nil, err
	}

	return article, nil
}
