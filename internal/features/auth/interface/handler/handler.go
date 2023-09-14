package auth_handler_interface

import "net/http"

type IAuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	OAuth(w http.ResponseWriter, r *http.Request)
}
