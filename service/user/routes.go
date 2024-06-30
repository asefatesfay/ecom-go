package user

import (
	"fmt"
	"net/http"

	"github.com/asefatesfay/ecom-go/service/auth"
	"github.com/asefatesfay/ecom-go/types"
	"github.com/asefatesfay/ecom-go/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	var registerPayload types.RegisterPayload
	if err := utils.ParseJSON(r, &registerPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	_, err := h.store.GetUserByEmail(registerPayload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("User with email %s already exists", registerPayload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(registerPayload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: registerPayload.FirstName,
		LastName:  registerPayload.LastName,
		Email:     registerPayload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
