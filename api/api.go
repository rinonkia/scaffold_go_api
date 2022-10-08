package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rinonkia/scaffold_golang/controllers"
	"github.com/rinonkia/scaffold_golang/repositories"
	"github.com/rinonkia/scaffold_golang/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	// repository
	articleRepository := repositories.NewArticleRepository(db)
	commentRepository := repositories.NewCommentRepository(db)

	// service
	articleService := services.NewArticleService(articleRepository, commentRepository)
	commentService := services.NewCommentService(commentRepository)

	// controller
	article := controllers.NewArticleController(articleService)
	comment := controllers.NewCommentController(commentService)
	r := mux.NewRouter()

	r.HandleFunc("/article", article.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", article.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", article.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", article.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", comment.PostCommentHandler).Methods(http.MethodPost)

	return r
}
