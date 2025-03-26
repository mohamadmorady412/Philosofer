package routes

import (
	"github.com/gorilla/mux"
	"your_project/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.Handle("/auth/login", middlewares.IPBlockerMiddleware(middlewares.RateLimiterMiddleware(http.HandlerFunc(handlers.LoginHandler)))).Methods("POST")
	r.Handle("/auth/register", middlewares.IPBlockerMiddleware(middlewares.RateLimiterMiddleware(http.HandlerFunc(handlers.RegisterHandler)))).Methods("POST")
	
	r.HandleFunc("/auth/google/login", handlers.GoogleLoginHandler).Methods("GET")
	r.HandleFunc("/auth/google/callback", handlers.GoogleCallbackHandler).Methods("GET")

	r.HandleFunc("/auth/github/login", handlers.GitHubLoginHandler).Methods("GET")
	r.HandleFunc("/auth/github/callback", handlers.GitHubCallbackHandler).Methods("GET")
}
