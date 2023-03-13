package app

import (
	"html/template"
	"log"
	"net/http"
)


func newHandler(t *template.Template, name string) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request){
		data := struct {
			Title   string
			CSS     string
			JS      string
			Year    int
			Heading string
			Text    string
		}{
			Title:   "My App - About",
			//CSS:     template.URL("main.css"),
			JS:      "/static/js/main.js",
			Year:    2023,
			Heading: "About Us",
			Text:    "We are a team of developers building great apps!",
		}
		err := t.ExecuteTemplate(w, "index.html", data)
		log.Println(err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func StartApp() {
	templates := template.Must(template.ParseGlob("templates/*.html"))
	for _, x := range templates.Templates(){
		log.Println(x.Name())
	}

	// Create a file server for the static files
	fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/", newHandler(templates, "index.html"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}