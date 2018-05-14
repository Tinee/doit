package http

import (
	"net/http"

	"github.com/Tinee/doit/domain"

	"github.com/go-chi/chi"
)

// UserHandler is the struct that contains and handle all routing for credentials.
type UserHandler struct {
	*chi.Mux
	UserRepository domain.UserRepository
}

// NewCredentialHandler creates and setup the routes for credentials.
func NewCredentialHandler(r domain.UserRepository) *UserHandler {
	h := &UserHandler{
		Mux:            chi.NewRouter(),
		UserRepository: r,
	}

	h.Get("/api/user", h.HandleRegisterPost)
	return h
}

// HandleRegisterPost handles when a user want's to register on our platform.
func (h *UserHandler) HandleRegisterPost(w http.ResponseWriter, r *http.Request) {
	h.UserRepository.Create(domain.User{
		Email: "Marcus.Carssl@asd.com",
	})
}
