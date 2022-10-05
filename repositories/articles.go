package repositories

import (
	"database/sql"
	"github.com/rinonkia/go_api_tutorial/models"
)

const articleNumPerPage = 5

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) InsertArticle(article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values 
		(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := r.db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	newArticle.ID = int(id)
	return newArticle, nil
}

func (r *ArticleRepository) SelectArticleList(page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice from articles limit ? offset ?
	`

	rows, err := r.db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}
	return articleArray, nil
}

func (r *ArticleRepository) SelectArticleDetail(id int) (models.Article, error) {
	const sqlStr = `
		select * from articles where article_id = ?;
		`

	row := r.db.QueryRow(sqlStr, id)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func (r *ArticleRepository) UpdateNiceNum(articleID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice from articles where article_id = ?	
`
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
