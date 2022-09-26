package controllers

import (
	"encoding/json"
	services2 "github.com/rinonkia/go_api_tutorial/controllers/services"
	"github.com/rinonkia/go_api_tutorial/models"
	"log"
	"net/http"
)

type CommentController struct {
	service services2.CommentServicer
}

func NewCommentController(s services2.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
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
