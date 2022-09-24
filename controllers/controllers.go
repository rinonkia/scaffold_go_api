package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rinonkia/go_api_tutorial/models"
	"github.com/rinonkia/go_api_tutorial/services"
	"io"
	"log"
	"net/http"
	"strconv"
)

type MyAppController struct {
	service *services.MyAppService
}

func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

func (c *MyAppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!!")
}

func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		log.Print(err)
		return
	}

	newArticle, err := c.service.PostArticleService(article)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(newArticle)
}

func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		log.Print(err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "fail to decode json.", http.StatusBadRequest)
		log.Print(err)
		return
	}

	updatedArticle, err := c.service.PostNiceService(article)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(updatedArticle)
}

func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(w, "fail to decode json.", http.StatusBadRequest)
		log.Print(err)
		return
	}

	newComment, err := c.service.PostCommentService(comment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(newComment)
}
