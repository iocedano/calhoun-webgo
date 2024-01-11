package pages

import (
	"net/http"
	"training/webgo/internal"
)

var Home = internal.CreatePage("home", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	return nil
})
