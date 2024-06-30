package unit_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asefatesfay/ecom-go/service/user"
	"github.com/asefatesfay/ecom-go/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := user.NewHandler(userStore)

	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterPayload{
			FirstName: "adonai",
			LastName:  "testfay",
			Email:     "adonaitesfay@gamil.com",
			Password:  "supersecret",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (store *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("User with email %s not found", email)
}
func (store *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (store *mockUserStore) CreateUser(user types.User) error {
	return nil
}
