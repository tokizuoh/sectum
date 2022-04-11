package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		color := r.FormValue("color")
		log.Println(name, color)
		http.Redirect(w, r, "/", 301)
	})
	http.ListenAndServe(":8080", nil)
}
