package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rinonkia/go_api_tutorial/app/models"
	"github.com/rinonkia/go_api_tutorial/services/interfaces"
)

type ArticleController struct {
	article interfaces.ArticleService
}

func NewArticleController(s interfaces.ArticleService) *ArticleController {
	return &ArticleController{article: s}
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		log.Print(err)
		return
	}

	err := c.article.PostArticleService(context.Background(), article)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if err = json.NewEncoder(w).Encode("OK"); err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
	}
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			log.Print(err)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.article.GetArticleListService(context.Background(), page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if err = json.NewEncoder(w).Encode(articleList); err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
	}
}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		log.Print(err)
		return
	}

	article, err := c.article.GetArticleService(context.Background(), articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if err = json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
	}
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: PDO実装
	var articleID int
	log.Printf("%T\n", req.Body)
	log.Printf("%s\n", req.Body)
	if err := json.NewDecoder(req.Body).Decode(&articleID); err != nil {
		http.Error(w, "fail to decode json.", http.StatusBadRequest)
		log.Print(err)
		return
	}

	updatedArticle, err := c.article.PostNiceService(context.Background(), articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	if err = json.NewEncoder(w).Encode(updatedArticle); err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
	}
}
