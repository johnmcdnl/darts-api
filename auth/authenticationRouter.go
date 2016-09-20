package auth

import (
	"github.com/johnmcdnl/auth/auth"
	"github.com/johnmcdnl/darts/utils/handlers"
	"github.com/pressly/chi"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(handlers.CORSHandler().Handler)
	r.Post("/login", auth.LoginHandler)
	r.Post("/register", auth.RegisterHandler)
	return r
}
