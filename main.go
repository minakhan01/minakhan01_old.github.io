package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("*.html"))

func main() {
	http.HandleFunc("/", indexHandler)

	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/font-awesome-4.1.0/", http.StripPrefix("/font-awesome-4.1.0", http.FileServer(http.Dir("./font-awesome-4.1.0"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./img"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server didn't start: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", struct{}{})
	if err != nil {
		log.Printf("Failed to render template: %v\n", err)
	}
}
