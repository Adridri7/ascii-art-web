package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, s string) {
	t, err := template.ParseFiles("web/templates/" + tmpl + ".html")
	if err != nil {
		// En cas d'erreur lors du parsing du template
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusBadRequest)+". BAD REQUEST", http.StatusBadRequest)
		fmt.Println("Erreur : bad request")
		return
	}

	if err := t.Execute(w, s); err != nil {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusBadRequest)+". BAD REQUEST", http.StatusBadRequest)
		fmt.Println("Erreur : bad request")
		return
	}
}
