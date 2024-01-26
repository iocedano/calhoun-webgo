package pages

import (
	"net/http"
	"training/webgo/internal"
)

var SignUp = internal.CreatePage("signup", func(w http.ResponseWriter, r *http.Request) interface{} {
	return nil
})
