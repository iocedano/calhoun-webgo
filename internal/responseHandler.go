package internal

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type CustomResponseHandler func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{}

type IResponseHandler interface {
	ServeHTTP(ResponseWriter http.ResponseWriter, Request *http.Request)
	Handler(ResponseWriter http.ResponseWriter, Request *http.Request) interface{}
	WriteResponse(w http.ResponseWriter, data interface{})
}

// Page ---
type Page struct {
	name     string
	handler  CustomResponseHandler
	template *template.Template
}

func (ph *Page) WriteResponse(w http.ResponseWriter, data interface{}) {
	if err := ph.template.Execute(w, data); err != nil {
		log.Panicf("parsing template: %v", err)
		http.Error(w, "Parsing Error", http.StatusBadRequest)
	}
}

func (ph *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ph.WriteResponse(w, ph.handler(w, r))
}

// API Endpoint ---
type API struct {
	handler CustomResponseHandler
}

func (*API) WriteResponse(w http.ResponseWriter, data interface{}) {
	jData, err := json.Marshal(data)
	if err != nil {
		log.Panicf("marshal error: %v", err)
		http.Error(w, "Marshal Error", http.StatusBadRequest)
		return
	}
	w.Write(jData)
}

func (endpoint *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint.WriteResponse(w, endpoint.handler(w, r))
}

func CreatePage(name string, handler CustomResponseHandler) *Page {
	return &Page{
		name:     name,
		handler:  handler,
		template: GetFSTemplate(name),
	}
}

func CreateAPIEndpoint(handler CustomResponseHandler) *API {
	return &API{
		handler: handler,
	}
}
