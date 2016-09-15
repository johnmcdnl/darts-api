package darts

import (
	"github.com/johnmcdnl/darts/resources"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/johnmcdnl/darts/auth"
	"net/http"
)

func StartServer() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/darts/api", resources.ResourcesRouter())
	r.Mount("/darts/api/auth", auth.AuthRouter())

	http.ListenAndServe(":4500", r)
}