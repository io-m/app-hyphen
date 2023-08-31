package helpers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetUrlParam(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}
