package routes

import (
	"fmt"
	"github.com/acgtubio/eydsh-auth/handlers"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/api/login", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hello")
	})
	mux.Handle("/api/login2", handlers.HandleLogin())
}

func NewServer() http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
	)

	return mux
}
