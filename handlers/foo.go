package handlers

import (
	"net/http"

	"github.com/HsiaoCz/gohtt/views/foo"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return foo.Index().Render(r.Context(), w)
}
