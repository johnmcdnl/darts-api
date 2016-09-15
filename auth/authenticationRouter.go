package auth

import (
	"github.com/pressly/chi"
	"github.com/johnmcdnl/auth/auth"
	"github.com/goware/cors"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsHandler.Handler)
	r.Post("/login", auth.LoginHandler)
	r.Post("/register", auth.RegisterHandler)
	return r
}