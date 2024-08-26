package controllers

import (
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type homeController struct{}

func NewHomeController() *homeController {
	return &homeController{}
}

func (c *homeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/home/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
