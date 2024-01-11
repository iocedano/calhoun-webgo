package pages

import (
	"net/http"
	"training/webgo/internal"
)

var Contact = internal.CreatePage("contact", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	return nil
})
