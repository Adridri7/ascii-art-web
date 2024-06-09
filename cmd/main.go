package main

import (
	"ascii"
	"fmt"
	"handlers"
	"net/http"
	"strconv"
)

var Text string

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusNotFound)+". PAGE NOT FOUND ", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		handlers.RenderTemplate(w, "home")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//handlers.RenderTemplate(w, "home")
		text := r.FormValue("input")
		theme := r.FormValue("themes")
		http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)
		Text = ascii.CheckArgs([]string{text, theme})
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func DisplayResult(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL)
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusNotFound)+". PAGE NOT FOUND ", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		handlers.RenderTemplate(w, "ascii-art")

	}
	fmt.Fprintf(w, Text)

	// Display to the screen

}

func main() {
	// Setting up the home page
	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", DisplayResult)

	// Start the serv on port : 8080
	port := 8080
	adresse := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost:%d...\n", port)
	http.ListenAndServe(adresse, nil)
}
