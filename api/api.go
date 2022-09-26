package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rinonkia/go_api_tutorial/controllers"
	"github.com/rinonkia/go_api_tutorial/services"
	"net/http"
)

func NewRouter(db *sql.DB) *mux.Router {
	service := services.NewMyAppService(db)
	article := controllers.NewArticleController(service)
	comment := controllers.NewCommentController(service)
	r := mux.NewRouter()

	r.HandleFunc("/article", article.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", article.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", article.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", article.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", comment.PostCommentHandler).Methods(http.MethodPost)

	return r
}
