package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to my home</h1>")
}
func contactHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to my contact</h1>")
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", r)
}