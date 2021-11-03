package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	dir := http.Dir("./htdocs")
	fileServer := http.FileServer(dir)
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/", fileServer)
	http.HandleFunc("/submit", formHandler)
	fmt.Println("Starting Web Server on Port 8000")
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("In Handler")
	io.WriteString(w, "Response handler \n")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", 400)
	}

	content := r.FormValue("content")
	content = strings.ToLower(strings.Trim(content, " "))
	contact := r.FormValue("contact")
	contact = strings.ToLower(strings.Trim(contact, " "))
	contact += ".log"
	path := "./htdocs/"
	//fmt.Println(contact)

	file, err := os.OpenFile(path+contact, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print(content)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
