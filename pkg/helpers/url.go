package helpers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUrlParam(r *http.Request, param string) uuid.UUID {
	return uuid.MustParse(chi.URLParam(r, param))
}
