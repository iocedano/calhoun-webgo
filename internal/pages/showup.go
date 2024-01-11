package pages

import (
	"net/http"
	"training/webgo/internal"

	"github.com/go-chi/chi/v5"
)

var ShowUp = internal.CreatePage("showup", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	data := struct {
		Username string
	}{
		Username: chi.URLParam(Request, "username"),
	}
	return data
})
