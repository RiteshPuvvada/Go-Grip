package main

import (
	"fmt"
	"log"
	"net/http"
)

func fh(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err : %v ", err)
	}
	fmt.Fprint(w, "POST request sucessful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func helloH(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
	}
	fmt.Fprint(w, "Hello !")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/form", fh)
	http.HandleFunc("/hello", helloH)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
