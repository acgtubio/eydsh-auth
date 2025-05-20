package handlers

import (
	"fmt"
	"net/http"
)

type loginHandler struct{}

func (*loginHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "served\r\n")
}

func HandleLogin() http.Handler {
	return &loginHandler{}
}
