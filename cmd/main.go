package main

import (
	"ascii"
	"fmt"
	"handlers"
	"net/http"
	"os"
	"strconv"
)

var Text string

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusNotFound)+". PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		handlers.RenderTemplate(w, "home", "")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		text, theme := r.FormValue("input"), r.FormValue("themes")
		lines := ascii.ThemeToLines(theme)
		input := ascii.GetTextInput(text)
		Text = ascii.PrintAsciiArt(input, lines)
		SaveOutput(Text)
		http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)

	default:
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusMethodNotAllowed)+". METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		fmt.Println("ERROR: " + strconv.Itoa(http.StatusMethodNotAllowed) + ". METHOD NOT ALLOWED")
	}
}

func DisplayResult(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusNotFound)+". PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		handlers.RenderTemplate(w, "ascii-art", Text)
		//w.Header().Set("Content-Disposition", "attachment; filename=test.txt")
	default:
		http.Error(w, "ERROR: "+strconv.Itoa(http.StatusMethodNotAllowed)+". METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		fmt.Println("ERROR: " + strconv.Itoa(http.StatusMethodNotAllowed) + ". METHOD NOT ALLOWED")
	}
}

func SaveOutput(text string) {
	output := CreateFile("../web/download/res.txt")
	output.WriteString(text)
}

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	// Setting up the home page
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static"))))
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("../web/download"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", DisplayResult)

	// Start the serv on port : 8080
	port := 8080
	adresse := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost:%d...\n", port)
	http.ListenAndServe(adresse, nil)
}
