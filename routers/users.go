package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UsersRouter() chi.Router {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome users"))
	})

	return r

}
