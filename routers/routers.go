package routers

import (
	"github.com/gorilla/mux"
	"github.com/rinonkia/go_api_tutorial/controllers"
	"net/http"
)

func NewRouter(articleCon *controllers.ArticleController, commentCon *controllers.CommentController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/article", articleCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", articleCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", articleCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", articleCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", commentCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
