package handlers

import (
	"ascii"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var Text string

func RenderTemplate(w http.ResponseWriter, tmpl string, s string) {
	t, err := template.ParseFiles("./web/templates/" + tmpl + ".html")
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

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusNotFound)+". PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "home", "")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		text, theme := r.FormValue("input"), r.FormValue("themes")
		lines := ascii.ThemeToLines(theme)
		input, err := ascii.GetTextInput(text)

		if err != nil {
			http.Error(w, "ERROR: "+strconv.Itoa(http.StatusInternalServerError)+". INTERNAL SERVER ERROR", http.StatusInternalServerError)
		} else {
			Text = ascii.PrintAsciiArt(input, lines)
			SaveOutput(w, Text)
			http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)
		}

	default:
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusMethodNotAllowed)+". METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		log.Println("ERROR: " + strconv.Itoa(http.StatusMethodNotAllowed) + ". METHOD NOT ALLOWED")
	}
}

func DisplayResult(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusNotFound)+". PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "ascii-art", Text)
		//w.Header().Set("Content-Disposition", "attachment; filename=res.txt")
		//w.Write()
	default:
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusMethodNotAllowed)+". METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		log.Println("ERROR: " + strconv.Itoa(http.StatusMethodNotAllowed) + ". METHOD NOT ALLOWED")
	}
}

func SaveOutput(w http.ResponseWriter, text string) {
	output := CreateFile(w, "./web/download/res.txt")
	output.WriteString(text)
}

func CreateFile(w http.ResponseWriter, path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusInternalServerError)+". INTERNAL SERVER ERROR", http.StatusInternalServerError)
	}
	return file
}
