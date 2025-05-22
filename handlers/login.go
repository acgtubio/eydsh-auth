package handlers

import (
	"fmt"
	"net/http"

	"github.com/acgtubio/eydsh-auth/internal/auth"
)

type loginHandler struct{}

func (*loginHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "served\r\n")
}

func HandleLogin(authService auth.AuthService) http.Handler {
	return &loginHandler{}
}
