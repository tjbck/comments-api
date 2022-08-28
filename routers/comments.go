package routers

import (
	comments "comments-api/models/comments"

	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Comment = comments.Comment

type CommentForm struct {
	UserID  string `json:"userId"`
	Content string `json:"content"`
}

func addComment(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var formData CommentForm
	err := decoder.Decode(&formData)
	if err != nil {
		panic(err)
	}

	commentId := uuid.New().String()
	timestamp := time.Now().Unix()

	comment := Comment{
		ID:        commentId,
		ParentID:  "1",
		UserID:    formData.UserID,
		Content:   formData.Content,
		CreatedAt: int(timestamp),
		EditedAt:  int(timestamp),
	}

	comment_list := comments.AddComment(comment)

	render.JSON(w, r, comment_list)
}

func CommentsRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		comment_list := comments.GetComments()
		render.JSON(w, r, comment_list)
	})

	r.Get("/{commentId}", func(w http.ResponseWriter, r *http.Request) {

		commentId := chi.URLParam(r, "commentId")

		comment := comments.GetCommentsById(commentId)
		render.JSON(w, r, comment)
	})

	r.Post("/add", addComment)

	return r
}
