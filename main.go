package main

import (
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", Anasayfa)
	http.HandleFunc("/detay", Detay)
	http.ListenAndServe(":8080", nil)

}
func Anasayfa(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html", "navbar.html")
	data := "Go'dan gelen veri"
	view.ExecuteTemplate(w, "anasayfa", data)

}

func Detay(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("detay.html", "navbar.html")
	view.ExecuteTemplate(w, "detay", nil)
}
