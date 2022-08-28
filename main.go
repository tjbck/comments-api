package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"comments-api/routers"
	"comments-api/utils"
)

func main() {
	PORT := "9000"

	utils.LogInfo(PORT)

	// collection := db.GetCollection("comments")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/comments", routers.CommentsRouter())
	r.Mount("/users", routers.UsersRouter())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
