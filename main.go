package main

import (
	"log"
	"net/http"

	"github.com/thongsoi/jwt/handlers"

	"github.com/thongsoi/jwt/db"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Auth routes
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST") // New route

	// Protected routes
	r.Handle("/dashboard", handlers.AuthMiddleware(http.HandlerFunc(handlers.DashboardHandler))).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
