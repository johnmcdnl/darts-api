package board

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"net/http"
)

func BoardsRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getBoardHandler)

	return r
}

func getBoardHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, newBoard())
}
