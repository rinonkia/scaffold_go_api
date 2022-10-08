package controllers

import (
	"context"
	"encoding/json"
	"github.com/rinonkia/go_api_tutorial/app/models"
	"github.com/rinonkia/go_api_tutorial/services/interfaces"
	"log"
	"net/http"
)

type CommentController struct {
	comment interfaces.CommentService
}

func NewCommentController(s interfaces.CommentService) *CommentController {
	return &CommentController{comment: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: PDO実装
	var comment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(w, "fail to decode json.", http.StatusBadRequest)
		log.Print(err)
		return
	}

	err := c.comment.PostCommentService(context.Background(), comment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode("OK")
}
