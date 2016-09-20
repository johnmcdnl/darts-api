package handlers

import (
	"github.com/goware/cors"
)

//TODO Force users to define what params are
func CORSHandler() *cors.Cors {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{""},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	return corsHandler
}
