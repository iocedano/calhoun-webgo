package pages

import (
	"net/http"
	"training/webgo/internal"
)

var NotFoudPage = internal.CreatePage("error", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	ResponseWriter.WriteHeader(http.StatusNotFound)
	return nil
})
