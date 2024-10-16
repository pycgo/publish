package main

import (
	"net/http"
	"publish/web"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "index.html")
	//})
	http.HandleFunc("/submit", web.SubmitHandler)
	http.ListenAndServe(":8080", nil)
}
