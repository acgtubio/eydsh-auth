package routes

import (
	"fmt"
	"net/http"

	"github.com/acgtubio/eydsh-auth/handlers"
	"github.com/acgtubio/eydsh-auth/internal/auth"
)

func addRoutes(
	mux *http.ServeMux,
	authService auth.AuthService,
) {
	mux.HandleFunc("/api/login", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hello")
	})
	mux.Handle("/api/login2", handlers.HandleLogin(authService))
	mux.Handle("/api/register", handlers.HandleRegister(authService))
}

func NewServer(
	authService auth.AuthService,
) http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
		authService,
	)

	return mux
}
