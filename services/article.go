package services

import (
	"github.com/rinonkia/go_api_tutorial/apperrors"
	"github.com/rinonkia/go_api_tutorial/models"
	"github.com/rinonkia/go_api_tutorial/repositories/interfaces"
)

type ArticleService struct {
	articleRepository interfaces.ArticleRepository
	commentRepository interfaces.CommentRepository
}

func NewArticleService(
	article interfaces.ArticleRepository,
	comment interfaces.CommentRepository,
) *ArticleService {
	return &ArticleService{
		articleRepository: article,
		commentRepository: comment,
	}
}

func (s *ArticleService) GetArticleService(articleID int) (models.Article, error) {
	article, err := s.articleRepository.SelectArticleDetail(articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := s.commentRepository.SelectCommentList(articleID)
	if err != nil {
		return models.Article{}, err
	}
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *ArticleService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := s.articleRepository.InsertArticle(article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *ArticleService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := s.articleRepository.SelectArticleList(page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Article{}, err
	}

	if len(articleList) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *ArticleService) PostNiceService(article models.Article) (models.Article, error) {

	err := s.articleRepository.UpdateNiceNum(article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
