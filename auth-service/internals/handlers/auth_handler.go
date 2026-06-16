package handler

import (
	"auth-service/internals/service"
	"html/template"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

type DashboardData struct {
	Name  string
	Role  string
	Email string
}

func (h *AuthHandler) DashboardPage(w http.ResponseWriter, r *http.Request) {

	data := DashboardData{
		Name:  "Afolabi Adewale",
		Email: "afolabi@stage.com",
		Role:  "speaker",
	}

	tmpl := template.Must(
		template.ParseFiles("templates/dashboard.html"),
	)

	tmpl.Execute(w, data)
}

func (h *AuthHandler) HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(
		template.ParseFiles("templates/home.html"),
	)

	tmpl.Execute(w, nil)
}

func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(
		template.ParseFiles("templates/register.html"),
	)

	tmpl.Execute(w, nil)
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(
		template.ParseFiles("templates/login.html"),
	)

	tmpl.Execute(w, nil)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")

	err := h.service.Register(email, password, role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := h.service.Login(email, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
