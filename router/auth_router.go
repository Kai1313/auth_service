package router

import (
	"AuthService/handlers"
	"AuthService/repositories"
	"github.com/gorilla/mux"
)

func AuthRouter(repository repositories.AuthRepository, router *mux.Router) {
	AuthHandler := handlers.NewAuthRepository(&repository)

	router.Use(handlers.CorsMiddleware)

	router.HandleFunc("/api/login", AuthHandler.LoginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/verifyToken", AuthHandler.VerifyTokenHandler).Methods("POST", "OPTIONS")
}
