package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rinonkia/go_api_tutorial/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!!")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	// 出力することで暫定的にエラーを回避
	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 出力することで暫定的にエラーを回避
	log.Println(articleID)

	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "fail to decode json.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(w, "fail to decode json.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
