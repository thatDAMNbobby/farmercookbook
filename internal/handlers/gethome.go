package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type HomeHandler struct {
	tmpl *template.Template
}

type HomeHandlerDeps struct {
	Tmpl *template.Template
}

func NewHomeHandler(deps HomeHandlerDeps) *HomeHandler {
	return &HomeHandler{
		tmpl: deps.Tmpl,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")
	err := h.tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
