package api

import (
	"net/http"
	"training/webgo/internal"
)

type User struct {
	UserName string `json:"user"`
}

var UsersSignUp = internal.CreateAPIEndpoint(func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	return User{
		UserName: Request.FormValue("email"),
	}
})
