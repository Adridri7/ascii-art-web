package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("../web/templates/" + tmpl + ".html")
	if err != nil {
		// En cas d'erreur lors du parsing du template
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusInternalServerError)+". INTERNAL SERVER ERROR", http.StatusInternalServerError)
		fmt.Printf("Erreur lors du parsing du template: %s.html", tmpl)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusInternalServerError)+". INTERNAL SERVER ERROR", http.StatusInternalServerError)
		fmt.Printf("Erreur lors de l'exécution du template: %s", err)
		return
	}
}
