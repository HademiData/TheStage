package main

import (
    "auth-service/internals/handlers" // Ensure this matches your actual package name
    "auth-service/internals/repository"
    "auth-service/internals/service"
    "log"
    "net/http"

    "github.com/gorilla/mux" // You must import the package
)

func main() {
    repo := repository.NewMemoryRepository()
    authService := service.NewAuthService(repo)
    authHandler := handler.NewAuthHandler(authService)

    // Initialize the gorilla/mux router
    r := mux.NewRouter()

    // Static file serving
    fs := http.FileServer(http.Dir("./static"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

    // Routes

    r.HandleFunc("/", authHandler.HomePage).Methods(http.MethodGet)

    r.HandleFunc("/register", authHandler.RegisterPage).Methods(http.MethodGet)
    r.HandleFunc("/login", authHandler.LoginPage).Methods(http.MethodGet)

    r.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)
    r.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)

    r.HandleFunc("/dashboard", authHandler.DashboardPage).Methods(http.MethodGet)

    log.Println("running on :8080")

    http.ListenAndServe(":8080", r)
}