package handlers

import (
	"fmt"
	"net/http"

	"github.com/acgtubio/eydsh-auth/internal/auth"
	"github.com/acgtubio/eydsh-auth/model"
)

type registerHandler struct {
	auth auth.AuthService
}

func (handler *registerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	user := model.User{
		Id:       "asdadf",
		Username: "First user",
		Password: "asdfadf",
		Salt:     "setsets",
	}

	err := handler.auth.Register(user)

	if err != nil {
		http.Error(w, "error register", http.StatusInternalServerError)
	}

	fmt.Fprint(w, "registered")
}

func HandleRegister(authService auth.AuthService) http.Handler {
	return &registerHandler{
		auth: authService,
	}
}
