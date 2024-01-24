package pages

import (
	"net/http"
	"training/webgo/internal"
)

var SignUp = internal.CreatePage("signup", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	return nil
})
